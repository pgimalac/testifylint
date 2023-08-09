package analyzer

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	"github.com/Antonboom/testifylint/config"
	"github.com/Antonboom/testifylint/internal/analysisutil"
	"github.com/Antonboom/testifylint/internal/checkers"
)

const (
	name = "testifylint"
	doc  = "Checks usage of github.com/stretchr/testify."
)

// New accepts validated config.Config and returns testifylint analyzer.
func New(cfg config.Config) *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: name,
		Doc:  doc,
		Run: func(pass *analysis.Pass) (any, error) {
			if err := config.Validate(cfg); err != nil {
				return nil, fmt.Errorf("invalid config: %v", err)
			}

			callCheckers, advancedCheckers, err := newCheckers(cfg)
			if err != nil {
				return nil, fmt.Errorf("build checkers: %v", err)
			}

			tl := &testifyLint{
				callCheckers:     callCheckers,
				advancedCheckers: advancedCheckers,
			}
			return tl.run(pass)
		},
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}
}

type testifyLint struct {
	callCheckers     []checkers.CallChecker
	advancedCheckers []checkers.AdvancedChecker
}

func (tl *testifyLint) run(pass *analysis.Pass) (any, error) {
	insp := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	for _, f := range pass.Files {
		if !analysisutil.IsTestFile(pass, f) {
			continue
		}

		if analysisutil.Imports(f, "github.com/stretchr/testify/assert") ||
			analysisutil.Imports(f, "github.com/stretchr/testify/require") ||
			analysisutil.Imports(f, "github.com/stretchr/testify/suite") {
			insp.Nodes([]ast.Node{
				(*ast.CallExpr)(nil),
				(*ast.FuncDecl)(nil),
			}, tl.newCallCheckersRunner(pass))

			for _, ch := range tl.advancedCheckers {
				ch.Check(pass, insp)
			}
		}
	}
	return nil, nil
}

func (tl *testifyLint) newCallCheckersRunner(pass *analysis.Pass) func(ast.Node, bool) bool {
	var insideSuiteMethod bool

	return func(node ast.Node, push bool) (proceed bool) {
		switch v := node.(type) {
		case *ast.FuncDecl:
			if analysisutil.IsSuiteMethod(pass, v) {
				if push {
					insideSuiteMethod = true
				} else {
					insideSuiteMethod = false
				}
			}

		case *ast.CallExpr:
			tl.checkCall(v, pass, insideSuiteMethod)
		}
		return true
	}
}

func (tl *testifyLint) checkCall(ce *ast.CallExpr, pass *analysis.Pass, insideSuiteMethod bool) {
	se, ok := ce.Fun.(*ast.SelectorExpr)
	if !ok {
		return
	}
	if se.Sel == nil {
		return
	}
	fn := se.Sel.Name

	pkg := func() *types.Package {
		if sel, ok := pass.TypesInfo.Selections[se]; ok {
			return sel.Obj().Pkg()
		}

		if id, ok := se.X.(*ast.Ident); ok {
			if selObj := pass.TypesInfo.ObjectOf(id); selObj != nil {
				if pkg, ok := selObj.(*types.PkgName); ok {
					return pkg.Imported()
				}
			}
		}
		return nil
	}()
	if pkg == nil {
		return
	}

	isAssert := analysisutil.IsPkg(pkg, "assert", "github.com/stretchr/testify/assert")
	isRequire := analysisutil.IsPkg(pkg, "require", "github.com/stretchr/testify/require")
	if !(isAssert || isRequire) {
		return
	}

	call := checkers.CallMeta{
		Range:             ce,
		Selector:          se,
		IsAssert:          isAssert,
		IsRequire:         isRequire,
		InsideSuiteMethod: insideSuiteMethod,
		SelectorStr:       types.ExprString(se.X),
		Fn: checkers.FnMeta{
			Range: se.Sel,
			Name:  fn,
			IsFmt: strings.HasSuffix(fn, "f"),
		},
		Args:    trimTArg(pass, ce.Args),
		ArgsRaw: ce.Args,
	}
	for _, ch := range tl.callCheckers {
		ch.Check(pass, call)
	}
}

func trimTArg(pass *analysis.Pass, args []ast.Expr) []ast.Expr {
	if len(args) == 0 {
		return args
	}

	if analysisutil.IsTestingTPtr(pass, args[0]) {
		return args[1:]
	}
	return args
}