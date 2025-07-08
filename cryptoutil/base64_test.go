package cryptoutil

import "testing"

func TestBase64Encode(t *testing.T) {
	t.Log(Base64Encode("Hello World")) // SGVsbG8gV29ybGQ=
}

func TestBase64Decode(t *testing.T) {
	t.Log(Base64Decode("SGVsbG8gV29ybGQ=")) // Hello World
}
