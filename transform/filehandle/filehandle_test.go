package filehandle

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.ibm.com/CloudBroker/dash_utils/dashtest"
)

func TestInvalidTempFile(t *testing.T) {
	_, err := TempFile("./test", "txt")
	assert.NotEqualf(t, err, nil, "they should not be equal")
}

func TestValidTempFile(t *testing.T) {
	_, err := TempFile("../input", "txt")
	assert.Equalf(t, err, nil, "they should be equal")
}

func TestValidFile(t *testing.T) {
	testCase := []struct {
		filename string
		status   bool
	}{
		{
			"abc.jpeg",
			true,
		},
		{
			"abc",
			false,
		},
		{
			"",
			false,
		},
		{
			"abc.pdf",
			false,
		},
	}

	for _, test := range testCase {
		assert.Equalf(t, test.status, ValFile(test.filename), "they should be equal")
	}
}

func TestMain(m *testing.M) {
	dashtest.ControlCoverage(m)
	m.Run()
}
