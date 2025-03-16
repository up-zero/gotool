package gotool

import "testing"

func TestSysUptime(t *testing.T) {
	uptime, err := SysUptime()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(uptime)
}
