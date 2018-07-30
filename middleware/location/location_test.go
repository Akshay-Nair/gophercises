package location

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractLocation(t *testing.T) {
	testCase := []struct {
		content  string
		expected int
		msg      string
	}{
		{
			"/usr/local/go/src/runtime/panic.go:80",
			1,
			"they should be equal",
		},
		{
			"asdfad",
			0,
			"they should be  equal",
		},
		{
			"/usr/local/go/src/runtime/panic.go:43  /usr/local/go/src/net/http/server.go:80 /usr/local/go/src/net/http/se",
			2,
			"they should be equal",
		},
		{
			"/usr/local/go/src/runtime/panic.go  /usr/local/go/src/net/http/server.go:80 /usr/local/go/src/net/http/se",
			1,
			"they should be equal",
		},
	}
	for _, test := range testCase {
		data, err := ExtractLocation(test.content)
		if err != nil {
			t.Error(err)
		}
		assert.Equalf(t, len(data), test.expected, test.msg)
	}
}

func TestPath(t *testing.T) {
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
		_, data, err := ExtractPath(test.location)
		check1 := (data == 0)
		check2 := (err != nil)
		check := (check1 == test.dataZero) && (check2 == test.errorOccur)
		assert.Equalf(t, true, check, test.msg)
	}
}
