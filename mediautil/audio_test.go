package mediautil

import (
	"github.com/up-zero/gotool/testutil"
	"testing"
)

func TestPreEmphasis(t *testing.T) {
	testutil.EqualFloat(t, PreEmphasis([]float32{10.0, 20.0, 30.0}, float32(0.97)),
		[]float32{10.0, 10.3, 10.6})
}
