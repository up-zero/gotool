package stringutil

import "testing"

func TestReverse(t *testing.T) {
	t.Log(Reverse("hello world")) // dlrow olleh
	t.Log(Reverse("你好 世界!"))      // !界世 好你
}

func TestTruncate(t *testing.T) {
	t.Log(Truncate("hello world", 5)) // hello
	t.Log(Truncate("hello", 5))       // hello
	t.Log(Truncate("hel", 5))         // hel
	t.Log(Truncate("你好 世界!", 6))      // 你好
}
