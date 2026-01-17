package randomutil

import "testing"

func TestString(t *testing.T) {
	for i := 0; i < 5; i++ {
		t.Log(String("12345677890", 6))
	}
}

func TestDigits(t *testing.T) {
	for i := 0; i < 5; i++ {
		t.Log(Digits(6))
	}
}

func TestLetters(t *testing.T) {
	for i := 0; i < 5; i++ {
		t.Log(Letters(6))
	}
}

func TestAlphanumeric(t *testing.T) {
	for i := 0; i < 5; i++ {
		t.Log(Alphanumeric(6))
	}
}

func TestIntRange(t *testing.T) {
	for i := 0; i < 5; i++ {
		t.Log(IntRange(5, 100))
	}
}

func TestBool(t *testing.T) {
	for i := 0; i < 5; i++ {
		t.Log(Bool())
	}
}

func TestChoice(t *testing.T) {
	for i := 0; i < 5; i++ {
		t.Logf("%v", Choice([]string{"1", "2", "3", "4", "5"}))
	}
}
