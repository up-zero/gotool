package cryptoutil

import (
	"testing"
)

func TestSha1(t *testing.T) {
	t.Log(Sha1("hello world"))         // 2aae6c35c94fcfb415dbe95f408b9ce91ee846ed
	t.Log(Sha1([]byte("hello world"))) // 2aae6c35c94fcfb415dbe95f408b9ce91ee846ed
	t.Log(Sha1("hello", " world"))     // 2aae6c35c94fcfb415dbe95f408b9ce91ee846ed
}

func TestSha256(t *testing.T) {
	t.Log(Sha256("hello world"))         // b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9
	t.Log(Sha256([]byte("hello world"))) // b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9
	t.Log(Sha256("hello", " world"))     // b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9
}

func TestSha512(t *testing.T) {
	t.Log(Sha512("hello world"))         // 309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f
	t.Log(Sha512([]byte("hello world"))) // 309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f
	t.Log(Sha512("hello", " world"))     // 309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f
}
