package checkers

import (
	"fmt"
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"

	"github.com/Antonboom/testifylint/internal/analysisutil"
)

// ErrorIsAs detects situations like
//
//	assert.Error(t, err, errSentinel)
//	assert.NoError(t, err, errSentinel)
//	assert.True(t, errors.Is(err, errSentinel))
//	assert.False(t, errors.Is(err, errSentinel))
//	assert.True(t, errors.As(err, &target))
//
// and requires
//
//	assert.ErrorIs(t, err, errSentinel)
//	assert.NotErrorIs(t, err, errSentinel)
//	assert.ErrorAs(t, err, &target)
//
// Also ErrorIsAs repeats go vet's "errorsas" check logic.
type ErrorIsAs struct{}

// NewErrorIsAs constructs ErrorIsAs checker.
func NewErrorIsAs() ErrorIsAs  { return ErrorIsAs{} }
func (ErrorIsAs) Name() string { return "error-is-as" }

func (checker ErrorIsAs) Check(pass *analysis.Pass, call *CallMeta) *analysis.Diagnostic {
	switch call.Fn.Name {
	case "Error", "Errorf":
		if len(call.Args) >= 2 && isError(pass, call.Args[1]) {
			const proposed = "ErrorIs"
			msg := fmt.Sprintf("invalid usage of %[1]s.Error, use %[1]s.%[2]s instead", call.SelectorXStr, proposed)
			return newDiagnostic(checker.Name(), call, msg, newSuggestedFuncReplacement(call, proposed))
		}

	case "NoError", "NoErrorf":
		if len(call.Args) >= 2 && isError(pass, call.Args[1]) {
			const proposed = "NotErrorIs"
			msg := fmt.Sprintf("invalid usage of %[1]s.NoError, use %[1]s.%[2]s instead", call.SelectorXStr, proposed)
			return newDiagnostic(checker.Name(), call, msg, newSuggestedFuncReplacement(call, proposed))
		}

	case "True", "Truef":
		if len(call.Args) < 1 {
			return nil
		}

		ce, ok := call.Args[0].(*ast.CallExpr)
		if !ok {
			return nil
		}
		if len(ce.Args) != 2 {
			return nil
		}

		var proposed string
		switch {
		case isErrorsIsCall(pass, ce):
			proposed = "ErrorIs"
		case isErrorsAsCall(pass, ce):
			proposed = "ErrorAs"
		}
		if proposed != "" {
			return newUseFunctionDiagnostic(checker.Name(), call, proposed,
				newSuggestedFuncReplacement(call, proposed, analysis.TextEdit{
					Pos:     ce.Pos(),
					End:     ce.End(),
					NewText: formatAsCallArgs(pass, ce.Args[0], ce.Args[1]),
				}),
			)
		}

	case "False", "Falsef":
		if len(call.Args) < 1 {
			return nil
		}

		ce, ok := call.Args[0].(*ast.CallExpr)
		if !ok {
			return nil
		}
		if len(ce.Args) != 2 {
			return nil
		}

		if isErrorsIsCall(pass, ce) {
			const proposed = "NotErrorIs"
			return newUseFunctionDiagnostic(checker.Name(), call, proposed,
				newSuggestedFuncReplacement(call, proposed, analysis.TextEdit{
					Pos:     ce.Pos(),
					End:     ce.End(),
					NewText: formatAsCallArgs(pass, ce.Args[0], ce.Args[1]),
				}),
			)
		}

	case "ErrorAs", "ErrorAsf":
		if len(call.Args) < 2 {
			return nil
		}

		// NOTE(a.telyshev): Logic below must be consistent with
		// https://cs.opensource.google/go/x/tools/+/master:go/analysis/passes/errorsas/errorsas.go

		var (
			defaultReport  = fmt.Sprintf("second argument to %s must be a non-nil pointer to either a type that implements error, or to any interface type", call) //nolint:lll
			errorPtrReport = fmt.Sprintf("second argument to %s should not be *error", call)
		)

		target := call.Args[1]
		t := pass.TypesInfo.Types[target].Type
		if it, ok := t.Underlying().(*types.Interface); ok && it.NumMethods() == 0 {
			// `any` interface case. It is always allowed, since it often indicates
			// a value forwarded from another source.
			return nil
		}

		pt, ok := t.Underlying().(*types.Pointer)
		if !ok {
			return newDiagnostic(checker.Name(), call, defaultReport, nil)
		}
		if pt.Elem() == errorType {
			return newDiagnostic(checker.Name(), call, errorPtrReport, nil)
		}

		_, isInterface := pt.Elem().Underlying().(*types.Interface)
		if !isInterface && !types.Implements(pt.Elem(), errorIface) {
			return newDiagnostic(checker.Name(), call, defaultReport, nil)
		}
	}
	return nil
}

func isErrorsIsCall(pass *analysis.Pass, ce *ast.CallExpr) bool {
	return isErrorsPkgFnCall(pass, ce, "Is")
}

func isErrorsAsCall(pass *analysis.Pass, ce *ast.CallExpr) bool {
	return isErrorsPkgFnCall(pass, ce, "As")
}

func isErrorsPkgFnCall(pass *analysis.Pass, ce *ast.CallExpr, fn string) bool {
	se, ok := ce.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	errorsIsObj := analysisutil.ObjectOf(pass.Pkg, "errors", fn)
	if errorsIsObj == nil {
		return false
	}

	return analysisutil.IsObj(pass.TypesInfo, se.Sel, errorsIsObj)
}
