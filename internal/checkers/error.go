package checkers

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
)

const ErrorCheckerName = "error"

// Error checks situation like
//
//	assert.Nil(t, err)
//
// and requires e.g.
//
//	assert.NoError(t, err)
type Error struct{}        //
func NewError() Error      { return Error{} }
func (Error) Name() string { return ErrorCheckerName }

func (checker Error) Check(pass *analysis.Pass, call CallMeta) {
	if len(call.Args) < 1 {
		return
	}
	errArg := call.Args[0]

	switch call.Fn.Name {
	case "NotNil", "NotNilf":
		if isError(pass, errArg) {
			r.ReportUseFunction(pass, checker.Name(), call, "Error",
				newFixViaFnReplacement(call, "Error"))
		}

	case "Nil", "Nilf":
		if isError(pass, errArg) {
			r.ReportUseFunction(pass, checker.Name(), call, "NoError",
				newFixViaFnReplacement(call, "NoError"))
		}
	}
}

var errIface = types.Universe.Lookup("error").Type().Underlying().(*types.Interface)

func isError(pass *analysis.Pass, expr ast.Expr) bool {
	t := pass.TypesInfo.TypeOf(expr)
	if t == nil {
		return false
	}

	_, ok := t.Underlying().(*types.Interface)
	return ok && types.Implements(t, errIface)
}