package cmd

import (
	"gophercises/task/db"
	"os"
	"testing"

	"github.ibm.com/CloudBroker/dash_utils/dashtest"

	"github.com/atrox/homedir"
)

func TestMain(m *testing.M) {
	closeDB()
	dir, _ := homedir.Dir()
	dir = dir + "/task_db.db"
	os.Remove(dir)
	db.Connection()
	dashtest.ControlCoverage(m)
	m.Run()
}
