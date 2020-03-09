package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckPasswordHash(t *testing.T) {
	password := "Test"
	hash, err := HashPassword(password)
	assert.NotEmpty(t, hash)
	assert.NoError(t, err)

	check := CheckPasswordHash("Test", hash)

	assert.Equal(t, true, check)
}

func TestHashPassword(t *testing.T) {
	password := "Test"
	hash, err := HashPassword(password)
	assert.NotEmpty(t, hash)
	assert.NoError(t, err)
}
