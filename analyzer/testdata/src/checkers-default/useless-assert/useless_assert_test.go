// Code generated by testifylint/internal/testgen. DO NOT EDIT.
package uselessassert

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUselessAssertChecker(t *testing.T) {
	var value any
	var err error
	var elapsed time.Time
	var str string
	var tc testCase

	// Invalid.
	{
		assert.Equal(t, 42, 42)                                                                           // want "useless-assert: asserting of the same variable"
		assert.Equal(t, 42, 42, "msg")                                                                    // want "useless-assert: asserting of the same variable"
		assert.Equal(t, 42, 42, "msg with arg %d", 42)                                                    // want "useless-assert: asserting of the same variable"
		assert.Equal(t, 42, 42, "msg with args %d %s", 42, "42")                                          // want "useless-assert: asserting of the same variable"
		assert.Equalf(t, 42, 42, "msg")                                                                   // want "useless-assert: asserting of the same variable"
		assert.Equalf(t, 42, 42, "msg with arg %d", 42)                                                   // want "useless-assert: asserting of the same variable"
		assert.Equalf(t, 42, 42, "msg with args %d %s", 42, "42")                                         // want "useless-assert: asserting of the same variable"
		assert.Equal(t, "value", "value")                                                                 // want "useless-assert: asserting of the same variable"
		assert.Equal(t, "value", "value", "msg")                                                          // want "useless-assert: asserting of the same variable"
		assert.Equal(t, "value", "value", "msg with arg %d", 42)                                          // want "useless-assert: asserting of the same variable"
		assert.Equal(t, "value", "value", "msg with args %d %s", 42, "42")                                // want "useless-assert: asserting of the same variable"
		assert.Equalf(t, "value", "value", "msg")                                                         // want "useless-assert: asserting of the same variable"
		assert.Equalf(t, "value", "value", "msg with arg %d", 42)                                         // want "useless-assert: asserting of the same variable"
		assert.Equalf(t, "value", "value", "msg with args %d %s", 42, "42")                               // want "useless-assert: asserting of the same variable"
		assert.Equal(t, value, value)                                                                     // want "useless-assert: asserting of the same variable"
		assert.Equal(t, value, value, "msg")                                                              // want "useless-assert: asserting of the same variable"
		assert.Equal(t, value, value, "msg with arg %d", 42)                                              // want "useless-assert: asserting of the same variable"
		assert.Equal(t, value, value, "msg with args %d %s", 42, "42")                                    // want "useless-assert: asserting of the same variable"
		assert.Equalf(t, value, value, "msg")                                                             // want "useless-assert: asserting of the same variable"
		assert.Equalf(t, value, value, "msg with arg %d", 42)                                             // want "useless-assert: asserting of the same variable"
		assert.Equalf(t, value, value, "msg with args %d %s", 42, "42")                                   // want "useless-assert: asserting of the same variable"
		assert.Equal(t, tc.A(), tc.A())                                                                   // want "useless-assert: asserting of the same variable"
		assert.Equal(t, tc.A(), tc.A(), "msg")                                                            // want "useless-assert: asserting of the same variable"
		assert.Equal(t, tc.A(), tc.A(), "msg with arg %d", 42)                                            // want "useless-assert: asserting of the same variable"
		assert.Equal(t, tc.A(), tc.A(), "msg with args %d %s", 42, "42")                                  // want "useless-assert: asserting of the same variable"
		assert.Equalf(t, tc.A(), tc.A(), "msg")                                                           // want "useless-assert: asserting of the same variable"
		assert.Equalf(t, tc.A(), tc.A(), "msg with arg %d", 42)                                           // want "useless-assert: asserting of the same variable"
		assert.Equalf(t, tc.A(), tc.A(), "msg with args %d %s", 42, "42")                                 // want "useless-assert: asserting of the same variable"
		assert.Equal(t, testCase{}.A().B().C(), testCase{}.A().B().C())                                   // want "useless-assert: asserting of the same variable"
		assert.Equal(t, testCase{}.A().B().C(), testCase{}.A().B().C(), "msg")                            // want "useless-assert: asserting of the same variable"
		assert.Equal(t, testCase{}.A().B().C(), testCase{}.A().B().C(), "msg with arg %d", 42)            // want "useless-assert: asserting of the same variable"
		assert.Equal(t, testCase{}.A().B().C(), testCase{}.A().B().C(), "msg with args %d %s", 42, "42")  // want "useless-assert: asserting of the same variable"
		assert.Equalf(t, testCase{}.A().B().C(), testCase{}.A().B().C(), "msg")                           // want "useless-assert: asserting of the same variable"
		assert.Equalf(t, testCase{}.A().B().C(), testCase{}.A().B().C(), "msg with arg %d", 42)           // want "useless-assert: asserting of the same variable"
		assert.Equalf(t, testCase{}.A().B().C(), testCase{}.A().B().C(), "msg with args %d %s", 42, "42") // want "useless-assert: asserting of the same variable"
		assert.IsType(t, (*testCase)(nil), (*testCase)(nil))                                              // want "useless-assert: asserting of the same variable"
		assert.IsType(t, (*testCase)(nil), (*testCase)(nil), "msg")                                       // want "useless-assert: asserting of the same variable"
		assert.IsType(t, (*testCase)(nil), (*testCase)(nil), "msg with arg %d", 42)                       // want "useless-assert: asserting of the same variable"
		assert.IsType(t, (*testCase)(nil), (*testCase)(nil), "msg with args %d %s", 42, "42")             // want "useless-assert: asserting of the same variable"
		assert.IsTypef(t, (*testCase)(nil), (*testCase)(nil), "msg")                                      // want "useless-assert: asserting of the same variable"
		assert.IsTypef(t, (*testCase)(nil), (*testCase)(nil), "msg with arg %d", 42)                      // want "useless-assert: asserting of the same variable"
		assert.IsTypef(t, (*testCase)(nil), (*testCase)(nil), "msg with args %d %s", 42, "42")            // want "useless-assert: asserting of the same variable"
		assert.GreaterOrEqual(t, value, value)                                                            // want "useless-assert: asserting of the same variable"
		assert.GreaterOrEqualf(t, value, value, "msg with args %d %s", 42, "42")                          // want "useless-assert: asserting of the same variable"
		assert.Less(t, value, value)                                                                      // want "useless-assert: asserting of the same variable"
		assert.Lessf(t, value, value, "msg with args %d %s", 42, "42")                                    // want "useless-assert: asserting of the same variable"
		assert.NotRegexp(t, value, value)                                                                 // want "useless-assert: asserting of the same variable"
		assert.NotRegexpf(t, value, value, "msg with args %d %s", 42, "42")                               // want "useless-assert: asserting of the same variable"
		assert.NotSame(t, value, value)                                                                   // want "useless-assert: asserting of the same variable"
		assert.NotSamef(t, value, value, "msg with args %d %s", 42, "42")                                 // want "useless-assert: asserting of the same variable"
		assert.ElementsMatch(t, value, value)                                                             // want "useless-assert: asserting of the same variable"
		assert.ElementsMatchf(t, value, value, "msg with args %d %s", 42, "42")                           // want "useless-assert: asserting of the same variable"
		assert.EqualExportedValues(t, value, value)                                                       // want "useless-assert: asserting of the same variable"
		assert.EqualExportedValuesf(t, value, value, "msg with args %d %s", 42, "42")                     // want "useless-assert: asserting of the same variable"
		assert.EqualValues(t, value, value)                                                               // want "useless-assert: asserting of the same variable"
		assert.EqualValuesf(t, value, value, "msg with args %d %s", 42, "42")                             // want "useless-assert: asserting of the same variable"
		assert.Same(t, value, value)                                                                      // want "useless-assert: asserting of the same variable"
		assert.Samef(t, value, value, "msg with args %d %s", 42, "42")                                    // want "useless-assert: asserting of the same variable"
		assert.Equal(t, value, value)                                                                     // want "useless-assert: asserting of the same variable"
		assert.Equalf(t, value, value, "msg with args %d %s", 42, "42")                                   // want "useless-assert: asserting of the same variable"
		assert.InDeltaMapValues(t, value, value, 0.01)                                                    // want "useless-assert: asserting of the same variable"
		assert.InDeltaMapValuesf(t, value, value, 0.01, "msg with args %d %s", 42, "42")                  // want "useless-assert: asserting of the same variable"
		assert.NotEqualValues(t, value, value)                                                            // want "useless-assert: asserting of the same variable"
		assert.NotEqualValuesf(t, value, value, "msg with args %d %s", 42, "42")                          // want "useless-assert: asserting of the same variable"
		assert.NotSubset(t, value, value)                                                                 // want "useless-assert: asserting of the same variable"
		assert.NotSubsetf(t, value, value, "msg with args %d %s", 42, "42")                               // want "useless-assert: asserting of the same variable"
		assert.InEpsilonSlice(t, value, value, 0.0001)                                                    // want "useless-assert: asserting of the same variable"
		assert.InEpsilonSlicef(t, value, value, 0.0001, "msg with args %d %s", 42, "42")                  // want "useless-assert: asserting of the same variable"
		assert.WithinDuration(t, elapsed, elapsed, time.Second)                                           // want "useless-assert: asserting of the same variable"
		assert.WithinDurationf(t, elapsed, elapsed, time.Second, "msg with args %d %s", 42, "42")         // want "useless-assert: asserting of the same variable"
		assert.InDelta(t, value, value, 0.01)                                                             // want "useless-assert: asserting of the same variable"
		assert.InDeltaf(t, value, value, 0.01, "msg with args %d %s", 42, "42")                           // want "useless-assert: asserting of the same variable"
		assert.JSONEq(t, str, str)                                                                        // want "useless-assert: asserting of the same variable"
		assert.JSONEqf(t, str, str, "msg with args %d %s", 42, "42")                                      // want "useless-assert: asserting of the same variable"
		assert.Regexp(t, value, value)                                                                    // want "useless-assert: asserting of the same variable"
		assert.Regexpf(t, value, value, "msg with args %d %s", 42, "42")                                  // want "useless-assert: asserting of the same variable"
		assert.Exactly(t, value, value)                                                                   // want "useless-assert: asserting of the same variable"
		assert.Exactlyf(t, value, value, "msg with args %d %s", 42, "42")                                 // want "useless-assert: asserting of the same variable"
		assert.Greater(t, value, value)                                                                   // want "useless-assert: asserting of the same variable"
		assert.Greaterf(t, value, value, "msg with args %d %s", 42, "42")                                 // want "useless-assert: asserting of the same variable"
		assert.Implements(t, value, value)                                                                // want "useless-assert: asserting of the same variable"
		assert.Implementsf(t, value, value, "msg with args %d %s", 42, "42")                              // want "useless-assert: asserting of the same variable"
		assert.InEpsilon(t, value, value, 0.0001)                                                         // want "useless-assert: asserting of the same variable"
		assert.InEpsilonf(t, value, value, 0.0001, "msg with args %d %s", 42, "42")                       // want "useless-assert: asserting of the same variable"
		assert.LessOrEqual(t, value, value)                                                               // want "useless-assert: asserting of the same variable"
		assert.LessOrEqualf(t, value, value, "msg with args %d %s", 42, "42")                             // want "useless-assert: asserting of the same variable"
		assert.YAMLEq(t, str, str)                                                                        // want "useless-assert: asserting of the same variable"
		assert.YAMLEqf(t, str, str, "msg with args %d %s", 42, "42")                                      // want "useless-assert: asserting of the same variable"
		assert.ErrorAs(t, err, err)                                                                       // want "useless-assert: asserting of the same variable"
		assert.ErrorAsf(t, err, err, "msg with args %d %s", 42, "42")                                     // want "useless-assert: asserting of the same variable"
		assert.ErrorIs(t, err, err)                                                                       // want "useless-assert: asserting of the same variable"
		assert.ErrorIsf(t, err, err, "msg with args %d %s", 42, "42")                                     // want "useless-assert: asserting of the same variable"
		assert.InDeltaSlice(t, value, value, 0.01)                                                        // want "useless-assert: asserting of the same variable"
		assert.InDeltaSlicef(t, value, value, 0.01, "msg with args %d %s", 42, "42")                      // want "useless-assert: asserting of the same variable"
		assert.IsType(t, value, value)                                                                    // want "useless-assert: asserting of the same variable"
		assert.IsTypef(t, value, value, "msg with args %d %s", 42, "42")                                  // want "useless-assert: asserting of the same variable"
		assert.Contains(t, value, value)                                                                  // want "useless-assert: asserting of the same variable"
		assert.Containsf(t, value, value, "msg with args %d %s", 42, "42")                                // want "useless-assert: asserting of the same variable"
		assert.NotEqual(t, value, value)                                                                  // want "useless-assert: asserting of the same variable"
		assert.NotEqualf(t, value, value, "msg with args %d %s", 42, "42")                                // want "useless-assert: asserting of the same variable"
		assert.NotErrorIs(t, err, err)                                                                    // want "useless-assert: asserting of the same variable"
		assert.NotErrorIsf(t, err, err, "msg with args %d %s", 42, "42")                                  // want "useless-assert: asserting of the same variable"
		assert.Subset(t, value, value)                                                                    // want "useless-assert: asserting of the same variable"
		assert.Subsetf(t, value, value, "msg with args %d %s", 42, "42")                                  // want "useless-assert: asserting of the same variable"
	}

	// Valid.
	{
		assert.Equal(t, value, 42)
		assert.Equal(t, value, 42, "msg")
		assert.Equal(t, value, 42, "msg with arg %d", 42)
		assert.Equal(t, value, 42, "msg with args %d %s", 42, "42")
		assert.Equalf(t, value, 42, "msg")
		assert.Equalf(t, value, 42, "msg with arg %d", 42)
		assert.Equalf(t, value, 42, "msg with args %d %s", 42, "42")
		assert.Equal(t, value, "value")
		assert.Equal(t, value, "value", "msg")
		assert.Equal(t, value, "value", "msg with arg %d", 42)
		assert.Equal(t, value, "value", "msg with args %d %s", 42, "42")
		assert.Equalf(t, value, "value", "msg")
		assert.Equalf(t, value, "value", "msg with arg %d", 42)
		assert.Equalf(t, value, "value", "msg with args %d %s", 42, "42")
		assert.Equal(t, tc.A(), "tc.A()")
		assert.Equal(t, tc.A(), "tc.A()", "msg")
		assert.Equal(t, tc.A(), "tc.A()", "msg with arg %d", 42)
		assert.Equal(t, tc.A(), "tc.A()", "msg with args %d %s", 42, "42")
		assert.Equalf(t, tc.A(), "tc.A()", "msg")
		assert.Equalf(t, tc.A(), "tc.A()", "msg with arg %d", 42)
		assert.Equalf(t, tc.A(), "tc.A()", "msg with args %d %s", 42, "42")
		assert.Equal(t, testCase{}.A().B().C(), tc.A().B().C())
		assert.Equal(t, testCase{}.A().B().C(), tc.A().B().C(), "msg")
		assert.Equal(t, testCase{}.A().B().C(), tc.A().B().C(), "msg with arg %d", 42)
		assert.Equal(t, testCase{}.A().B().C(), tc.A().B().C(), "msg with args %d %s", 42, "42")
		assert.Equalf(t, testCase{}.A().B().C(), tc.A().B().C(), "msg")
		assert.Equalf(t, testCase{}.A().B().C(), tc.A().B().C(), "msg with arg %d", 42)
		assert.Equalf(t, testCase{}.A().B().C(), tc.A().B().C(), "msg with args %d %s", 42, "42")
		assert.IsType(t, tc, testCase{})
		assert.IsType(t, tc, testCase{}, "msg")
		assert.IsType(t, tc, testCase{}, "msg with arg %d", 42)
		assert.IsType(t, tc, testCase{}, "msg with args %d %s", 42, "42")
		assert.IsTypef(t, tc, testCase{}, "msg")
		assert.IsTypef(t, tc, testCase{}, "msg with arg %d", 42)
		assert.IsTypef(t, tc, testCase{}, "msg with args %d %s", 42, "42")
	}
}

type testCase struct{}

func (testCase) A() testCase { return testCase{} }
func (testCase) B() testCase { return testCase{} }
func (testCase) C() testCase { return testCase{} }
