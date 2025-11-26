package convertutil

import (
	"github.com/up-zero/gotool/testutil"
	"testing"
)

func TestStrToInt(t *testing.T) {
	testutil.Equal(t, StrToInt("123"), 123)
}

func TestStrToInt8(t *testing.T) {
	testutil.Equal(t, StrToInt8("123"), int8(123))
	testutil.Equal(t, StrToInt8("128"), int8(127))
}

func TestStrToInt16(t *testing.T) {
	testutil.Equal(t, StrToInt16("32760"), int16(32760))
	testutil.Equal(t, StrToInt16("32768"), int16(32767))
}

func TestStrToInt32(t *testing.T) {
	testutil.Equal(t, StrToInt32("2147483646"), int32(2147483646))
	testutil.Equal(t, StrToInt32("2147483690"), int32(2147483647))
}

func TestStrToInt64(t *testing.T) {
	testutil.Equal(t, StrToInt64("123"), int64(123))
}

func TestStrToUint(t *testing.T) {
	testutil.Equal(t, StrToUint("123"), uint(123))
}

func TestStrToUint8(t *testing.T) {
	testutil.Equal(t, StrToUint8("255"), uint8(255))
	testutil.Equal(t, StrToUint8("256"), uint8(255))
}

func TestStrToUint16(t *testing.T) {
	testutil.Equal(t, StrToUint16("65535"), uint16(65535))
	testutil.Equal(t, StrToUint16("65536"), uint16(65535))
}

func TestStrToUint32(t *testing.T) {
	testutil.Equal(t, StrToUint32("4294967295"), uint32(4294967295))
	testutil.Equal(t, StrToUint32("4294967296"), uint32(4294967295))
}

func TestStrToUint64(t *testing.T) {
	testutil.Equal(t, StrToUint64("123"), uint64(123))
}

func TestStrToFloat32(t *testing.T) {
	testutil.Equal(t, StrToFloat32("123.1001"), float32(123.1001))
}

func TestStrToFloat64(t *testing.T) {
	testutil.Equal(t, StrToFloat64("123.1001"), 123.1001)
}

func TestInt64ToStr(t *testing.T) {
	testutil.Equal(t, Int64ToStr(1231111111), "1231111111")
}

func TestUint64ToStr(t *testing.T) {
	testutil.Equal(t, Uint64ToStr(1231111111), "1231111111")
}

func TestFloat64ToStr(t *testing.T) {
	testutil.Equal(t, Float64ToStr(123.123), "123.123")
	testutil.Equal(t, Float64ToStr(123.123, 2), "123.12")
}

func TestInt64ToHex(t *testing.T) {
	testutil.Equal(t, Int64ToHex(15), "F")
	testutil.Equal(t, Int64ToHex(447, "08"), "000001BF")
}

func TestHexToInt64(t *testing.T) {
	testutil.Equal(t, HexToInt64("000001BF"), int64(447))
}

func TestDigitToChinese(t *testing.T) {
	testutil.Equal(t, DigitToChinese("1"), "一")
	testutil.Equal(t, DigitToChinese("138000"), "一三八零零零")
}
