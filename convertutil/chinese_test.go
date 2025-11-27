package convertutil

import (
	"github.com/up-zero/gotool/testutil"
	"testing"
)

func TestDigitToChinese(t *testing.T) {
	testutil.Equal(t, DigitToChinese("1"), "一")
	testutil.Equal(t, DigitToChinese("138000"), "一三八零零零")
}

func TestNumberToChinese(t *testing.T) {
	testutil.Equal(t, IntegerToChinese(1), "一")
	testutil.Equal(t, IntegerToChinese(138000), "十三万八千")
	testutil.Equal(t, IntegerToChinese(-10001), "负一万零一")
}
