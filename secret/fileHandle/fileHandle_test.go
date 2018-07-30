package fileHandle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	tempName := fileName
	fileName = "../temp.csv"
	defer func() {
		fileName = tempName
	}()
	testCase := []struct {
		key      string
		expected string
	}{
		{
			"abc",
			"",
			//"c647e6a250b9a8cd61087c03f15eb914a98808a15ecc",
		},
		{
			"def",
			"",
		},
		{
			"google",
			"",
		},
		{
			"",
			"",
		},
	}
	for _, test := range testCase {
		key, err := GetSecret(test.key)
		if err != nil {
			fmt.Println(err)
		}
		assert.Equalf(t, key, test.expected, "should be equal")
	}
}

func TestInvalid(t *testing.T) {
	tempName := fileName
	fileName = "../abc.csv" //testing for non existing file
	defer func() {
		fileName = tempName
	}()
	_, err := GetSecret("abc")
	assert.Equalf(t, err.Error(), "invalid argument", "they should be equal")
	err = SetSecret("abc", "def")
	assert.NotEqualf(t, err, "nil", "they should not be equal")
}

func TestWrite(t *testing.T) {
	tempName := fileName
	fileName = "../temp.csv"
	defer func() {
		fileName = tempName
	}()
	testCase := []struct {
		serviceName string
		key         string
		expected    error
	}{
		{
			"twitter",
			"abc",
			nil,
		},
		{
			"google",
			"",
			nil,
		},
		{
			"",
			"",
			nil,
		},
		{
			"",
			"abc123",
			nil,
		},
	}
	for _, test := range testCase {
		err := SetSecret(test.serviceName, test.key)
		if err != nil {
			fmt.Println(err)
		}
		assert.Equalf(t, err, test.expected, "should be equal")
	}
}
