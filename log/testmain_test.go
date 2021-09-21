package log_test

import (
	"github.com/annakallo/travel-log-server/testutil"
	"testing"
)

func TestMain(m *testing.M) {
	testutil.GlobalTearUp()
	code := m.Run()
	testutil.GlobalTearDown(code)
}
