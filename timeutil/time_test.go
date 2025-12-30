package timeutil

import (
	"github.com/up-zero/gotool/testutil"
	"testing"
	"time"
)

func TestTransformLayout(t *testing.T) {
	res, _ := TransformLayout("2023-08-17T14:30:00+08:00", time.RFC3339, time.DateTime)
	testutil.Equal(t, res, "2023-08-17 14:30:00")
}

func TestFormatRFC3339(t *testing.T) {
	rfc3339Time := "2023-08-17T14:30:00+08:00"
	res, _ := FormatRFC3339(rfc3339Time)
	testutil.Equal(t, res, "2023-08-17 14:30:00")
}

func TestFormatRFC1123(t *testing.T) {
	rfc1123Time := "Mon, 02 Jan 2006 15:04:05 MST"
	res, _ := FormatRFC1123(rfc1123Time)
	testutil.Equal(t, res, "2006-01-02 15:04:05")
}
