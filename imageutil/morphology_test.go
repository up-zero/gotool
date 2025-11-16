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

func TestMorphologyOpenFile(t *testing.T) {
	if err := MorphologyOpenFile("test.png", "test_open.png", NewRectKernel(3, 3)); err != nil {
		t.Fatal(err)
	}
}

func TestMorphologyCloseFile(t *testing.T) {
	if err := MorphologyCloseFile("test.png", "test_close.png", NewRectKernel(3, 3)); err != nil {
		t.Fatal(err)
	}
}
