package cryptoutil

import "testing"

func TestAesCbcEncrypt(t *testing.T) {
	s, err := AesCbcEncrypt("123456", []byte("1234567890123456"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log("encrypt:", s)
}

func TestAesCbcDecrypt(t *testing.T) {
	s, err := AesCbcDecrypt("N34m6eMn3SceMhb9EG7EJ7+MgMYGeQQqOPOG1k+Oi9M=", []byte("1234567890123456"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log("decrypt:", s)
}

func TestAesGcmEncrypt(t *testing.T) {
	s, err := AesGcmEncrypt("123456", []byte("1234567890123456"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log("encrypt:", s)
}

func TestAesGcmDecrypt(t *testing.T) {
	s, err := AesGcmDecrypt("+pNHBma82FQqxmRkh1iX2P3KcTanDOrR4sP3Dm72qXr0cQ==", []byte("1234567890123456"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log("decrypt:", s)
}
