package filehandle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetContent(t *testing.T) {
	testCase := []struct {
		filePath   string
		content    bool
		errorOccur bool
		msg        string
	}{
		{
			"/usr/local/abc.go",
			false,
			true,
			"they should be equal",
		},
		{
			"/home/gslab/go/src/middleware/dir/dir.go",
			true,
			false,
			"they should be equal",
		},
		//test case for permission restricted files
		{
			"/home/gslab/Desktop/nonreadable.txt",
			false,
			true,
			"they should be equal",
		},
	}

	for _, test := range testCase {
		data, err := GetData(test.filePath)
		check1 := (len(data) != 0)
		check2 := (err != nil)
		check := (check1 == test.content) && (check2 == test.errorOccur)
		assert.Equalf(t, true, check, test.msg)
	}

}
