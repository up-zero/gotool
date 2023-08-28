package gotool

import (
	"testing"
)

func TestIpv4sLocal(t *testing.T) {
	res, err := Ipv4sLocal()
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range res {
		t.Log(v)
	}
}
