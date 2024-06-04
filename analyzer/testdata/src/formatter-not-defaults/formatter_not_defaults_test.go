// Code generated by testifylint/internal/testgen. DO NOT EDIT.

package formatter

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFormatterChecker(t *testing.T) {
	var err error
	var args []any
	assert.Error(t, err, "Parse(%v) should fail.", args) // want "formatter: use assert\\.Errorf$"
	assert.Equal(t, 1, 2)
	assert.Equal(t, 1, 2, "msg")                           // want "formatter: use assert\\.Equalf"
	assert.Equal(t, 1, 2, "msg with arg %d", 42)           // want "formatter: use assert\\.Equalf"
	assert.Equal(t, 1, 2, "msg with args %d %s", 42, "42") // want "formatter: use assert\\.Equalf"
	assert.Equalf(t, 1, 2, "msg")
	assert.Equalf(t, 1, 2, "msg with arg %d", 42)
	assert.Equalf(t, 1, 2, "msg with args %d %s", 42, "42")

	assert.Equal(t, 1, 2, fmt.Sprintf("msg"))                            // want "formatter: remove unnecessary fmt\\.Sprintf and use assert\\.Equalf"
	assert.Equal(t, 1, 2, fmt.Sprintf("msg with arg %d", 42))            // want "formatter: remove unnecessary fmt\\.Sprintf and use assert\\.Equalf"
	assert.Equal(t, 1, 2, fmt.Sprintf("msg with args %d %s", 42, "42"))  // want "formatter: remove unnecessary fmt\\.Sprintf and use assert\\.Equalf"
	assert.Equal(t, 1, 2, fmt.Sprintf("msg"), 42)                        // want "formatter: use assert\\.Equalf"
	assert.Equal(t, 1, 2, fmt.Sprintf("msg with arg %d", 42), "42")      // want "formatter: use assert\\.Equalf"
	assert.Equalf(t, 1, 2, fmt.Sprintf("msg"))                           // want "formatter: remove unnecessary fmt\\.Sprintf"
	assert.Equalf(t, 1, 2, fmt.Sprintf("msg with arg %d", 42))           // want "formatter: remove unnecessary fmt\\.Sprintf"
	assert.Equalf(t, 1, 2, fmt.Sprintf("msg with args %d %s", 42, "42")) // want "formatter: remove unnecessary fmt\\.Sprintf"
	assert.Equalf(t, 1, 2, fmt.Sprintf("msg"), 42)
	assert.Equalf(t, 1, 2, fmt.Sprintf("msg with arg %d", 42), "42")
}

func TestFormatterChecker_AllAssertions(t *testing.T) {
	assert.Condition(t, nil)
	assert.Condition(t, nil, "msg") // want "formatter: use assert\\.Conditionf"
	assert.Conditionf(t, nil, "msg")
	assert.Conditionf(t, nil, "msg with arg", 42)
	assert.Conditionf(t, nil, "msg with arg %d", 42)
	assert.Contains(t, nil, nil)
	assert.Contains(t, nil, nil, "msg") // want "formatter: use assert\\.Containsf"
	assert.Containsf(t, nil, nil, "msg")
	assert.Containsf(t, nil, nil, "msg with arg", 42)
	assert.Containsf(t, nil, nil, "msg with arg %d", 42)
	assert.DirExists(t, "")
	assert.DirExists(t, "", "msg") // want "formatter: use assert\\.DirExistsf"
	assert.DirExistsf(t, "", "msg")
	assert.DirExistsf(t, "", "msg with arg", 42)
	assert.DirExistsf(t, "", "msg with arg %d", 42)
	assert.ElementsMatch(t, nil, nil)
	assert.ElementsMatch(t, nil, nil, "msg") // want "formatter: use assert\\.ElementsMatchf"
	assert.ElementsMatchf(t, nil, nil, "msg")
	assert.ElementsMatchf(t, nil, nil, "msg with arg", 42)
	assert.ElementsMatchf(t, nil, nil, "msg with arg %d", 42)
	assert.Empty(t, nil)
	assert.Empty(t, nil, "msg") // want "formatter: use assert\\.Emptyf"
	assert.Emptyf(t, nil, "msg")
	assert.Emptyf(t, nil, "msg with arg", 42)
	assert.Emptyf(t, nil, "msg with arg %d", 42)
	assert.Equal(t, nil, nil)
	assert.Equal(t, nil, nil, "msg") // want "formatter: use assert\\.Equalf"
	assert.Equalf(t, nil, nil, "msg")
	assert.Equalf(t, nil, nil, "msg with arg", 42)
	assert.Equalf(t, nil, nil, "msg with arg %d", 42)
	assert.EqualError(t, nil, "")
	assert.EqualError(t, nil, "", "msg") // want "formatter: use assert\\.EqualErrorf"
	assert.EqualErrorf(t, nil, "", "msg")
	assert.EqualErrorf(t, nil, "", "msg with arg", 42)
	assert.EqualErrorf(t, nil, "", "msg with arg %d", 42)
	assert.EqualExportedValues(t, nil, nil)
	assert.EqualExportedValues(t, nil, nil, "msg") // want "formatter: use assert\\.EqualExportedValuesf"
	assert.EqualExportedValuesf(t, nil, nil, "msg")
	assert.EqualExportedValuesf(t, nil, nil, "msg with arg", 42)
	assert.EqualExportedValuesf(t, nil, nil, "msg with arg %d", 42)
	assert.EqualValues(t, nil, nil)
	assert.EqualValues(t, nil, nil, "msg") // want "formatter: use assert\\.EqualValuesf"
	assert.EqualValuesf(t, nil, nil, "msg")
	assert.EqualValuesf(t, nil, nil, "msg with arg", 42)
	assert.EqualValuesf(t, nil, nil, "msg with arg %d", 42)
	assert.Error(t, nil)
	assert.Error(t, nil, "msg") // want "formatter: use assert\\.Errorf"
	assert.Errorf(t, nil, "msg")
	assert.Errorf(t, nil, "msg with arg", 42)
	assert.Errorf(t, nil, "msg with arg %d", 42)
	assert.ErrorAs(t, nil, nil)
	assert.ErrorAs(t, nil, nil, "msg") // want "formatter: use assert\\.ErrorAsf"
	assert.ErrorAsf(t, nil, nil, "msg")
	assert.ErrorAsf(t, nil, nil, "msg with arg", 42)
	assert.ErrorAsf(t, nil, nil, "msg with arg %d", 42)
	assert.ErrorContains(t, nil, "")
	assert.ErrorContains(t, nil, "", "msg") // want "formatter: use assert\\.ErrorContainsf"
	assert.ErrorContainsf(t, nil, "", "msg")
	assert.ErrorContainsf(t, nil, "", "msg with arg", 42)
	assert.ErrorContainsf(t, nil, "", "msg with arg %d", 42)
	assert.ErrorIs(t, nil, nil)
	assert.ErrorIs(t, nil, nil, "msg") // want "formatter: use assert\\.ErrorIsf"
	assert.ErrorIsf(t, nil, nil, "msg")
	assert.ErrorIsf(t, nil, nil, "msg with arg", 42)
	assert.ErrorIsf(t, nil, nil, "msg with arg %d", 42)
	assert.Eventually(t, nil, 0, 0)
	assert.Eventually(t, nil, 0, 0, "msg") // want "formatter: use assert\\.Eventuallyf"
	assert.Eventuallyf(t, nil, 0, 0, "msg")
	assert.Eventuallyf(t, nil, 0, 0, "msg with arg", 42)
	assert.Eventuallyf(t, nil, 0, 0, "msg with arg %d", 42)
	assert.EventuallyWithT(t, nil, 0, 0)
	assert.EventuallyWithT(t, nil, 0, 0, "msg") // want "formatter: use assert\\.EventuallyWithTf"
	assert.EventuallyWithTf(t, nil, 0, 0, "msg")
	assert.EventuallyWithTf(t, nil, 0, 0, "msg with arg", 42)
	assert.EventuallyWithTf(t, nil, 0, 0, "msg with arg %d", 42)
	assert.Exactly(t, nil, nil)
	assert.Exactly(t, nil, nil, "msg") // want "formatter: use assert\\.Exactlyf"
	assert.Exactlyf(t, nil, nil, "msg")
	assert.Exactlyf(t, nil, nil, "msg with arg", 42)
	assert.Exactlyf(t, nil, nil, "msg with arg %d", 42)
	assert.Fail(t, "")
	assert.Fail(t, "", "msg") // want "formatter: use assert\\.Failf"
	assert.Failf(t, "", "msg")
	assert.Failf(t, "", "msg with arg", 42)
	assert.Failf(t, "", "msg with arg %d", 42)
	assert.FailNow(t, "")
	assert.FailNow(t, "", "msg") // want "formatter: use assert\\.FailNowf"
	assert.FailNowf(t, "", "msg")
	assert.FailNowf(t, "", "msg with arg", 42)
	assert.FailNowf(t, "", "msg with arg %d", 42)
	assert.False(t, false)
	assert.False(t, false, "msg") // want "formatter: use assert\\.Falsef"
	assert.Falsef(t, false, "msg")
	assert.Falsef(t, false, "msg with arg", 42)
	assert.Falsef(t, false, "msg with arg %d", 42)
	assert.FileExists(t, "")
	assert.FileExists(t, "", "msg") // want "formatter: use assert\\.FileExistsf"
	assert.FileExistsf(t, "", "msg")
	assert.FileExistsf(t, "", "msg with arg", 42)
	assert.FileExistsf(t, "", "msg with arg %d", 42)
	assert.Implements(t, nil, nil)
	assert.Implements(t, nil, nil, "msg") // want "formatter: use assert\\.Implementsf"
	assert.Implementsf(t, nil, nil, "msg")
	assert.Implementsf(t, nil, nil, "msg with arg", 42)
	assert.Implementsf(t, nil, nil, "msg with arg %d", 42)
	assert.InDelta(t, 0., 0., 0.)
	assert.InDelta(t, 0., 0., 0., "msg") // want "formatter: use assert\\.InDeltaf"
	assert.InDeltaf(t, 0., 0., 0., "msg")
	assert.InDeltaf(t, 0., 0., 0., "msg with arg", 42)
	assert.InDeltaf(t, 0., 0., 0., "msg with arg %d", 42)
	assert.InDeltaMapValues(t, nil, nil, 0.)
	assert.InDeltaMapValues(t, nil, nil, 0., "msg") // want "formatter: use assert\\.InDeltaMapValuesf"
	assert.InDeltaMapValuesf(t, nil, nil, 0., "msg")
	assert.InDeltaMapValuesf(t, nil, nil, 0., "msg with arg", 42)
	assert.InDeltaMapValuesf(t, nil, nil, 0., "msg with arg %d", 42)
	assert.InDeltaSlice(t, nil, nil, 0.)
	assert.InDeltaSlice(t, nil, nil, 0., "msg") // want "formatter: use assert\\.InDeltaSlicef"
	assert.InDeltaSlicef(t, nil, nil, 0., "msg")
	assert.InDeltaSlicef(t, nil, nil, 0., "msg with arg", 42)
	assert.InDeltaSlicef(t, nil, nil, 0., "msg with arg %d", 42)
	assert.InEpsilon(t, nil, nil, 0.)
	assert.InEpsilon(t, nil, nil, 0., "msg") // want "formatter: use assert\\.InEpsilonf"
	assert.InEpsilonf(t, nil, nil, 0., "msg")
	assert.InEpsilonf(t, nil, nil, 0., "msg with arg", 42)
	assert.InEpsilonf(t, nil, nil, 0., "msg with arg %d", 42)
	assert.InEpsilonSlice(t, nil, nil, 0.)
	assert.InEpsilonSlice(t, nil, nil, 0., "msg") // want "formatter: use assert\\.InEpsilonSlicef"
	assert.InEpsilonSlicef(t, nil, nil, 0., "msg")
	assert.InEpsilonSlicef(t, nil, nil, 0., "msg with arg", 42)
	assert.InEpsilonSlicef(t, nil, nil, 0., "msg with arg %d", 42)
	assert.IsType(t, nil, nil)
	assert.IsType(t, nil, nil, "msg") // want "formatter: use assert\\.IsTypef"
	assert.IsTypef(t, nil, nil, "msg")
	assert.IsTypef(t, nil, nil, "msg with arg", 42)
	assert.IsTypef(t, nil, nil, "msg with arg %d", 42)
	assert.JSONEq(t, "", "")
	assert.JSONEq(t, "", "", "msg") // want "formatter: use assert\\.JSONEqf"
	assert.JSONEqf(t, "", "", "msg")
	assert.JSONEqf(t, "", "", "msg with arg", 42)
	assert.JSONEqf(t, "", "", "msg with arg %d", 42)
	assert.Len(t, nil, 0)
	assert.Len(t, nil, 0, "msg") // want "formatter: use assert\\.Lenf"
	assert.Lenf(t, nil, 0, "msg")
	assert.Lenf(t, nil, 0, "msg with arg", 42)
	assert.Lenf(t, nil, 0, "msg with arg %d", 42)
	assert.Never(t, nil, 0, 0)
	assert.Never(t, nil, 0, 0, "msg") // want "formatter: use assert\\.Neverf"
	assert.Neverf(t, nil, 0, 0, "msg")
	assert.Neverf(t, nil, 0, 0, "msg with arg", 42)
	assert.Neverf(t, nil, 0, 0, "msg with arg %d", 42)
	assert.Nil(t, nil)
	assert.Nil(t, nil, "msg") // want "formatter: use assert\\.Nilf"
	assert.Nilf(t, nil, "msg")
	assert.Nilf(t, nil, "msg with arg", 42)
	assert.Nilf(t, nil, "msg with arg %d", 42)
	assert.NoDirExists(t, "")
	assert.NoDirExists(t, "", "msg") // want "formatter: use assert\\.NoDirExistsf"
	assert.NoDirExistsf(t, "", "msg")
	assert.NoDirExistsf(t, "", "msg with arg", 42)
	assert.NoDirExistsf(t, "", "msg with arg %d", 42)
	assert.NoError(t, nil)
	assert.NoError(t, nil, "msg") // want "formatter: use assert\\.NoErrorf"
	assert.NoErrorf(t, nil, "msg")
	assert.NoErrorf(t, nil, "msg with arg", 42)
	assert.NoErrorf(t, nil, "msg with arg %d", 42)
	assert.NoFileExists(t, "")
	assert.NoFileExists(t, "", "msg") // want "formatter: use assert\\.NoFileExistsf"
	assert.NoFileExistsf(t, "", "msg")
	assert.NoFileExistsf(t, "", "msg with arg", 42)
	assert.NoFileExistsf(t, "", "msg with arg %d", 42)
	assert.NotContains(t, nil, nil)
	assert.NotContains(t, nil, nil, "msg") // want "formatter: use assert\\.NotContainsf"
	assert.NotContainsf(t, nil, nil, "msg")
	assert.NotContainsf(t, nil, nil, "msg with arg", 42)
	assert.NotContainsf(t, nil, nil, "msg with arg %d", 42)
	assert.NotEmpty(t, nil)
	assert.NotEmpty(t, nil, "msg") // want "formatter: use assert\\.NotEmptyf"
	assert.NotEmptyf(t, nil, "msg")
	assert.NotEmptyf(t, nil, "msg with arg", 42)
	assert.NotEmptyf(t, nil, "msg with arg %d", 42)
	assert.NotEqual(t, nil, nil)
	assert.NotEqual(t, nil, nil, "msg") // want "formatter: use assert\\.NotEqualf"
	assert.NotEqualf(t, nil, nil, "msg")
	assert.NotEqualf(t, nil, nil, "msg with arg", 42)
	assert.NotEqualf(t, nil, nil, "msg with arg %d", 42)
	assert.NotEqualValues(t, nil, nil)
	assert.NotEqualValues(t, nil, nil, "msg") // want "formatter: use assert\\.NotEqualValuesf"
	assert.NotEqualValuesf(t, nil, nil, "msg")
	assert.NotEqualValuesf(t, nil, nil, "msg with arg", 42)
	assert.NotEqualValuesf(t, nil, nil, "msg with arg %d", 42)
	assert.NotErrorIs(t, nil, nil)
	assert.NotErrorIs(t, nil, nil, "msg") // want "formatter: use assert\\.NotErrorIsf"
	assert.NotErrorIsf(t, nil, nil, "msg")
	assert.NotErrorIsf(t, nil, nil, "msg with arg", 42)
	assert.NotErrorIsf(t, nil, nil, "msg with arg %d", 42)
	assert.NotNil(t, nil)
	assert.NotNil(t, nil, "msg") // want "formatter: use assert\\.NotNilf"
	assert.NotNilf(t, nil, "msg")
	assert.NotNilf(t, nil, "msg with arg", 42)
	assert.NotNilf(t, nil, "msg with arg %d", 42)
	assert.NotPanics(t, nil)
	assert.NotPanics(t, nil, "msg") // want "formatter: use assert\\.NotPanicsf"
	assert.NotPanicsf(t, nil, "msg")
	assert.NotPanicsf(t, nil, "msg with arg", 42)
	assert.NotPanicsf(t, nil, "msg with arg %d", 42)
	assert.NotRegexp(t, nil, "")
	assert.NotRegexp(t, nil, "", "msg") // want "formatter: use assert\\.NotRegexpf"
	assert.NotRegexpf(t, nil, "", "msg")
	assert.NotRegexpf(t, nil, "", "msg with arg", 42)
	assert.NotRegexpf(t, nil, "", "msg with arg %d", 42)
	assert.NotSame(t, nil, nil)
	assert.NotSame(t, nil, nil, "msg") // want "formatter: use assert\\.NotSamef"
	assert.NotSamef(t, nil, nil, "msg")
	assert.NotSamef(t, nil, nil, "msg with arg", 42)
	assert.NotSamef(t, nil, nil, "msg with arg %d", 42)
	assert.NotSubset(t, nil, nil)
	assert.NotSubset(t, nil, nil, "msg") // want "formatter: use assert\\.NotSubsetf"
	assert.NotSubsetf(t, nil, nil, "msg")
	assert.NotSubsetf(t, nil, nil, "msg with arg", 42)
	assert.NotSubsetf(t, nil, nil, "msg with arg %d", 42)
	assert.NotZero(t, nil)
	assert.NotZero(t, nil, "msg") // want "formatter: use assert\\.NotZerof"
	assert.NotZerof(t, nil, "msg")
	assert.NotZerof(t, nil, "msg with arg", 42)
	assert.NotZerof(t, nil, "msg with arg %d", 42)
	assert.Panics(t, nil)
	assert.Panics(t, nil, "msg") // want "formatter: use assert\\.Panicsf"
	assert.Panicsf(t, nil, "msg")
	assert.Panicsf(t, nil, "msg with arg", 42)
	assert.Panicsf(t, nil, "msg with arg %d", 42)
	assert.PanicsWithError(t, "", nil)
	assert.PanicsWithError(t, "", nil, "msg") // want "formatter: use assert\\.PanicsWithErrorf"
	assert.PanicsWithErrorf(t, "", nil, "msg")
	assert.PanicsWithErrorf(t, "", nil, "msg with arg", 42)
	assert.PanicsWithErrorf(t, "", nil, "msg with arg %d", 42)
	assert.PanicsWithValue(t, nil, nil)
	assert.PanicsWithValue(t, nil, nil, "msg") // want "formatter: use assert\\.PanicsWithValuef"
	assert.PanicsWithValuef(t, nil, nil, "msg")
	assert.PanicsWithValuef(t, nil, nil, "msg with arg", 42)
	assert.PanicsWithValuef(t, nil, nil, "msg with arg %d", 42)
	assert.Regexp(t, nil, nil)
	assert.Regexp(t, nil, nil, "msg") // want "formatter: use assert\\.Regexpf"
	assert.Regexpf(t, nil, nil, "msg")
	assert.Regexpf(t, nil, nil, "msg with arg", 42)
	assert.Regexpf(t, nil, nil, "msg with arg %d", 42)
	assert.Same(t, nil, nil)
	assert.Same(t, nil, nil, "msg") // want "formatter: use assert\\.Samef"
	assert.Samef(t, nil, nil, "msg")
	assert.Samef(t, nil, nil, "msg with arg", 42)
	assert.Samef(t, nil, nil, "msg with arg %d", 42)
	assert.Subset(t, nil, nil)
	assert.Subset(t, nil, nil, "msg") // want "formatter: use assert\\.Subsetf"
	assert.Subsetf(t, nil, nil, "msg")
	assert.Subsetf(t, nil, nil, "msg with arg", 42)
	assert.Subsetf(t, nil, nil, "msg with arg %d", 42)
	assert.True(t, true)
	assert.True(t, true, "msg") // want "formatter: use assert\\.Truef"
	assert.Truef(t, true, "msg")
	assert.Truef(t, true, "msg with arg", 42)
	assert.Truef(t, true, "msg with arg %d", 42)
	assert.WithinDuration(t, time.Time{}, time.Time{}, 0)
	assert.WithinDuration(t, time.Time{}, time.Time{}, 0, "msg") // want "formatter: use assert\\.WithinDurationf"
	assert.WithinDurationf(t, time.Time{}, time.Time{}, 0, "msg")
	assert.WithinDurationf(t, time.Time{}, time.Time{}, 0, "msg with arg", 42)
	assert.WithinDurationf(t, time.Time{}, time.Time{}, 0, "msg with arg %d", 42)
	assert.WithinRange(t, time.Time{}, time.Time{}, time.Time{})
	assert.WithinRange(t, time.Time{}, time.Time{}, time.Time{}, "msg") // want "formatter: use assert\\.WithinRangef"
	assert.WithinRangef(t, time.Time{}, time.Time{}, time.Time{}, "msg")
	assert.WithinRangef(t, time.Time{}, time.Time{}, time.Time{}, "msg with arg", 42)
	assert.WithinRangef(t, time.Time{}, time.Time{}, time.Time{}, "msg with arg %d", 42)
	assert.YAMLEq(t, "", "")
	assert.YAMLEq(t, "", "", "msg") // want "formatter: use assert\\.YAMLEqf"
	assert.YAMLEqf(t, "", "", "msg")
	assert.YAMLEqf(t, "", "", "msg with arg", 42)
	assert.YAMLEqf(t, "", "", "msg with arg %d", 42)
	assert.Zero(t, nil)
	assert.Zero(t, nil, "msg") // want "formatter: use assert\\.Zerof"
	assert.Zerof(t, nil, "msg")
	assert.Zerof(t, nil, "msg with arg", 42)
	assert.Zerof(t, nil, "msg with arg %d", 42)
}
