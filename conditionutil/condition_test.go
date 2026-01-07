package conditionutil

import (
	"github.com/up-zero/gotool/testutil"
	"testing"
)

func TestIsZero(t *testing.T) {
	type User struct {
		Name string
	}

	testutil.Equal(t, IsZero(User{}), true)
	testutil.Equal(t, IsZero(User{Name: "test"}), false)
}
