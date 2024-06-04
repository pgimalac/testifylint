// Code generated by testifylint/internal/testgen. DO NOT EDIT.

package formatter

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestFormatterChecker(t *testing.T) {
	var err error
	var args []any
	assert.Error(t, err, "Parse(%v) should fail.", args)

	assert.Equal(t, 1, 2)
	assert.Equal(t, 1, 2, "msg")
	assert.Equal(t, 1, 2, "msg with arg %d", 42)
	assert.Equal(t, 1, 2, "msg with args %d %s", 42, "42")
	assert.Equalf(t, 1, 2, "msg")
	assert.Equalf(t, 1, 2, "msg with arg %d", 42)
	assert.Equalf(t, 1, 2, "msg with args %d %s", 42, "42")

	assert.Equal(t, 1, 2, fmt.Sprintf("msg"))                           // want "formatter: remove unnecessary fmt\\.Sprintf"
	assert.Equal(t, 1, 2, fmt.Sprintf("msg with arg %d", 42))           // want "formatter: remove unnecessary fmt\\.Sprintf"
	assert.Equal(t, 1, 2, fmt.Sprintf("msg with args %d %s", 42, "42")) // want "formatter: remove unnecessary fmt\\.Sprintf"
	assert.Equal(t, 1, 2, fmt.Sprintf("msg"), 42)
	assert.Equal(t, 1, 2, fmt.Sprintf("msg with arg %d", 42), "42")
	assert.Equalf(t, 1, 2, fmt.Sprintf("msg"))                           // want "formatter: remove unnecessary fmt\\.Sprintf"
	assert.Equalf(t, 1, 2, fmt.Sprintf("msg with arg %d", 42))           // want "formatter: remove unnecessary fmt\\.Sprintf"
	assert.Equalf(t, 1, 2, fmt.Sprintf("msg with args %d %s", 42, "42")) // want "formatter: remove unnecessary fmt\\.Sprintf"
	assert.Equalf(t, 1, 2, fmt.Sprintf("msg"), 42)
	assert.Equalf(t, 1, 2, fmt.Sprintf("msg with arg %d", 42), "42")
}

func TestFormatterChecker_PrintfChecks(t *testing.T) {
	assert.Equalf(t, 1, 2, "msg with arg", "arg")        // want "formatter: assert\\.Equalf call has arguments but no formatting directives"
	assert.Equalf(t, 1, 2, "msg with arg: %w", nil)      // want "formatter: assert\\.Equalf does not support error-wrapping directive %w"
	assert.Equalf(t, 1, 2, "msg with args %d", 42, "42") // want "formatter: assert\\.Equalf call needs 1 arg but has 2 args"
	assert.Equalf(t, 1, 2, "msg with arg %[xd", 42)      // want "formatter: assert\\.Equalf format %\\[xd is missing closing \\]"

	assert.Equalf(t, 1, 2, "msg with arg %[3]*.[2*[1]f", 1, 2, 3) // want "formatter: assert\\.Equalf format has invalid argument index \\[2\\*\\[1\\]"

	assert.Equalf(t, 1, 2, "msg with arg %", 42)           // want "formatter: assert\\.Equalf format % is missing verb at end of string"
	assert.Equalf(t, 1, 2, "msg with arg %r", 42)          // want "formatter: assert\\.Equalf format %r has unknown verb r"
	assert.Equalf(t, 1, 2, "msg with arg %#s", 42)         // want "formatter: assert\\.Equalf format %#s has unrecognized flag #"
	assert.Equalf(t, 1, 2, "msg with arg %d", assertFalse) // want "formatter: assert\\.Equalf format %d arg assertFalse is a func value, not called"

	assert.Equalf(t, 1, 2, "msg with args %s %s", "42") // want "formatter: assert\\.Equalf format %s reads arg #2, but call has 1 arg$"
}

type FormatterCheckerSuite struct {
	suite.Suite
}

func TestFormatterCheckerSuite(t *testing.T) {
	suite.Run(t, new(FormatterCheckerSuite))
}

func (suite *FormatterCheckerSuite) TestFuncNameInDiagnostic() {
	require.Equalf(suite.T(), 1, 2, "msg with arg", "arg") // want "formatter: require\\.Equalf call has arguments but no formatting directives"

	suite.Require().Equalf(1, 2, "msg with arg", "arg") // want "formatter: suite\\.Require\\(\\)\\.Equalf call has arguments but no formatting directives"
	suite.Equalf(1, 2, "msg with arg", "arg")           // want "formatter: suite\\.Equalf call has arguments but no formatting directives"

	assertObj := assert.New(suite.T())
	assertObj.Equalf(1, 2, "msg with arg", "arg") // want "formatter: assertObj\\.Equalf call has arguments but no formatting directives"

	requireObj := require.New(suite.T())
	requireObj.Equalf(1, 2, "msg with arg", "arg") // want "formatter: requireObj\\.Equalf call has arguments but no formatting directives"
}

func assertFalse(t *testing.T, v bool, arg1 string, arg2 ...interface{}) {
	t.Helper()
	assert.Falsef(t, v, arg1, arg2...)
}

func TestFormatterChecker_AllAssertions(t *testing.T) {
	assert.Condition(t, nil)
	assert.Condition(t, nil, "msg")
	assert.Conditionf(t, nil, "msg")
	assert.Conditionf(t, nil, "msg with arg", 42) // want "formatter: assert\\.Conditionf call has arguments but no formatting directives"
	assert.Conditionf(t, nil, "msg with arg %d", 42)
	assert.Contains(t, nil, nil)
	assert.Contains(t, nil, nil, "msg")
	assert.Containsf(t, nil, nil, "msg")
	assert.Containsf(t, nil, nil, "msg with arg", 42) // want "formatter: assert\\.Containsf call has arguments but no formatting directives"
	assert.Containsf(t, nil, nil, "msg with arg %d", 42)
	assert.DirExists(t, "")
	assert.DirExists(t, "", "msg")
	assert.DirExistsf(t, "", "msg")
	assert.DirExistsf(t, "", "msg with arg", 42) // want "formatter: assert\\.DirExistsf call has arguments but no formatting directives"
	assert.DirExistsf(t, "", "msg with arg %d", 42)
	assert.ElementsMatch(t, nil, nil)
	assert.ElementsMatch(t, nil, nil, "msg")
	assert.ElementsMatchf(t, nil, nil, "msg")
	assert.ElementsMatchf(t, nil, nil, "msg with arg", 42) // want "formatter: assert\\.ElementsMatchf call has arguments but no formatting directives"
	assert.ElementsMatchf(t, nil, nil, "msg with arg %d", 42)
	assert.Empty(t, nil)
	assert.Empty(t, nil, "msg")
	assert.Emptyf(t, nil, "msg")
	assert.Emptyf(t, nil, "msg with arg", 42) // want "formatter: assert\\.Emptyf call has arguments but no formatting directives"
	assert.Emptyf(t, nil, "msg with arg %d", 42)
	assert.Equal(t, nil, nil)
	assert.Equal(t, nil, nil, "msg")
	assert.Equalf(t, nil, nil, "msg")
	assert.Equalf(t, nil, nil, "msg with arg", 42) // want "formatter: assert\\.Equalf call has arguments but no formatting directives"
	assert.Equalf(t, nil, nil, "msg with arg %d", 42)
	assert.EqualError(t, nil, "")
	assert.EqualError(t, nil, "", "msg")
	assert.EqualErrorf(t, nil, "", "msg")
	assert.EqualErrorf(t, nil, "", "msg with arg", 42) // want "formatter: assert\\.EqualErrorf call has arguments but no formatting directives"
	assert.EqualErrorf(t, nil, "", "msg with arg %d", 42)
	assert.EqualExportedValues(t, nil, nil)
	assert.EqualExportedValues(t, nil, nil, "msg")
	assert.EqualExportedValuesf(t, nil, nil, "msg")
	assert.EqualExportedValuesf(t, nil, nil, "msg with arg", 42) // want "formatter: assert\\.EqualExportedValuesf call has arguments but no formatting directives"
	assert.EqualExportedValuesf(t, nil, nil, "msg with arg %d", 42)
	assert.EqualValues(t, nil, nil)
	assert.EqualValues(t, nil, nil, "msg")
	assert.EqualValuesf(t, nil, nil, "msg")
	assert.EqualValuesf(t, nil, nil, "msg with arg", 42) // want "formatter: assert\\.EqualValuesf call has arguments but no formatting directives"
	assert.EqualValuesf(t, nil, nil, "msg with arg %d", 42)
	assert.Error(t, nil)
	assert.Error(t, nil, "msg")
	assert.Errorf(t, nil, "msg")
	assert.Errorf(t, nil, "msg with arg", 42) // want "formatter: assert\\.Errorf call has arguments but no formatting directives"
	assert.Errorf(t, nil, "msg with arg %d", 42)
	assert.ErrorAs(t, nil, nil)
	assert.ErrorAs(t, nil, nil, "msg")
	assert.ErrorAsf(t, nil, nil, "msg")
	assert.ErrorAsf(t, nil, nil, "msg with arg", 42) // want "formatter: assert\\.ErrorAsf call has arguments but no formatting directives"
	assert.ErrorAsf(t, nil, nil, "msg with arg %d", 42)
	assert.ErrorContains(t, nil, "")
	assert.ErrorContains(t, nil, "", "msg")
	assert.ErrorContainsf(t, nil, "", "msg")
	assert.ErrorContainsf(t, nil, "", "msg with arg", 42) // want "formatter: assert\\.ErrorContainsf call has arguments but no formatting directives"
	assert.ErrorContainsf(t, nil, "", "msg with arg %d", 42)
	assert.ErrorIs(t, nil, nil)
	assert.ErrorIs(t, nil, nil, "msg")
	assert.ErrorIsf(t, nil, nil, "msg")
	assert.ErrorIsf(t, nil, nil, "msg with arg", 42) // want "formatter: assert\\.ErrorIsf call has arguments but no formatting directives"
	assert.ErrorIsf(t, nil, nil, "msg with arg %d", 42)
	assert.Eventually(t, nil, 0, 0)
	assert.Eventually(t, nil, 0, 0, "msg")
	assert.Eventuallyf(t, nil, 0, 0, "msg")
	assert.Eventuallyf(t, nil, 0, 0, "msg with arg", 42) // want "formatter: assert\\.Eventuallyf call has arguments but no formatting directives"
	assert.Eventuallyf(t, nil, 0, 0, "msg with arg %d", 42)
	assert.EventuallyWithT(t, nil, 0, 0)
	assert.EventuallyWithT(t, nil, 0, 0, "msg")
	assert.EventuallyWithTf(t, nil, 0, 0, "msg")
	assert.EventuallyWithTf(t, nil, 0, 0, "msg with arg", 42) // want "formatter: assert\\.EventuallyWithTf call has arguments but no formatting directives"
	assert.EventuallyWithTf(t, nil, 0, 0, "msg with arg %d", 42)
	assert.Exactly(t, nil, nil)
	assert.Exactly(t, nil, nil, "msg")
	assert.Exactlyf(t, nil, nil, "msg")
	assert.Exactlyf(t, nil, nil, "msg with arg", 42) // want "formatter: assert\\.Exactlyf call has arguments but no formatting directives"
	assert.Exactlyf(t, nil, nil, "msg with arg %d", 42)
	assert.Fail(t, "")
	assert.Fail(t, "", "msg")
	assert.Failf(t, "", "msg")
	assert.Failf(t, "", "msg with arg", 42) // want "formatter: assert\\.Failf call has arguments but no formatting directives"
	assert.Failf(t, "", "msg with arg %d", 42)
	assert.FailNow(t, "")
	assert.FailNow(t, "", "msg")
	assert.FailNowf(t, "", "msg")
	assert.FailNowf(t, "", "msg with arg", 42) // want "formatter: assert\\.FailNowf call has arguments but no formatting directives"
	assert.FailNowf(t, "", "msg with arg %d", 42)
	assert.False(t, false)
	assert.False(t, false, "msg")
	assert.Falsef(t, false, "msg")
	assert.Falsef(t, false, "msg with arg", 42) // want "formatter: assert\\.Falsef call has arguments but no formatting directives"
	assert.Falsef(t, false, "msg with arg %d", 42)
	assert.FileExists(t, "")
	assert.FileExists(t, "", "msg")
	assert.FileExistsf(t, "", "msg")
	assert.FileExistsf(t, "", "msg with arg", 42) // want "formatter: assert\\.FileExistsf call has arguments but no formatting directives"
	assert.FileExistsf(t, "", "msg with arg %d", 42)
	assert.Implements(t, nil, nil)
	assert.Implements(t, nil, nil, "msg")
	assert.Implementsf(t, nil, nil, "msg")
	assert.Implementsf(t, nil, nil, "msg with arg", 42) // want "formatter: assert\\.Implementsf call has arguments but no formatting directives"
	assert.Implementsf(t, nil, nil, "msg with arg %d", 42)
	assert.InDelta(t, 0., 0., 0.)
	assert.InDelta(t, 0., 0., 0., "msg")
	assert.InDeltaf(t, 0., 0., 0., "msg")
	assert.InDeltaf(t, 0., 0., 0., "msg with arg", 42) // want "formatter: assert\\.InDeltaf call has arguments but no formatting directives"
	assert.InDeltaf(t, 0., 0., 0., "msg with arg %d", 42)
	assert.InDeltaMapValues(t, nil, nil, 0.)
	assert.InDeltaMapValues(t, nil, nil, 0., "msg")
	assert.InDeltaMapValuesf(t, nil, nil, 0., "msg")
	assert.InDeltaMapValuesf(t, nil, nil, 0., "msg with arg", 42) // want "formatter: assert\\.InDeltaMapValuesf call has arguments but no formatting directives"
	assert.InDeltaMapValuesf(t, nil, nil, 0., "msg with arg %d", 42)
	assert.InDeltaSlice(t, nil, nil, 0.)
	assert.InDeltaSlice(t, nil, nil, 0., "msg")
	assert.InDeltaSlicef(t, nil, nil, 0., "msg")
	assert.InDeltaSlicef(t, nil, nil, 0., "msg with arg", 42) // want "formatter: assert\\.InDeltaSlicef call has arguments but no formatting directives"
	assert.InDeltaSlicef(t, nil, nil, 0., "msg with arg %d", 42)
	assert.InEpsilon(t, nil, nil, 0.)
	assert.InEpsilon(t, nil, nil, 0., "msg")
	assert.InEpsilonf(t, nil, nil, 0., "msg")
	assert.InEpsilonf(t, nil, nil, 0., "msg with arg", 42) // want "formatter: assert\\.InEpsilonf call has arguments but no formatting directives"
	assert.InEpsilonf(t, nil, nil, 0., "msg with arg %d", 42)
	assert.InEpsilonSlice(t, nil, nil, 0.)
	assert.InEpsilonSlice(t, nil, nil, 0., "msg")
	assert.InEpsilonSlicef(t, nil, nil, 0., "msg")
	assert.InEpsilonSlicef(t, nil, nil, 0., "msg with arg", 42) // want "formatter: assert\\.InEpsilonSlicef call has arguments but no formatting directives"
	assert.InEpsilonSlicef(t, nil, nil, 0., "msg with arg %d", 42)
	assert.IsType(t, nil, nil)
	assert.IsType(t, nil, nil, "msg")
	assert.IsTypef(t, nil, nil, "msg")
	assert.IsTypef(t, nil, nil, "msg with arg", 42) // want "formatter: assert\\.IsTypef call has arguments but no formatting directives"
	assert.IsTypef(t, nil, nil, "msg with arg %d", 42)
	assert.JSONEq(t, "", "")
	assert.JSONEq(t, "", "", "msg")
	assert.JSONEqf(t, "", "", "msg")
	assert.JSONEqf(t, "", "", "msg with arg", 42) // want "formatter: assert\\.JSONEqf call has arguments but no formatting directives"
	assert.JSONEqf(t, "", "", "msg with arg %d", 42)
	assert.Len(t, nil, 0)
	assert.Len(t, nil, 0, "msg")
	assert.Lenf(t, nil, 0, "msg")
	assert.Lenf(t, nil, 0, "msg with arg", 42) // want "formatter: assert\\.Lenf call has arguments but no formatting directives"
	assert.Lenf(t, nil, 0, "msg with arg %d", 42)
	assert.Never(t, nil, 0, 0)
	assert.Never(t, nil, 0, 0, "msg")
	assert.Neverf(t, nil, 0, 0, "msg")
	assert.Neverf(t, nil, 0, 0, "msg with arg", 42) // want "formatter: assert\\.Neverf call has arguments but no formatting directives"
	assert.Neverf(t, nil, 0, 0, "msg with arg %d", 42)
	assert.Nil(t, nil)
	assert.Nil(t, nil, "msg")
	assert.Nilf(t, nil, "msg")
	assert.Nilf(t, nil, "msg with arg", 42) // want "formatter: assert\\.Nilf call has arguments but no formatting directives"
	assert.Nilf(t, nil, "msg with arg %d", 42)
	assert.NoDirExists(t, "")
	assert.NoDirExists(t, "", "msg")
	assert.NoDirExistsf(t, "", "msg")
	assert.NoDirExistsf(t, "", "msg with arg", 42) // want "formatter: assert\\.NoDirExistsf call has arguments but no formatting directives"
	assert.NoDirExistsf(t, "", "msg with arg %d", 42)
	assert.NoError(t, nil)
	assert.NoError(t, nil, "msg")
	assert.NoErrorf(t, nil, "msg")
	assert.NoErrorf(t, nil, "msg with arg", 42) // want "formatter: assert\\.NoErrorf call has arguments but no formatting directives"
	assert.NoErrorf(t, nil, "msg with arg %d", 42)
	assert.NoFileExists(t, "")
	assert.NoFileExists(t, "", "msg")
	assert.NoFileExistsf(t, "", "msg")
	assert.NoFileExistsf(t, "", "msg with arg", 42) // want "formatter: assert\\.NoFileExistsf call has arguments but no formatting directives"
	assert.NoFileExistsf(t, "", "msg with arg %d", 42)
	assert.NotContains(t, nil, nil)
	assert.NotContains(t, nil, nil, "msg")
	assert.NotContainsf(t, nil, nil, "msg")
	assert.NotContainsf(t, nil, nil, "msg with arg", 42) // want "formatter: assert\\.NotContainsf call has arguments but no formatting directives"
	assert.NotContainsf(t, nil, nil, "msg with arg %d", 42)
	assert.NotEmpty(t, nil)
	assert.NotEmpty(t, nil, "msg")
	assert.NotEmptyf(t, nil, "msg")
	assert.NotEmptyf(t, nil, "msg with arg", 42) // want "formatter: assert\\.NotEmptyf call has arguments but no formatting directives"
	assert.NotEmptyf(t, nil, "msg with arg %d", 42)
	assert.NotEqual(t, nil, nil)
	assert.NotEqual(t, nil, nil, "msg")
	assert.NotEqualf(t, nil, nil, "msg")
	assert.NotEqualf(t, nil, nil, "msg with arg", 42) // want "formatter: assert\\.NotEqualf call has arguments but no formatting directives"
	assert.NotEqualf(t, nil, nil, "msg with arg %d", 42)
	assert.NotEqualValues(t, nil, nil)
	assert.NotEqualValues(t, nil, nil, "msg")
	assert.NotEqualValuesf(t, nil, nil, "msg")
	assert.NotEqualValuesf(t, nil, nil, "msg with arg", 42) // want "formatter: assert\\.NotEqualValuesf call has arguments but no formatting directives"
	assert.NotEqualValuesf(t, nil, nil, "msg with arg %d", 42)
	assert.NotErrorIs(t, nil, nil)
	assert.NotErrorIs(t, nil, nil, "msg")
	assert.NotErrorIsf(t, nil, nil, "msg")
	assert.NotErrorIsf(t, nil, nil, "msg with arg", 42) // want "formatter: assert\\.NotErrorIsf call has arguments but no formatting directives"
	assert.NotErrorIsf(t, nil, nil, "msg with arg %d", 42)
	assert.NotNil(t, nil)
	assert.NotNil(t, nil, "msg")
	assert.NotNilf(t, nil, "msg")
	assert.NotNilf(t, nil, "msg with arg", 42) // want "formatter: assert\\.NotNilf call has arguments but no formatting directives"
	assert.NotNilf(t, nil, "msg with arg %d", 42)
	assert.NotPanics(t, nil)
	assert.NotPanics(t, nil, "msg")
	assert.NotPanicsf(t, nil, "msg")
	assert.NotPanicsf(t, nil, "msg with arg", 42) // want "formatter: assert\\.NotPanicsf call has arguments but no formatting directives"
	assert.NotPanicsf(t, nil, "msg with arg %d", 42)
	assert.NotRegexp(t, nil, "")
	assert.NotRegexp(t, nil, "", "msg")
	assert.NotRegexpf(t, nil, "", "msg")
	assert.NotRegexpf(t, nil, "", "msg with arg", 42) // want "formatter: assert\\.NotRegexpf call has arguments but no formatting directives"
	assert.NotRegexpf(t, nil, "", "msg with arg %d", 42)
	assert.NotSame(t, nil, nil)
	assert.NotSame(t, nil, nil, "msg")
	assert.NotSamef(t, nil, nil, "msg")
	assert.NotSamef(t, nil, nil, "msg with arg", 42) // want "formatter: assert\\.NotSamef call has arguments but no formatting directives"
	assert.NotSamef(t, nil, nil, "msg with arg %d", 42)
	assert.NotSubset(t, nil, nil)
	assert.NotSubset(t, nil, nil, "msg")
	assert.NotSubsetf(t, nil, nil, "msg")
	assert.NotSubsetf(t, nil, nil, "msg with arg", 42) // want "formatter: assert\\.NotSubsetf call has arguments but no formatting directives"
	assert.NotSubsetf(t, nil, nil, "msg with arg %d", 42)
	assert.NotZero(t, nil)
	assert.NotZero(t, nil, "msg")
	assert.NotZerof(t, nil, "msg")
	assert.NotZerof(t, nil, "msg with arg", 42) // want "formatter: assert\\.NotZerof call has arguments but no formatting directives"
	assert.NotZerof(t, nil, "msg with arg %d", 42)
	assert.Panics(t, nil)
	assert.Panics(t, nil, "msg")
	assert.Panicsf(t, nil, "msg")
	assert.Panicsf(t, nil, "msg with arg", 42) // want "formatter: assert\\.Panicsf call has arguments but no formatting directives"
	assert.Panicsf(t, nil, "msg with arg %d", 42)
	assert.PanicsWithError(t, "", nil)
	assert.PanicsWithError(t, "", nil, "msg")
	assert.PanicsWithErrorf(t, "", nil, "msg")
	assert.PanicsWithErrorf(t, "", nil, "msg with arg", 42) // want "formatter: assert\\.PanicsWithErrorf call has arguments but no formatting directives"
	assert.PanicsWithErrorf(t, "", nil, "msg with arg %d", 42)
	assert.PanicsWithValue(t, nil, nil)
	assert.PanicsWithValue(t, nil, nil, "msg")
	assert.PanicsWithValuef(t, nil, nil, "msg")
	assert.PanicsWithValuef(t, nil, nil, "msg with arg", 42) // want "formatter: assert\\.PanicsWithValuef call has arguments but no formatting directives"
	assert.PanicsWithValuef(t, nil, nil, "msg with arg %d", 42)
	assert.Regexp(t, nil, nil)
	assert.Regexp(t, nil, nil, "msg")
	assert.Regexpf(t, nil, nil, "msg")
	assert.Regexpf(t, nil, nil, "msg with arg", 42) // want "formatter: assert\\.Regexpf call has arguments but no formatting directives"
	assert.Regexpf(t, nil, nil, "msg with arg %d", 42)
	assert.Same(t, nil, nil)
	assert.Same(t, nil, nil, "msg")
	assert.Samef(t, nil, nil, "msg")
	assert.Samef(t, nil, nil, "msg with arg", 42) // want "formatter: assert\\.Samef call has arguments but no formatting directives"
	assert.Samef(t, nil, nil, "msg with arg %d", 42)
	assert.Subset(t, nil, nil)
	assert.Subset(t, nil, nil, "msg")
	assert.Subsetf(t, nil, nil, "msg")
	assert.Subsetf(t, nil, nil, "msg with arg", 42) // want "formatter: assert\\.Subsetf call has arguments but no formatting directives"
	assert.Subsetf(t, nil, nil, "msg with arg %d", 42)
	assert.True(t, true)
	assert.True(t, true, "msg")
	assert.Truef(t, true, "msg")
	assert.Truef(t, true, "msg with arg", 42) // want "formatter: assert\\.Truef call has arguments but no formatting directives"
	assert.Truef(t, true, "msg with arg %d", 42)
	assert.WithinDuration(t, time.Time{}, time.Time{}, 0)
	assert.WithinDuration(t, time.Time{}, time.Time{}, 0, "msg")
	assert.WithinDurationf(t, time.Time{}, time.Time{}, 0, "msg")
	assert.WithinDurationf(t, time.Time{}, time.Time{}, 0, "msg with arg", 42) // want "formatter: assert\\.WithinDurationf call has arguments but no formatting directives"
	assert.WithinDurationf(t, time.Time{}, time.Time{}, 0, "msg with arg %d", 42)
	assert.WithinRange(t, time.Time{}, time.Time{}, time.Time{})
	assert.WithinRange(t, time.Time{}, time.Time{}, time.Time{}, "msg")
	assert.WithinRangef(t, time.Time{}, time.Time{}, time.Time{}, "msg")
	assert.WithinRangef(t, time.Time{}, time.Time{}, time.Time{}, "msg with arg", 42) // want "formatter: assert\\.WithinRangef call has arguments but no formatting directives"
	assert.WithinRangef(t, time.Time{}, time.Time{}, time.Time{}, "msg with arg %d", 42)
	assert.YAMLEq(t, "", "")
	assert.YAMLEq(t, "", "", "msg")
	assert.YAMLEqf(t, "", "", "msg")
	assert.YAMLEqf(t, "", "", "msg with arg", 42) // want "formatter: assert\\.YAMLEqf call has arguments but no formatting directives"
	assert.YAMLEqf(t, "", "", "msg with arg %d", 42)
	assert.Zero(t, nil)
	assert.Zero(t, nil, "msg")
	assert.Zerof(t, nil, "msg")
	assert.Zerof(t, nil, "msg with arg", 42) // want "formatter: assert\\.Zerof call has arguments but no formatting directives"
	assert.Zerof(t, nil, "msg with arg %d", 42)
}

func TestFormatterChecker_Ignored(t *testing.T) {
	assert.ObjectsAreEqual(nil, nil)
	assert.ObjectsAreEqualValues(nil, nil)
	assert.ObjectsExportedFieldsAreEqual(nil, nil)
}
