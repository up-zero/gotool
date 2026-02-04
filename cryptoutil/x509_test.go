package cryptoutil

import (
	"fmt"
	"testing"
)

func TestGenSelfSignedCert(t *testing.T) {
	certPEM, keyPEM, err := GenSelfSignedCert(CertConfig{
		CommonName:   "192.168.1.8",
		Organization: []string{"UpZero"},
		Years:        1,
		Hosts:        []string{"192.168.1.8", "getcharzp.cn"},
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v \n%v\n", certPEM, keyPEM)
}
