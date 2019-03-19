package model

import (
	"os"
	"testing"

	"github.com/sdk_go/src/config"
	"github.com/sdk_go/src/util"
)

func TestInsert(t *testing.T) {
	// Load the configuration file
	util.Load("../config"+string(os.PathSeparator)+"config.json", config.Config)
	// createPrefences()
	id := InsertOrder()
	if id == 0 {
		t.Error("Error during to insert the order")
	}
}
