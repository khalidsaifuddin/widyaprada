package helper

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObfuscateErrorWithContext(t *testing.T) {
	err := errors.New("db connection failed")
	got := ObfuscateErrorWithContext(err, "create user")
	assert.Error(t, got)
	assert.Contains(t, got.Error(), "create user")
	assert.Contains(t, got.Error(), "db connection failed")
	assert.ErrorIs(t, got, err)
}

func TestObfuscateErrorWithContext_NilError(t *testing.T) {
	got := ObfuscateErrorWithContext(nil, "create user")
	assert.NoError(t, got)
	assert.Nil(t, got)
}

func TestObfuscateError(t *testing.T) {
	err := errors.New("something went wrong")
	got := ObfuscateError(err)
	assert.Equal(t, err, got)
}

func TestObfuscateError_NilError(t *testing.T) {
	got := ObfuscateError(nil)
	assert.Nil(t, got)
}
