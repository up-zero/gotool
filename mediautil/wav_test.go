package mediautil

import "testing"

func TestReadWavHeader(t *testing.T) {
	header, err := ReadWavHeader("test.wav")
	if err != nil {
		t.Fatalf("read wav header error: %v", err)
		return
	}
	t.Logf("wav header: %+v", header)
}
