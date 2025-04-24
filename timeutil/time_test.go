package timeutil

import (
	"testing"
)

func TestRFC3339ToNormalTime(t *testing.T) {
	rfc3339Time := "2023-08-17T14:30:00+08:00"
	t.Log(RFC3339ToNormalTime(rfc3339Time))
}

func TestRFC1123ToNormalTime(t *testing.T) {
	rfc1123Time := "Mon, 02 Jan 2006 15:04:05 MST"
	t.Log(RFC1123ToNormalTime(rfc1123Time))
}
