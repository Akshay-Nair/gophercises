package crypt

import (
	"testing"

	"github.ibm.com/CloudBroker/dash_utils/dashtest"

	"github.com/stretchr/testify/assert"
)

var hexCode string

func TestEncrypt(t *testing.T) {
	secret, err := Encrypt("hello123", "abc")
	hexCode = secret
	assert.NotEqualf(t, len(secret), 0, "they must not be equal")
	assert.Equalf(t, err, nil, "they should be equal")
}

func TestDecrypt(t *testing.T) {
	secret, err := Decrypt("hello123", hexCode)
	assert.NotEqualf(t, len(secret), 0, "they must not be equal")
	assert.Equalf(t, err, nil, "they should be equal")
}

func TestInvalidKeyDecrypt(t *testing.T) {
	secret, err := Decrypt("hello123", "123333")
	assert.Equalf(t, len(secret), 0, "they must not be equal")
	assert.NotEqualf(t, err, nil, "they should be equal")
}

func TestInvalidHexcodeDecrypt(t *testing.T) {
	secret, err := Decrypt("hello123", "1+3333")
	assert.Equalf(t, len(secret), 0, "they must not be equal")
	assert.NotEqualf(t, err, nil, "they should be equal")
}

func TestMain(m *testing.M) {
	dashtest.ControlCoverage(m)
}
