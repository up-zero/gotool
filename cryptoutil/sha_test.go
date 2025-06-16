package cryptoutil

import (
	"testing"
)

func TestSha1(t *testing.T) {
	t.Log(Sha1("hello world"))         // 2aae6c35c94fcfb415dbe95f408b9ce91ee846ed
	t.Log(Sha1([]byte("hello world"))) // 2aae6c35c94fcfb415dbe95f408b9ce91ee846ed
	t.Log(Sha1("hello", " world"))     // 2aae6c35c94fcfb415dbe95f408b9ce91ee846ed
}
