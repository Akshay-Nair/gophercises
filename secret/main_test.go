package main

import (
	"testing"

	"github.ibm.com/CloudBroker/dash_utils/dashtest"
)

func TestM(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			t.Errorf("error occured while calling the main function")
		}
	}()
	main()
}

func TestMain(m *testing.M) {
	dashtest.ControlCoverage(m)
}
