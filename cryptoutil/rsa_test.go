package cryptoutil

import (
	"fmt"
	"testing"
)

func TestRsaGenerateKey(t *testing.T) {
	prvKey, pubKey, err := RsaGenerateKey(2048)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v \n", prvKey)
	fmt.Printf(pubKey)
}

func TestRsaEncrypt(t *testing.T) {
	key := `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDmWcb/iq27I5xz2dXtdHZXHFoT
bDex7f0c4HVOUzvUOpxMWgV+4yHo2BwtPHk1M9udd+S00iYZ0hErG9pOka3vp4+/
XcJtbo28MrUUMArl02PajhQ00rCYL6lM8X6mYOneX8HLHP/UOsPbdqPmaqpZOvgE
L8v87Pz0ZsvQTwIBwQIDAQAB
-----END PUBLIC KEY-----`
	res, err := RsaEncrypt([]byte("hello world"), key)
	if err != nil {
		t.Fatal(err)
	}
	base64 := Base64Encode(string(res))
	t.Log(base64)
}

func TestRsaDecrypt(t *testing.T) {
	key := `-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDmWcb/iq27I5xz2dXtdHZXHFoTbDex7f0c4HVOUzvUOpxMWgV+
4yHo2BwtPHk1M9udd+S00iYZ0hErG9pOka3vp4+/XcJtbo28MrUUMArl02PajhQ0
0rCYL6lM8X6mYOneX8HLHP/UOsPbdqPmaqpZOvgEL8v87Pz0ZsvQTwIBwQIDAQAB
AoGBAJpriOn6b6jKoLFRUQZUUpjnGsN2goe3QT1Ag6M0TSFjATB2EMUVQsAzUABL
T/4Ie6s+moVVc9FPb870Zw7QvkZq9fV7sErGxwfUNXuZMqgVc/Oc6wLdTUQO1tRt
rD+4ceInJPXnppaCY+ZjCinLzLV/LsXQ8nwq09kwOdUJik75AkEA9UAJvUaG1/OF
13wymyy5+45GdftjlKkQt0piA6YrccSlOuoqYu3gr297EBqU6UTEDqurRhoWvKhx
I2ckih6+hwJBAPByjUiIoru+i6TDD9re+sSj5FZiEx6iGQymLG6yv6QkwOq31H1Z
yOKvNriNjjXmd3fQRlChrqOQMSZbieUgR3cCQE314mKxpbHGLti2GVwslp55trpQ
hHJAYBjz4z5nt02+Bgw5Xen+1jrOhF81I/sXKf/D4HkzV+D25qgrZHknlscCQDzx
m/PNhgm2Eyjws/0S1Vav/7kRZK04Asdc+xgXwFE3a5pSe85FypACPwlp/6iTwKYi
qR/Yyy3z5zFEtF/Z3aMCQQCxCxXAB82CUWn3K1d9ybnDCv1zl1rTRJpRKb6S3rbY
phqBC6+MR83aZ/3MVO0rVAeBfYyNVghjd5Ktbt216CjW
-----END RSA PRIVATE KEY-----`
	base64 := Base64Decode("HOcBHh5A6bN82kxFwBKg/0kGOOVe4DyBs2rX41rBC0M/6GiFTdtpwdraOPjUaFrfWRmn7+1wnjtcln7Fxh1EdLolKKS2ySaY8tCEEv4uGwaUA0qbe9gKwT6xtmCu5XhAdR3/OBT6YHFeCs38hq45wZS5zpx4mQF5vL9zakWpBbU=")
	res, err := RsaDecrypt([]byte(base64), key)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(res))
}
