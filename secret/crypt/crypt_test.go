package crypt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidEncrypt(t *testing.T) {
	secret, err := Encrypt("hello123", "abc")
	fmt.Println("secret")
	assert.NotEqualf(t, len(secret), 0, "they must not be equal")
	assert.Equalf(t, len(err.Error()), 0, "they should be equal")
}

func TestInValidKeyEncrypt(t *testing.T) {
	secret, err := Encrypt("", "abc")
	assert.Equalf(t, len(secret), 0, "they must not be equal")
	assert.NotEqualf(t, len(err.Error()), 0, "they should be equal")
}

func TestInValidSecretEncrypt(t *testing.T) {
	secret, err := Encrypt("hello123", "")
	assert.Equalf(t, len(secret), 0, "they must not be equal")
	assert.NotEqualf(t, len(err.Error()), 0, "they should be equal")
}
