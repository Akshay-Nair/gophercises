package primitive

import (
	"testing"

	"github.ibm.com/CloudBroker/dash_utils/dashtest"

	"github.com/atrox/homedir"
	"github.com/stretchr/testify/assert"
)

func TestTransform(t *testing.T) {
	hDir, _ := homedir.Dir()
	_, err := Transform(hDir+"/go/src/gophercises/transform/sample.jpg", 3, 120, "jpg")
	if err != nil {
		t.Error(err.Error())
	}
	assert.Equalf(t, nil, err, "they should be equal")
}

func TestBatchTransformMode(t *testing.T) {
	testOutput := BatchTransform("/home/gslab/go/src/gophercises/transform/sample.jpg", []int{0, 1, 2, 3}, []int{50}, "jpg")
	assert.Equalf(t, len(testOutput), 4, "they should be equal")
}

func TestBatchTransformNumber(t *testing.T) {
	testOutput := BatchTransform("/home/gslab/go/src/gophercises/transform/sample.jpg", []int{0}, []int{50, 100, 150, 200}, "jpg")
	assert.Equalf(t, len(testOutput), 4, "they should be equal")
}

func TestMain(m *testing.M) {
	dashtest.ControlCoverage(m)
	m.Run()
}
