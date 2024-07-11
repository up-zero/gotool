package gotool

import "testing"

func TestRandomStr(t *testing.T) {
	t.Log(RandomStr("12345677890", 6))
}

func TestRandomAlpha(t *testing.T) {
	t.Log(RandomAlpha(6))
}

func TestRandomNumber(t *testing.T) {
	t.Log(RandomNumber(6))
}

func TestRandomAlphaNumber(t *testing.T) {
	t.Log(RandomAlphaNumber(6))
}

func TestRandomRangeInt(t *testing.T) {
	t.Log(RandomRangeInt(5, 100))
}
