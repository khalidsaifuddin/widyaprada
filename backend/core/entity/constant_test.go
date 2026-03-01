package entity

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrapRecordNotFound(t *testing.T) {
	err := WrapRecordNotFound("user not found")
	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())
	var rnf *RecordNotFoundError
	assert.True(t, errors.As(err, &rnf))
	assert.Equal(t, "user not found", rnf.Message)
	assert.ErrorIs(t, err, ErrRecordNotFound)
}

func TestWrapRecordNotFoundf(t *testing.T) {
	err := WrapRecordNotFoundf("user %s not found", "123")
	assert.Error(t, err)
	assert.Equal(t, "user 123 not found", err.Error())
	var rnf *RecordNotFoundError
	assert.True(t, errors.As(err, &rnf))
	assert.Equal(t, "user 123 not found", rnf.Message)
}

func TestRecordNotFoundError_Error(t *testing.T) {
	t.Run("with message", func(t *testing.T) {
		e := &RecordNotFoundError{Message: "custom", Err: ErrRecordNotFound}
		assert.Equal(t, "custom", e.Error())
	})
	t.Run("empty message uses Err", func(t *testing.T) {
		e := &RecordNotFoundError{Err: ErrRecordNotFound}
		assert.Equal(t, ErrRecordNotFound.Error(), e.Error())
	})
}

func TestRecordNotFoundError_Unwrap(t *testing.T) {
	e := &RecordNotFoundError{Err: ErrRecordNotFound}
	assert.Equal(t, ErrRecordNotFound, e.Unwrap())
}

func TestIsRecordNotFound(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{"wrapped RecordNotFoundError", WrapRecordNotFound("x"), true},
		{"plain ErrRecordNotFound", ErrRecordNotFound, true},
		{"other error", errors.New("other"), false},
		{"nil", nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsRecordNotFound(tt.err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestEmailRegex(t *testing.T) {
	valid := []string{
		"a@b.com", "user@example.org", "test+tag@domain.co.uk",
		"user.name@domain.com", "user_name@domain.com",
	}
	invalid := []string{
		"", "no-at-sign", "@domain.com", "user@", "user@.com",
		"user @domain.com", "user@domain",
	}
	for _, e := range valid {
		assert.True(t, EmailRegex.MatchString(e), "expected valid: %s", e)
	}
	for _, e := range invalid {
		assert.False(t, EmailRegex.MatchString(e), "expected invalid: %s", e)
	}
}

func TestConstants(t *testing.T) {
	assert.Equal(t, "PESERTA", DefaultRoleCodePeserta)
	assert.GreaterOrEqual(t, BcryptCost, 4)
	assert.LessOrEqual(t, BcryptCost, 31)
}
