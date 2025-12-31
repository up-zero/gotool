package conditionutil

import (
	"github.com/up-zero/gotool/testutil"
	"testing"
)

func TestIf(t *testing.T) {
	testutil.Equal(t, If(true, 1, 2), 1)
	testutil.Equal(t, If(true, 1.2, 2.4), 1.2)
	testutil.Equal(t, If(false, "1.3", "2.6"), "2.6")
}

func TestIfLazy(t *testing.T) {
	testutil.Equal(t, IfLazy(true, func() int { return 1 }, func() int { return 2 }), 1)
	testutil.Equal(t, IfLazy(false, func() string { return "1.3" }, func() string { return "2.6" }), "2.6")
}
