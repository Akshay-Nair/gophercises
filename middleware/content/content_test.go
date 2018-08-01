package content

import (
	"regexp"
	"testing"

	"github.ibm.com/CloudBroker/dash_utils/dashtest"

	"github.com/stretchr/testify/assert"
)

func TestExtractPath(t *testing.T) {
	testCase := []struct {
		location   string
		dataZero   bool
		errorOccur bool
		msg        string
	}{
		{
			"/usr/local/go/src/runtime/panic.go:80",
			false,
			false,
			"they should be equal",
		},
		{
			"asdfad",
			true,
			true,
			"they should be  equal",
		},
		{
			"/usr/local/go/src/runtime/panic.go",
			true,
			true,
			"they should be equal",
		},
	}
	for _, test := range testCase {
		_, data, err := extractPath(test.location)
		check1 := (data == 0)
		check2 := (err != nil)
		check := (check1 == test.dataZero) && (check2 == test.errorOccur)
		assert.Equalf(t, true, check, test.msg)
	}
}

func TestFindAndReplacePath(t *testing.T) {
	re, _ := regexp.Compile("<a href=.+</a>")
	testCase := []struct {
		content  string
		urlCount int
		msg      string
	}{
		{
			"abcde",
			0,
			"they should be equal",
		},
		{
			`goroutine 9 [running]:
			runtime/debug.Stack(0xc420057b98, 0x8268e0, 0x9ecb20)
					/usr/local/go/src/runtime/debug/stack.go:24 +0xa7
			main.middleware.func1.1(0x9f57c0, 0xc420064000)
					/home/gslab/go/src/middleware/main.go:52 +0x5f`,
			2,
			"they should be equal",
		},
		{
			"goroutine 4[running]",
			0,
			"they should be equal",
		},
	}
	for _, test := range testCase {
		testOutput, _ := FindAndReplacePath(test.content)
		testResult := re.FindAllString(testOutput, -1)
		assert.Equalf(t, len(testResult), test.urlCount, test.msg)
	}
}

func TestMain(m *testing.M) {
	dashtest.ControlCoverage(m)
}
