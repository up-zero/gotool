package idutil

import (
	"testing"
)

func TestNewObjectID(t *testing.T) {
	for i := 0; i < 3; i++ {
		id := NewObjectID()
		t.Log(id)
	}
}
