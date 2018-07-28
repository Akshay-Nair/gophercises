package cmd

import (
	"fmt"
	"os/exec"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	r, _ := regexp.Compile("Secret Key :  ")

	testCases := []struct {
		flag        string
		serviceName string
		expected    bool
		msg         string
	}{
		{
			"hello123",
			"twitter",
			true,
			"it should be true",
		},
		{
			"hello123",
			"google",
			false,
			"it should be false",
		}}

	for _, test := range testCases {
		cmdResult, err := exec.Command("../secret", "get", "--key", test.flag, test.serviceName).CombinedOutput()
		if err != nil {
			if err.Error() != "exit status 1" {
				t.Error("error occured while testing : ", err)
			}
		}
		fmt.Println(test.serviceName, test.expected, "obtained : ", r.MatchString(string(cmdResult)))
		assert.Equalf(t, r.MatchString(string(cmdResult)), test.expected, test.msg)
	}
}

func TestSet(t *testing.T) {
	r, _ := regexp.Compile("Secret saved successfully")

	testCases := []struct {
		flag        string
		serviceName string
		secret      string
		expected    bool
		msg         string
	}{
		{
			"hello123",
			"a",
			"",
			false,
			"it should be equal",
		},
	}

	for _, test := range testCases {
		cmdResult, err := exec.Command("../secret", "set", "--key", test.flag, test.serviceName, test.secret).CombinedOutput()
		if err != nil {
			if err.Error() != "exit status 1" {
				t.Error("error occured while testing : ", err)
			}
		}
		fmt.Println(string(cmdResult))
		assert.Equalf(t, r.MatchString(string(cmdResult)), test.expected, test.msg)
	}
}
