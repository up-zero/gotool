package imageutil

import (
	"image/png"
	"os"
	"testing"
)

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
