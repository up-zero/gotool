package gotool

import "testing"

func TestIf(t *testing.T) {
	t.Log(If(true, 1, 2))
}
