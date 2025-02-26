package gotool

import "testing"

func TestStrToInt8(t *testing.T) {
	t.Log(StrToInt8("123")) // 123
	t.Log(StrToInt8("128")) // 127
}

func TestStrToInt16(t *testing.T) {
	t.Log(StrToInt16("32760")) // 32760
	t.Log(StrToInt16("32768")) // 32767
}

func TestStrToInt32(t *testing.T) {
	t.Log(StrToInt32("2147483646")) // 2147483646
	t.Log(StrToInt32("2147483648")) // 2147483647
}

func TestStrToInt64(t *testing.T) {
	t.Log(StrToInt64("123"))
}

func TestStrToUint8(t *testing.T) {
	t.Log(StrToUint8("255")) // 255
	t.Log(StrToUint8("256")) // 255
}

func TestStrToUint16(t *testing.T) {
	t.Log(StrToUint16("65535")) // 65535
	t.Log(StrToUint16("65536")) // 65535
}

func TestStrToUint32(t *testing.T) {
	t.Log(StrToUint32("4294967295")) // 4294967295
	t.Log(StrToUint32("4294967296")) // 4294967295
}

func TestStrToUint64(t *testing.T) {
	t.Log(StrToUint64("123"))
}

func TestStrToFloat32(t *testing.T) {
	t.Log(StrToFloat32("123.1001"))
}

func TestStrToFloat64(t *testing.T) {
	t.Log(StrToFloat64("123.1001"))
}

func TestInt64ToStr(t *testing.T) {
	t.Log(Int64ToStr(1231111111))
}

func TestUint64ToStr(t *testing.T) {
	t.Log(Uint64ToStr(1231111111))
}

func TestFloat64ToStr(t *testing.T) {
	t.Log(Float64ToStr(123.123))
	t.Log(Float64ToStr(123.123, 2))
}

func TestInt64ToHex(t *testing.T) {
	t.Log(Int64ToHex(15))
	t.Log(Int64ToHex(447, "08"))
}

func TestHexToInt64(t *testing.T) {
	t.Log(HexToInt64("000001BF"))
}
