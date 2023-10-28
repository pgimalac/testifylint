// Code generated by testifylint/internal/testgen. DO NOT EDIT.

package nilcompare

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNilCompareChecker(t *testing.T) {
	var value any
	var Row map[string]any

	// Invalid.
	{
		assert.Equal(t, value, nil)                                            // want "nil-compare: use assert\\.Nil"
		assert.Equalf(t, value, nil, "msg with args %d %s", 42, "42")          // want "nil-compare: use assert\\.Nilf"
		require.Equal(t, value, nil)                                           // want "nil-compare: use require\\.Nil"
		require.Equalf(t, value, nil, "msg with args %d %s", 42, "42")         // want "nil-compare: use require\\.Nilf"
		assert.Equal(t, nil, value)                                            // want "nil-compare: use assert\\.Nil"
		assert.Equalf(t, nil, value, "msg with args %d %s", 42, "42")          // want "nil-compare: use assert\\.Nilf"
		require.Equal(t, nil, value)                                           // want "nil-compare: use require\\.Nil"
		require.Equalf(t, nil, value, "msg with args %d %s", 42, "42")         // want "nil-compare: use require\\.Nilf"
		assert.Equal(t, Row["col"], nil)                                       // want "nil-compare: use assert\\.Nil"
		assert.Equalf(t, Row["col"], nil, "msg with args %d %s", 42, "42")     // want "nil-compare: use assert\\.Nilf"
		require.Equal(t, Row["col"], nil)                                      // want "nil-compare: use require\\.Nil"
		require.Equalf(t, Row["col"], nil, "msg with args %d %s", 42, "42")    // want "nil-compare: use require\\.Nilf"
		assert.Equal(t, nil, Row["col"])                                       // want "nil-compare: use assert\\.Nil"
		assert.Equalf(t, nil, Row["col"], "msg with args %d %s", 42, "42")     // want "nil-compare: use assert\\.Nilf"
		require.Equal(t, nil, Row["col"])                                      // want "nil-compare: use require\\.Nil"
		require.Equalf(t, nil, Row["col"], "msg with args %d %s", 42, "42")    // want "nil-compare: use require\\.Nilf"
		assert.NotEqual(t, value, nil)                                         // want "nil-compare: use assert\\.NotNil"
		assert.NotEqualf(t, value, nil, "msg with args %d %s", 42, "42")       // want "nil-compare: use assert\\.NotNilf"
		require.NotEqual(t, value, nil)                                        // want "nil-compare: use require\\.NotNil"
		require.NotEqualf(t, value, nil, "msg with args %d %s", 42, "42")      // want "nil-compare: use require\\.NotNilf"
		assert.NotEqual(t, nil, value)                                         // want "nil-compare: use assert\\.NotNil"
		assert.NotEqualf(t, nil, value, "msg with args %d %s", 42, "42")       // want "nil-compare: use assert\\.NotNilf"
		require.NotEqual(t, nil, value)                                        // want "nil-compare: use require\\.NotNil"
		require.NotEqualf(t, nil, value, "msg with args %d %s", 42, "42")      // want "nil-compare: use require\\.NotNilf"
		assert.NotEqual(t, Row["col"], nil)                                    // want "nil-compare: use assert\\.NotNil"
		assert.NotEqualf(t, Row["col"], nil, "msg with args %d %s", 42, "42")  // want "nil-compare: use assert\\.NotNilf"
		require.NotEqual(t, Row["col"], nil)                                   // want "nil-compare: use require\\.NotNil"
		require.NotEqualf(t, Row["col"], nil, "msg with args %d %s", 42, "42") // want "nil-compare: use require\\.NotNilf"
		assert.NotEqual(t, nil, Row["col"])                                    // want "nil-compare: use assert\\.NotNil"
		assert.NotEqualf(t, nil, Row["col"], "msg with args %d %s", 42, "42")  // want "nil-compare: use assert\\.NotNilf"
		require.NotEqual(t, nil, Row["col"])                                   // want "nil-compare: use require\\.NotNil"
		require.NotEqualf(t, nil, Row["col"], "msg with args %d %s", 42, "42") // want "nil-compare: use require\\.NotNilf"
	}

	// Valid.
	{
		assert.Nil(t, value)
		assert.Nilf(t, value, "msg with args %d %s", 42, "42")
		require.Nil(t, value)
		require.Nilf(t, value, "msg with args %d %s", 42, "42")
		assert.NotNil(t, value)
		assert.NotNilf(t, value, "msg with args %d %s", 42, "42")
		require.NotNil(t, value)
		require.NotNilf(t, value, "msg with args %d %s", 42, "42")
	}

	// Ignored.
	{
		assert.Equal(t, value, value)
		assert.Equalf(t, value, value, "msg with args %d %s", 42, "42")
		require.Equal(t, value, value)
		require.Equalf(t, value, value, "msg with args %d %s", 42, "42")
		assert.Equal(t, nil, nil)
		assert.Equalf(t, nil, nil, "msg with args %d %s", 42, "42")
		require.Equal(t, nil, nil)
		require.Equalf(t, nil, nil, "msg with args %d %s", 42, "42")
		assert.Equal(t, Row["col"], "foo")
		assert.Equalf(t, Row["col"], "foo", "msg with args %d %s", 42, "42")
		require.Equal(t, Row["col"], "foo")
		require.Equalf(t, Row["col"], "foo", "msg with args %d %s", 42, "42")
		assert.Equal(t, "foo", Row["col"])
		assert.Equalf(t, "foo", Row["col"], "msg with args %d %s", 42, "42")
		require.Equal(t, "foo", Row["col"])
		require.Equalf(t, "foo", Row["col"], "msg with args %d %s", 42, "42")
		assert.Equal(t, Row["col"], Row["col"])
		assert.Equalf(t, Row["col"], Row["col"], "msg with args %d %s", 42, "42")
		require.Equal(t, Row["col"], Row["col"])
		require.Equalf(t, Row["col"], Row["col"], "msg with args %d %s", 42, "42")
		assert.NotEqual(t, value, value)
		assert.NotEqualf(t, value, value, "msg with args %d %s", 42, "42")
		require.NotEqual(t, value, value)
		require.NotEqualf(t, value, value, "msg with args %d %s", 42, "42")
		assert.NotEqual(t, nil, nil)
		assert.NotEqualf(t, nil, nil, "msg with args %d %s", 42, "42")
		require.NotEqual(t, nil, nil)
		require.NotEqualf(t, nil, nil, "msg with args %d %s", 42, "42")
		assert.NotEqual(t, Row["col"], "foo")
		assert.NotEqualf(t, Row["col"], "foo", "msg with args %d %s", 42, "42")
		require.NotEqual(t, Row["col"], "foo")
		require.NotEqualf(t, Row["col"], "foo", "msg with args %d %s", 42, "42")
		assert.NotEqual(t, "foo", Row["col"])
		assert.NotEqualf(t, "foo", Row["col"], "msg with args %d %s", 42, "42")
		require.NotEqual(t, "foo", Row["col"])
		require.NotEqualf(t, "foo", Row["col"], "msg with args %d %s", 42, "42")
		assert.NotEqual(t, Row["col"], Row["col"])
		assert.NotEqualf(t, Row["col"], Row["col"], "msg with args %d %s", 42, "42")
		require.NotEqual(t, Row["col"], Row["col"])
		require.NotEqualf(t, Row["col"], Row["col"], "msg with args %d %s", 42, "42")
	}
}
