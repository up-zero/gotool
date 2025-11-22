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

func TestContainsAll(t *testing.T) {
	testutil.Equal(t, ContainsAll("hello world", "hello", "world"), true)
	testutil.Equal(t, ContainsAll("hello world", "hello", "world", "123"), false)
}

func TestCamelToSnake(t *testing.T) {
	testutil.Equal(t, CamelToSnake("helloWorld"), "hello_world")
	testutil.Equal(t, CamelToSnake("HTTPRequest"), "http_request")
	testutil.Equal(t, CamelToSnake("v1G"), "v1_g")
}

func TestSnakeToCamel(t *testing.T) {
	testutil.Equal(t, SnakeToCamel("hello_world"), "helloWorld")
	testutil.Equal(t, SnakeToCamel("___hello__world"), "helloWorld")
}

func TestSnakeToPascal(t *testing.T) {
	testutil.Equal(t, SnakeToPascal("hello_world"), "HelloWorld")
	testutil.Equal(t, SnakeToPascal("__hello_world"), "HelloWorld")
}
