package stringutil

import "testing"

func TestReverse(t *testing.T) {
	t.Log(Reverse("hello world")) // dlrow olleh
	t.Log(Reverse("你好 世界!"))      // !界世 好你
}
