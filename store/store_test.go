package store_test

import (
	"os"
	"testing"

	_ "github.com/kk-no/testable-api/database/mysql"
	"github.com/kk-no/testable-api/test/testutils"
)

func TestMain(m *testing.M) {
	testutils.TruncateTables()
	os.Exit(m.Run())
}
