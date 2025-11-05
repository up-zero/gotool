package mathutil

import (
	"github.com/up-zero/gotool/testutil"
	"testing"
)

func TestAbs(t *testing.T) {
	testutil.Equal(t, Abs(-42), 42)
	testutil.Equal(t, Abs(-123456789), 123456789)
	testutil.Equal(t, Abs(42), 42)

	testutil.Equal(t, Abs(-3.14), 3.14)
	testutil.Equal(t, Abs(-0.0), 0.00)
}

func TestMin(t *testing.T) {
	testutil.Equal(t, Min(12, 20, -1), -1)
	testutil.Equal(t, Min(12, 20, -1.1), -1.1)
}

func TestMax(t *testing.T) {
	testutil.Equal(t, Max(12, 20, -1), 20)
	testutil.Equal(t, Max(-100, -1.12, -1.1), -1.1)
}

func TestSum(t *testing.T) {
	testutil.Equal(t, Sum(12, 200, -1), 211)
	testutil.Equal(t, Sum(-100, -1.12, -1.1), -102.22)
}

func TestAverage(t *testing.T) {
	testutil.Equal(t, Average(12, 200, -1), 70)
	testutil.Equal(t, Average(-100, -1.12, -1.1), -34.07333333333333)
}
