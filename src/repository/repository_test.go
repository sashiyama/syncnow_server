package repository_test

import (
	"github.com/sashiyama/syncnow_server/util"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	util.TruncateAllTables()
	code := m.Run()
	util.TruncateAllTables()
	os.Exit(code)
}
