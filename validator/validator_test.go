package validator

import (
	"github.com/up-zero/gotool/testutil"
	"testing"
)

func TestIsDigit(t *testing.T) {
	t.Log(IsDigit("123"))   // true
	t.Log(IsDigit("Hello")) // false
}

func TestIsAlpha(t *testing.T) {
	t.Log(IsAlpha("Hello")) // true
	t.Log(IsAlpha("你好"))    // true
	t.Log(IsAlpha("你好，"))   // false
}

func TestIsAlphaStrict(t *testing.T) {
	t.Log(IsAlphaStrict("Hello")) // true
	t.Log(IsAlphaStrict("你好"))    // false
}

func TestIsAlphaNumeric(t *testing.T) {
	t.Log(IsAlphaNumeric("Hello123"))  // true
	t.Log(IsAlphaNumeric("Hello 123")) // false
	t.Log(IsAlphaNumeric("你好123"))     // true
	t.Log(IsAlphaNumeric("你好,123"))    // false
}

func TestIsIp(t *testing.T) {
	testutil.Equal(t, IsIp("192.168.1.1"), true)
	testutil.Equal(t, IsIp("256.0.0.0"), false)
}

func TestIsIpv4(t *testing.T) {
	t.Log(IsIpv4("192.168.1.1"))                             // true
	t.Log(IsIpv4("256.0.0.1"))                               // false
	t.Log(IsIpv4("2001:0db8:85a3:0000:0000:8a2e:0370:7334")) // false
	t.Log(IsIpv4("192.168.1.1:8080"))                        // false
	t.Log(IsIpv4("::1"))                                     // false
}

func TestIsIpv6(t *testing.T) {
	t.Log(IsIpv6("192.168.1.1"))                             // false
	t.Log(IsIpv6("256.0.0.1"))                               // false
	t.Log(IsIpv6("2001:0db8:85a3:0000:0000:8a2e:0370:7334")) // true
	t.Log(IsIpv6("192.168.1.1:8080"))                        // false
	t.Log(IsIpv6("::1"))                                     // true
}

func TestContainChinese(t *testing.T) {
	testutil.Equal(t, ContainChinese("你好"), true)
	testutil.Equal(t, ContainChinese("Hello"), false)
	testutil.Equal(t, ContainChinese("，"), true)
	testutil.Equal(t, ContainChinese(","), false)
	testutil.Equal(t, ContainChinese("こんにちは"), false)
}

func TestIsEmail(t *testing.T) {
	testutil.Equal(t, IsEmail("test@example.com"), true)
	testutil.Equal(t, IsEmail("name <user@example.com>"), false)
}
