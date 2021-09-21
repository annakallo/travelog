package settings_test

import (
	"github.com/annakallo/travelog/testutil"
	"testing"
)

func TestMain(m *testing.M) {
	testutil.GlobalTearUp()
	code := m.Run()
	testutil.GlobalTearDown(code)
}
