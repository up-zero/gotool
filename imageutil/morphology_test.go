package imageutil

import "testing"

func TestErodeFile(t *testing.T) {
	if err := ErodeFile("test.png", "test_erode.png", NewRectKernel(3, 3)); err != nil {
		t.Fatal(err)
	}
}

func TestDilateFile(t *testing.T) {
	if err := DilateFile("test.png", "test_dilate.png", NewRectKernel(3, 3)); err != nil {
		t.Fatal(err)
	}
}
