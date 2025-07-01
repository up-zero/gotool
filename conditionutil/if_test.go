package conditionutil

import "testing"

func TestIf(t *testing.T) {
	t.Log(If(true, 1, 2))
	t.Log(If(true, 1.2, 2.4))
	t.Log(If(true, "1.3", "2.6"))
}
