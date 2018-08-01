package db

import (
	"os"
	"testing"

	"github.ibm.com/CloudBroker/dash_utils/dashtest"
)

func TestMain(m *testing.M) {
	dir, _ := getHomeDir()
	dir = dir + "/task_db.db"
	os.Remove(dir)
	Connection()
	dashtest.ControlCoverage(m)
	m.Run()
}
