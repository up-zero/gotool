package gotool

import "testing"

func TestStrToInt64(t *testing.T) {
	t.Log(StrToInt64("123"))
}

func TestStrToUint64(t *testing.T) {
	t.Log(StrToUint64("123"))
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
