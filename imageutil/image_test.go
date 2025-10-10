package imageutil

import (
	"image"
	"image/png"
	"os"
	"testing"
)

func TestCompression(t *testing.T) {
	if err := Compression("test.png", "compressed.png", 20); err != nil {
		t.Fatal("Error compressing image:", err)
	}
}

func TestSize(t *testing.T) {
	if reply, err := Size("test.png"); err != nil {
		t.Fatal("Error getting image size:", err)
	} else {
		t.Log("Image size:", reply.Width, "x", reply.Height)
	}
}

func TestGenerateCaptcha(t *testing.T) {
	img, err := GenerateCaptcha("5679")
	if err != nil {
		t.Fatal("Error generating captcha:", err)
		return
	}
	outFile, err := os.Create("captcha.png")
	if err != nil {
		t.Fatal("Error creating file:", err)
		return
	}
	defer outFile.Close()

	// 保存图片为 PNG 格式
	err = png.Encode(outFile, img)
	if err != nil {
		t.Fatal("Error encoding image:", err)
		return
	}
}

func TestCrop(t *testing.T) {
	if err := Crop("test.png", "cropped.png", image.Rect(100, 100, 200, 200)); err != nil {
		t.Fatal("Error cropping image:", err)
	}
}
