// Code generated by testifylint/internal/testgen. DO NOT EDIT.

package empty

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyChecker(t *testing.T) {
	var elems []string

	// assert.Empty cases.
	{
		// Invalid.
		assert.Len(t, elems, 0)                                                   // want "empty: use assert\\.Empty"
		assert.Lenf(t, elems, 0, "msg with args %d %s", 42, "42")                 // want "empty: use assert\\.Emptyf"
		assert.Equal(t, len(elems), 0)                                            // want "empty: use assert\\.Empty"
		assert.Equalf(t, len(elems), 0, "msg with args %d %s", 42, "42")          // want "empty: use assert\\.Emptyf"
		assert.Equal(t, 0, len(elems))                                            // want "empty: use assert\\.Empty"
		assert.Equalf(t, 0, len(elems), "msg with args %d %s", 42, "42")          // want "empty: use assert\\.Emptyf"
		assert.LessOrEqual(t, len(elems), 0)                                      // want "empty: use assert\\.Empty"
		assert.LessOrEqualf(t, len(elems), 0, "msg with args %d %s", 42, "42")    // want "empty: use assert\\.Emptyf"
		assert.GreaterOrEqual(t, 0, len(elems))                                   // want "empty: use assert\\.Empty"
		assert.GreaterOrEqualf(t, 0, len(elems), "msg with args %d %s", 42, "42") // want "empty: use assert\\.Emptyf"
		assert.Less(t, len(elems), 1)                                             // want "empty: use assert\\.Empty"
		assert.Lessf(t, len(elems), 1, "msg with args %d %s", 42, "42")           // want "empty: use assert\\.Emptyf"
		assert.Greater(t, 1, len(elems))                                          // want "empty: use assert\\.Empty"
		assert.Greaterf(t, 1, len(elems), "msg with args %d %s", 42, "42")        // want "empty: use assert\\.Emptyf"

		// Valid.
		assert.Empty(t, elems)
		assert.Emptyf(t, elems, "msg with args %d %s", 42, "42")
	}

	// assert.NotEmpty cases.
	{
		// Invalid.
		assert.NotEqual(t, len(elems), 0)                                   // want "empty: use assert\\.NotEmpty"
		assert.NotEqualf(t, len(elems), 0, "msg with args %d %s", 42, "42") // want "empty: use assert\\.NotEmptyf"
		assert.NotEqual(t, 0, len(elems))                                   // want "empty: use assert\\.NotEmpty"
		assert.NotEqualf(t, 0, len(elems), "msg with args %d %s", 42, "42") // want "empty: use assert\\.NotEmptyf"
		assert.Greater(t, len(elems), 0)                                    // want "empty: use assert\\.NotEmpty"
		assert.Greaterf(t, len(elems), 0, "msg with args %d %s", 42, "42")  // want "empty: use assert\\.NotEmptyf"
		assert.Less(t, 0, len(elems))                                       // want "empty: use assert\\.NotEmpty"
		assert.Lessf(t, 0, len(elems), "msg with args %d %s", 42, "42")     // want "empty: use assert\\.NotEmptyf"

		// Valid.
		assert.NotEmpty(t, elems)
		assert.NotEmptyf(t, elems, "msg with args %d %s", 42, "42")
	}
}

func TestEmptyChecker_LenVarIndependence(t *testing.T) {
	var (
		arr    [0]int
		arrPtr *[0]int
		sl     []int
		mp     map[int]int
		str    string
		ch     chan int
	)

	assert.Equal(t, 0, len(arr))    // want "empty: use assert\\.Empty"
	assert.Equal(t, 0, len(arrPtr)) // want "empty: use assert\\.Empty"
	assert.Equal(t, 0, len(sl))     // want "empty: use assert\\.Empty"
	assert.Equal(t, 0, len(mp))     // want "empty: use assert\\.Empty"
	assert.Equal(t, 0, len(str))    // want "empty: use assert\\.Empty"
	assert.Equal(t, 0, len(ch))     // want "empty: use assert\\.Empty"
}

func TestEmptyChecker_Ignored(t *testing.T) {
	var elems []string

	assert.Len(t, elems, len(elems))
	assert.Lenf(t, elems, len(elems), "msg with args %d %s", 42, "42")
	assert.Len(t, elems, 1)
	assert.Lenf(t, elems, 1, "msg with args %d %s", 42, "42")
	assert.Equal(t, len(elems), len(elems))
	assert.Equalf(t, len(elems), len(elems), "msg with args %d %s", 42, "42")
	assert.Equal(t, len(elems), 1)
	assert.Equalf(t, len(elems), 1, "msg with args %d %s", 42, "42")
	assert.Equal(t, 1, len(elems))
	assert.Equalf(t, 1, len(elems), "msg with args %d %s", 42, "42")
	assert.NotEqual(t, len(elems), len(elems))
	assert.NotEqualf(t, len(elems), len(elems), "msg with args %d %s", 42, "42")
	assert.NotEqual(t, len(elems), 1)
	assert.NotEqualf(t, len(elems), 1, "msg with args %d %s", 42, "42")
	assert.NotEqual(t, 1, len(elems))
	assert.NotEqualf(t, 1, len(elems), "msg with args %d %s", 42, "42")
	assert.Greater(t, len(elems), len(elems))
	assert.Greaterf(t, len(elems), len(elems), "msg with args %d %s", 42, "42")
	assert.Greater(t, len(elems), 2)
	assert.Greaterf(t, len(elems), 2, "msg with args %d %s", 42, "42")
	assert.Greater(t, 2, len(elems))
	assert.Greaterf(t, 2, len(elems), "msg with args %d %s", 42, "42")
	assert.GreaterOrEqual(t, len(elems), len(elems))
	assert.GreaterOrEqualf(t, len(elems), len(elems), "msg with args %d %s", 42, "42")
	assert.GreaterOrEqual(t, len(elems), 0)
	assert.GreaterOrEqualf(t, len(elems), 0, "msg with args %d %s", 42, "42")
	assert.GreaterOrEqual(t, len(elems), 2)
	assert.GreaterOrEqualf(t, len(elems), 2, "msg with args %d %s", 42, "42")
	assert.GreaterOrEqual(t, 2, len(elems))
	assert.GreaterOrEqualf(t, 2, len(elems), "msg with args %d %s", 42, "42")
	assert.Less(t, len(elems), len(elems))
	assert.Lessf(t, len(elems), len(elems), "msg with args %d %s", 42, "42")
	assert.Less(t, len(elems), 0)
	assert.Lessf(t, len(elems), 0, "msg with args %d %s", 42, "42")
	assert.Less(t, len(elems), 2)
	assert.Lessf(t, len(elems), 2, "msg with args %d %s", 42, "42")
	assert.Less(t, 2, len(elems))
	assert.Lessf(t, 2, len(elems), "msg with args %d %s", 42, "42")
	assert.LessOrEqual(t, len(elems), len(elems))
	assert.LessOrEqualf(t, len(elems), len(elems), "msg with args %d %s", 42, "42")
	assert.LessOrEqual(t, 0, len(elems))
	assert.LessOrEqualf(t, 0, len(elems), "msg with args %d %s", 42, "42")
	assert.LessOrEqual(t, len(elems), 2)
	assert.LessOrEqualf(t, len(elems), 2, "msg with args %d %s", 42, "42")
	assert.LessOrEqual(t, 2, len(elems))
	assert.LessOrEqualf(t, 2, len(elems), "msg with args %d %s", 42, "42")
	assert.Greater(t, len(elems), 1)
	assert.Greaterf(t, len(elems), 1, "msg with args %d %s", 42, "42")
	assert.Less(t, 1, len(elems))
	assert.Lessf(t, 1, len(elems), "msg with args %d %s", 42, "42")
	assert.GreaterOrEqual(t, len(elems), 1)
	assert.GreaterOrEqualf(t, len(elems), 1, "msg with args %d %s", 42, "42")
	assert.LessOrEqual(t, 1, len(elems))
	assert.LessOrEqualf(t, 1, len(elems), "msg with args %d %s", 42, "42")
}
