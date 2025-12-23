package mathutil

import (
	"github.com/up-zero/gotool/testutil"
	"testing"
)

func TestHammingDistance(t *testing.T) {
	testutil.Equal(t, HammingDistance(0x0000000000000000, 0x0000000000000000), 0)
	testutil.Equal(t, HammingDistance(0xFFFFFFFFFFFFFFFF, 0x0000000000000000), 64)
}
