package gotool

import (
	"testing"
)

func TestPsByName(t *testing.T) {
	ps, err := PsByName("chrome")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ps)
}
