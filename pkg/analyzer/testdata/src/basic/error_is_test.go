// Code generated by testifylint/internal/cmd/testgen. DO NOT EDIT.

package basic

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestErrorInsteadOfErrorIs(t *testing.T) {
	var errSentinel = errors.New("user not found")
	var err error

	t.Run("assert", func(t *testing.T) {
		{
			assert.Error(t, err, errSentinel)                        // want "error-is: invalid usage of assert\\.Error, use assert\\.ErrorIs instead"
			assert.Error(t, err, errSentinel, "msg")                 // want "error-is: invalid usage of assert\\.Error, use assert\\.ErrorIs instead"
			assert.Error(t, err, errSentinel, "msg with arg %d", 42) // want "error-is: invalid usage of assert\\.Error, use assert\\.ErrorIs instead"

			assert.NoError(t, err, errSentinel)                        // want "error-is: invalid usage of assert\\.NoError, use assert\\.NotErrorIs instead"
			assert.NoError(t, err, errSentinel, "msg")                 // want "error-is: invalid usage of assert\\.NoError, use assert\\.NotErrorIs instead"
			assert.NoError(t, err, errSentinel, "msg with arg %d", 42) // want "error-is: invalid usage of assert\\.NoError, use assert\\.NotErrorIs instead"
		}

		// Valid.

		{
			assert.Error(t, err)
			assert.Error(t, err, "msg")
			assert.Error(t, err, "msg with arg %d", 42)
			assert.Errorf(t, err, "msg")
			assert.Errorf(t, err, "msg with arg %d", 42)

			assert.ErrorIs(t, err, errSentinel)
			assert.ErrorIs(t, err, errSentinel, "msg")
			assert.ErrorIs(t, err, errSentinel, "msg with arg %d", 42)
			assert.ErrorIsf(t, err, errSentinel, "msg")
			assert.ErrorIsf(t, err, errSentinel, "msg with arg %d", 42)

			assert.NoError(t, err)
			assert.NoError(t, err, "msg")
			assert.NoError(t, err, "msg with arg %d", 42)
			assert.NoErrorf(t, err, "msg")
			assert.NoErrorf(t, err, "msg with arg %d", 42)

			assert.NotErrorIs(t, err, errSentinel)
			assert.NotErrorIs(t, err, errSentinel, "msg")
			assert.NotErrorIs(t, err, errSentinel, "msg with arg %d", 42)
			assert.NotErrorIsf(t, err, errSentinel, "msg")
			assert.NotErrorIsf(t, err, errSentinel, "msg with arg %d", 42)
		}
	})

	t.Run("require", func(t *testing.T) {
		{
			require.Error(t, err, errSentinel)                        // want "error-is: invalid usage of require\\.Error, use require\\.ErrorIs instead"
			require.Error(t, err, errSentinel, "msg")                 // want "error-is: invalid usage of require\\.Error, use require\\.ErrorIs instead"
			require.Error(t, err, errSentinel, "msg with arg %d", 42) // want "error-is: invalid usage of require\\.Error, use require\\.ErrorIs instead"

			require.NoError(t, err, errSentinel)                        // want "error-is: invalid usage of require\\.NoError, use require\\.NotErrorIs instead"
			require.NoError(t, err, errSentinel, "msg")                 // want "error-is: invalid usage of require\\.NoError, use require\\.NotErrorIs instead"
			require.NoError(t, err, errSentinel, "msg with arg %d", 42) // want "error-is: invalid usage of require\\.NoError, use require\\.NotErrorIs instead"
		}

		// Valid.

		{
			require.Error(t, err)
			require.Error(t, err, "msg")
			require.Error(t, err, "msg with arg %d", 42)
			require.Errorf(t, err, "msg")
			require.Errorf(t, err, "msg with arg %d", 42)

			require.ErrorIs(t, err, errSentinel)
			require.ErrorIs(t, err, errSentinel, "msg")
			require.ErrorIs(t, err, errSentinel, "msg with arg %d", 42)
			require.ErrorIsf(t, err, errSentinel, "msg")
			require.ErrorIsf(t, err, errSentinel, "msg with arg %d", 42)

			require.NoError(t, err)
			require.NoError(t, err, "msg")
			require.NoError(t, err, "msg with arg %d", 42)
			require.NoErrorf(t, err, "msg")
			require.NoErrorf(t, err, "msg with arg %d", 42)

			require.NotErrorIs(t, err, errSentinel)
			require.NotErrorIs(t, err, errSentinel, "msg")
			require.NotErrorIs(t, err, errSentinel, "msg with arg %d", 42)
			require.NotErrorIsf(t, err, errSentinel, "msg")
			require.NotErrorIsf(t, err, errSentinel, "msg with arg %d", 42)
		}
	})

	assObj, reqObj := assert.New(t), require.New(t)

	t.Run("assObj", func(t *testing.T) {
		{
			assObj.Error(err, errSentinel)                        // want "error-is: invalid usage of assObj\\.Error, use assObj\\.ErrorIs instead"
			assObj.Error(err, errSentinel, "msg")                 // want "error-is: invalid usage of assObj\\.Error, use assObj\\.ErrorIs instead"
			assObj.Error(err, errSentinel, "msg with arg %d", 42) // want "error-is: invalid usage of assObj\\.Error, use assObj\\.ErrorIs instead"

			assObj.NoError(err, errSentinel)                        // want "error-is: invalid usage of assObj\\.NoError, use assObj\\.NotErrorIs instead"
			assObj.NoError(err, errSentinel, "msg")                 // want "error-is: invalid usage of assObj\\.NoError, use assObj\\.NotErrorIs instead"
			assObj.NoError(err, errSentinel, "msg with arg %d", 42) // want "error-is: invalid usage of assObj\\.NoError, use assObj\\.NotErrorIs instead"
		}

		// Valid.

		{
			assObj.Error(err)
			assObj.Error(err, "msg")
			assObj.Error(err, "msg with arg %d", 42)
			assObj.Errorf(err, "msg")
			assObj.Errorf(err, "msg with arg %d", 42)

			assObj.ErrorIs(err, errSentinel)
			assObj.ErrorIs(err, errSentinel, "msg")
			assObj.ErrorIs(err, errSentinel, "msg with arg %d", 42)
			assObj.ErrorIsf(err, errSentinel, "msg")
			assObj.ErrorIsf(err, errSentinel, "msg with arg %d", 42)

			assObj.NoError(err)
			assObj.NoError(err, "msg")
			assObj.NoError(err, "msg with arg %d", 42)
			assObj.NoErrorf(err, "msg")
			assObj.NoErrorf(err, "msg with arg %d", 42)

			assObj.NotErrorIs(err, errSentinel)
			assObj.NotErrorIs(err, errSentinel, "msg")
			assObj.NotErrorIs(err, errSentinel, "msg with arg %d", 42)
			assObj.NotErrorIsf(err, errSentinel, "msg")
			assObj.NotErrorIsf(err, errSentinel, "msg with arg %d", 42)
		}
	})

	t.Run("reqObj", func(t *testing.T) {
		{
			reqObj.Error(err, errSentinel)                        // want "error-is: invalid usage of reqObj\\.Error, use reqObj\\.ErrorIs instead"
			reqObj.Error(err, errSentinel, "msg")                 // want "error-is: invalid usage of reqObj\\.Error, use reqObj\\.ErrorIs instead"
			reqObj.Error(err, errSentinel, "msg with arg %d", 42) // want "error-is: invalid usage of reqObj\\.Error, use reqObj\\.ErrorIs instead"

			reqObj.NoError(err, errSentinel)                        // want "error-is: invalid usage of reqObj\\.NoError, use reqObj\\.NotErrorIs instead"
			reqObj.NoError(err, errSentinel, "msg")                 // want "error-is: invalid usage of reqObj\\.NoError, use reqObj\\.NotErrorIs instead"
			reqObj.NoError(err, errSentinel, "msg with arg %d", 42) // want "error-is: invalid usage of reqObj\\.NoError, use reqObj\\.NotErrorIs instead"
		}

		// Valid.

		{
			reqObj.Error(err)
			reqObj.Error(err, "msg")
			reqObj.Error(err, "msg with arg %d", 42)
			reqObj.Errorf(err, "msg")
			reqObj.Errorf(err, "msg with arg %d", 42)

			reqObj.ErrorIs(err, errSentinel)
			reqObj.ErrorIs(err, errSentinel, "msg")
			reqObj.ErrorIs(err, errSentinel, "msg with arg %d", 42)
			reqObj.ErrorIsf(err, errSentinel, "msg")
			reqObj.ErrorIsf(err, errSentinel, "msg with arg %d", 42)

			reqObj.NoError(err)
			reqObj.NoError(err, "msg")
			reqObj.NoError(err, "msg with arg %d", 42)
			reqObj.NoErrorf(err, "msg")
			reqObj.NoErrorf(err, "msg with arg %d", 42)

			reqObj.NotErrorIs(err, errSentinel)
			reqObj.NotErrorIs(err, errSentinel, "msg")
			reqObj.NotErrorIs(err, errSentinel, "msg with arg %d", 42)
			reqObj.NotErrorIsf(err, errSentinel, "msg")
			reqObj.NotErrorIsf(err, errSentinel, "msg with arg %d", 42)
		}
	})
}

type ErrorInsteadOfErrorIsSuite struct {
	suite.Suite
}

func TestErrorInsteadOfErrorIsSuite(t *testing.T) {
	suite.Run(t, new(ErrorInsteadOfErrorIsSuite))
}

func (s *ErrorInsteadOfErrorIsSuite) TestAll() {
	var errSentinel = errors.New("user not found")
	var err error

	assObj, reqObj := s.Assert(), s.Require()

	{
		{
			s.Error(err, errSentinel)                        // want "error-is: invalid usage of s\\.Error, use s\\.ErrorIs instead"
			s.Error(err, errSentinel, "msg")                 // want "error-is: invalid usage of s\\.Error, use s\\.ErrorIs instead"
			s.Error(err, errSentinel, "msg with arg %d", 42) // want "error-is: invalid usage of s\\.Error, use s\\.ErrorIs instead"

			s.NoError(err, errSentinel)                        // want "error-is: invalid usage of s\\.NoError, use s\\.NotErrorIs instead"
			s.NoError(err, errSentinel, "msg")                 // want "error-is: invalid usage of s\\.NoError, use s\\.NotErrorIs instead"
			s.NoError(err, errSentinel, "msg with arg %d", 42) // want "error-is: invalid usage of s\\.NoError, use s\\.NotErrorIs instead"
		}

		// Valid.

		{
			s.Error(err)
			s.Error(err, "msg")
			s.Error(err, "msg with arg %d", 42)
			s.Errorf(err, "msg")
			s.Errorf(err, "msg with arg %d", 42)

			s.ErrorIs(err, errSentinel)
			s.ErrorIs(err, errSentinel, "msg")
			s.ErrorIs(err, errSentinel, "msg with arg %d", 42)
			s.ErrorIsf(err, errSentinel, "msg")
			s.ErrorIsf(err, errSentinel, "msg with arg %d", 42)

			s.NoError(err)
			s.NoError(err, "msg")
			s.NoError(err, "msg with arg %d", 42)
			s.NoErrorf(err, "msg")
			s.NoErrorf(err, "msg with arg %d", 42)

			s.NotErrorIs(err, errSentinel)
			s.NotErrorIs(err, errSentinel, "msg")
			s.NotErrorIs(err, errSentinel, "msg with arg %d", 42)
			s.NotErrorIsf(err, errSentinel, "msg")
			s.NotErrorIsf(err, errSentinel, "msg with arg %d", 42)
		}
	}

	{
		{
			s.Assert().Error(err, errSentinel)                        // want "error-is: invalid usage of s\\.Assert\\(\\)\\.Error, use s\\.Assert\\(\\)\\.ErrorIs instead"
			s.Assert().Error(err, errSentinel, "msg")                 // want "error-is: invalid usage of s\\.Assert\\(\\)\\.Error, use s\\.Assert\\(\\)\\.ErrorIs instead"
			s.Assert().Error(err, errSentinel, "msg with arg %d", 42) // want "error-is: invalid usage of s\\.Assert\\(\\)\\.Error, use s\\.Assert\\(\\)\\.ErrorIs instead"

			s.Assert().NoError(err, errSentinel)                        // want "error-is: invalid usage of s\\.Assert\\(\\)\\.NoError, use s\\.Assert\\(\\)\\.NotErrorIs instead"
			s.Assert().NoError(err, errSentinel, "msg")                 // want "error-is: invalid usage of s\\.Assert\\(\\)\\.NoError, use s\\.Assert\\(\\)\\.NotErrorIs instead"
			s.Assert().NoError(err, errSentinel, "msg with arg %d", 42) // want "error-is: invalid usage of s\\.Assert\\(\\)\\.NoError, use s\\.Assert\\(\\)\\.NotErrorIs instead"
		}

		// Valid.

		{
			s.Assert().Error(err)
			s.Assert().Error(err, "msg")
			s.Assert().Error(err, "msg with arg %d", 42)
			s.Assert().Errorf(err, "msg")
			s.Assert().Errorf(err, "msg with arg %d", 42)

			s.Assert().ErrorIs(err, errSentinel)
			s.Assert().ErrorIs(err, errSentinel, "msg")
			s.Assert().ErrorIs(err, errSentinel, "msg with arg %d", 42)
			s.Assert().ErrorIsf(err, errSentinel, "msg")
			s.Assert().ErrorIsf(err, errSentinel, "msg with arg %d", 42)

			s.Assert().NoError(err)
			s.Assert().NoError(err, "msg")
			s.Assert().NoError(err, "msg with arg %d", 42)
			s.Assert().NoErrorf(err, "msg")
			s.Assert().NoErrorf(err, "msg with arg %d", 42)

			s.Assert().NotErrorIs(err, errSentinel)
			s.Assert().NotErrorIs(err, errSentinel, "msg")
			s.Assert().NotErrorIs(err, errSentinel, "msg with arg %d", 42)
			s.Assert().NotErrorIsf(err, errSentinel, "msg")
			s.Assert().NotErrorIsf(err, errSentinel, "msg with arg %d", 42)
		}
	}

	{
		{
			assObj.Error(err, errSentinel)                        // want "error-is: invalid usage of assObj\\.Error, use assObj\\.ErrorIs instead"
			assObj.Error(err, errSentinel, "msg")                 // want "error-is: invalid usage of assObj\\.Error, use assObj\\.ErrorIs instead"
			assObj.Error(err, errSentinel, "msg with arg %d", 42) // want "error-is: invalid usage of assObj\\.Error, use assObj\\.ErrorIs instead"

			assObj.NoError(err, errSentinel)                        // want "error-is: invalid usage of assObj\\.NoError, use assObj\\.NotErrorIs instead"
			assObj.NoError(err, errSentinel, "msg")                 // want "error-is: invalid usage of assObj\\.NoError, use assObj\\.NotErrorIs instead"
			assObj.NoError(err, errSentinel, "msg with arg %d", 42) // want "error-is: invalid usage of assObj\\.NoError, use assObj\\.NotErrorIs instead"
		}

		// Valid.

		{
			assObj.Error(err)
			assObj.Error(err, "msg")
			assObj.Error(err, "msg with arg %d", 42)
			assObj.Errorf(err, "msg")
			assObj.Errorf(err, "msg with arg %d", 42)

			assObj.ErrorIs(err, errSentinel)
			assObj.ErrorIs(err, errSentinel, "msg")
			assObj.ErrorIs(err, errSentinel, "msg with arg %d", 42)
			assObj.ErrorIsf(err, errSentinel, "msg")
			assObj.ErrorIsf(err, errSentinel, "msg with arg %d", 42)

			assObj.NoError(err)
			assObj.NoError(err, "msg")
			assObj.NoError(err, "msg with arg %d", 42)
			assObj.NoErrorf(err, "msg")
			assObj.NoErrorf(err, "msg with arg %d", 42)

			assObj.NotErrorIs(err, errSentinel)
			assObj.NotErrorIs(err, errSentinel, "msg")
			assObj.NotErrorIs(err, errSentinel, "msg with arg %d", 42)
			assObj.NotErrorIsf(err, errSentinel, "msg")
			assObj.NotErrorIsf(err, errSentinel, "msg with arg %d", 42)
		}
	}

	{
		{
			s.Require().Error(err, errSentinel)                        // want "error-is: invalid usage of s\\.Require\\(\\)\\.Error, use s\\.Require\\(\\)\\.ErrorIs instead"
			s.Require().Error(err, errSentinel, "msg")                 // want "error-is: invalid usage of s\\.Require\\(\\)\\.Error, use s\\.Require\\(\\)\\.ErrorIs instead"
			s.Require().Error(err, errSentinel, "msg with arg %d", 42) // want "error-is: invalid usage of s\\.Require\\(\\)\\.Error, use s\\.Require\\(\\)\\.ErrorIs instead"

			s.Require().NoError(err, errSentinel)                        // want "error-is: invalid usage of s\\.Require\\(\\)\\.NoError, use s\\.Require\\(\\)\\.NotErrorIs instead"
			s.Require().NoError(err, errSentinel, "msg")                 // want "error-is: invalid usage of s\\.Require\\(\\)\\.NoError, use s\\.Require\\(\\)\\.NotErrorIs instead"
			s.Require().NoError(err, errSentinel, "msg with arg %d", 42) // want "error-is: invalid usage of s\\.Require\\(\\)\\.NoError, use s\\.Require\\(\\)\\.NotErrorIs instead"
		}

		// Valid.

		{
			s.Require().Error(err)
			s.Require().Error(err, "msg")
			s.Require().Error(err, "msg with arg %d", 42)
			s.Require().Errorf(err, "msg")
			s.Require().Errorf(err, "msg with arg %d", 42)

			s.Require().ErrorIs(err, errSentinel)
			s.Require().ErrorIs(err, errSentinel, "msg")
			s.Require().ErrorIs(err, errSentinel, "msg with arg %d", 42)
			s.Require().ErrorIsf(err, errSentinel, "msg")
			s.Require().ErrorIsf(err, errSentinel, "msg with arg %d", 42)

			s.Require().NoError(err)
			s.Require().NoError(err, "msg")
			s.Require().NoError(err, "msg with arg %d", 42)
			s.Require().NoErrorf(err, "msg")
			s.Require().NoErrorf(err, "msg with arg %d", 42)

			s.Require().NotErrorIs(err, errSentinel)
			s.Require().NotErrorIs(err, errSentinel, "msg")
			s.Require().NotErrorIs(err, errSentinel, "msg with arg %d", 42)
			s.Require().NotErrorIsf(err, errSentinel, "msg")
			s.Require().NotErrorIsf(err, errSentinel, "msg with arg %d", 42)
		}
	}

	{
		{
			reqObj.Error(err, errSentinel)                        // want "error-is: invalid usage of reqObj\\.Error, use reqObj\\.ErrorIs instead"
			reqObj.Error(err, errSentinel, "msg")                 // want "error-is: invalid usage of reqObj\\.Error, use reqObj\\.ErrorIs instead"
			reqObj.Error(err, errSentinel, "msg with arg %d", 42) // want "error-is: invalid usage of reqObj\\.Error, use reqObj\\.ErrorIs instead"

			reqObj.NoError(err, errSentinel)                        // want "error-is: invalid usage of reqObj\\.NoError, use reqObj\\.NotErrorIs instead"
			reqObj.NoError(err, errSentinel, "msg")                 // want "error-is: invalid usage of reqObj\\.NoError, use reqObj\\.NotErrorIs instead"
			reqObj.NoError(err, errSentinel, "msg with arg %d", 42) // want "error-is: invalid usage of reqObj\\.NoError, use reqObj\\.NotErrorIs instead"
		}

		// Valid.

		{
			reqObj.Error(err)
			reqObj.Error(err, "msg")
			reqObj.Error(err, "msg with arg %d", 42)
			reqObj.Errorf(err, "msg")
			reqObj.Errorf(err, "msg with arg %d", 42)

			reqObj.ErrorIs(err, errSentinel)
			reqObj.ErrorIs(err, errSentinel, "msg")
			reqObj.ErrorIs(err, errSentinel, "msg with arg %d", 42)
			reqObj.ErrorIsf(err, errSentinel, "msg")
			reqObj.ErrorIsf(err, errSentinel, "msg with arg %d", 42)

			reqObj.NoError(err)
			reqObj.NoError(err, "msg")
			reqObj.NoError(err, "msg with arg %d", 42)
			reqObj.NoErrorf(err, "msg")
			reqObj.NoErrorf(err, "msg with arg %d", 42)

			reqObj.NotErrorIs(err, errSentinel)
			reqObj.NotErrorIs(err, errSentinel, "msg")
			reqObj.NotErrorIs(err, errSentinel, "msg with arg %d", 42)
			reqObj.NotErrorIsf(err, errSentinel, "msg")
			reqObj.NotErrorIsf(err, errSentinel, "msg with arg %d", 42)
		}
	}
}
