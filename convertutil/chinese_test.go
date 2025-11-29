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

func TestTextToChinese(t *testing.T) {
	testutil.Equal(t, TextToChinese("13万8千"), "十三万八千")
	testutil.Equal(t, TextToChinese("2025年5月20日，也就是12:30，增长率达到了50.5%，数值是3.14159。"),
		"二零二五年五月二十日,也就是十二点三十分,增长率达到了百分之五十点五,数值是三点一四一五九.")
}
