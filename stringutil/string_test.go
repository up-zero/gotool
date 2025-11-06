package stringutil

import (
	"github.com/up-zero/gotool/testutil"
	"testing"
)

func TestReverse(t *testing.T) {
	testutil.Equal(t, Reverse("hello world"), "dlrow olleh")
	testutil.Equal(t, Reverse("你好 世界!"), "!界世 好你")
}

func TestTakeFirst(t *testing.T) {
	testutil.Equal(t, TakeFirst("hello world", 5), "hello")
	testutil.Equal(t, TakeFirst("hello", 5), "hello")
	testutil.Equal(t, TakeFirst("hel", 5), "hel")
	testutil.Equal(t, TakeFirst("你好 世界!", 2), "你好")
}

func TestContainsAny(t *testing.T) {
	testutil.Equal(t, ContainsAny("hello world", "hello", "123"), true)
	testutil.Equal(t, ContainsAny("hello world", "123"), false)
}
