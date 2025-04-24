package mathutil

import "testing"

func TestRandomStr(t *testing.T) {
	for i := 0; i < 5; i++ {
		t.Log(RandomStr("12345677890", 6))
	}
}

func TestRandomAlpha(t *testing.T) {
	for i := 0; i < 5; i++ {
		t.Log(RandomAlpha(6))
	}
}

func TestRandomNumber(t *testing.T) {
	for i := 0; i < 5; i++ {
		t.Log(RandomNumber(6))
	}
}

func TestRandomAlphaNumber(t *testing.T) {
	for i := 0; i < 5; i++ {
		t.Log(RandomAlphaNumber(6))
	}
}

func TestRandomRangeInt(t *testing.T) {
	for i := 0; i < 5; i++ {
		t.Log(RandomRangeInt(5, 100))
	}
}
