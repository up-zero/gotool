package imageutil

import (
	"image"
	"image/color"
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
	if err := CropFile("test.png", "cropped.png", image.Rect(100, 100, 200, 600)); err != nil {
		t.Fatal("Error cropping image:", err)
	}
}

func TestResizeFile(t *testing.T) {
	if err := ResizeFile("test.png", "resized.png", 200, 0); err != nil {
		t.Fatal("Error resizing image:", err)
	}
}

func TestRotateFile(t *testing.T) {
	if err := RotateFile("test.png", "rotated.png", RotateAngle270); err != nil {
		t.Fatal("Error rotating image:", err)
	}
}

func TestFlipFile(t *testing.T) {
	if err := FlipFile("test.png", "flipped.png", FlipModeHorizontal); err != nil {
		t.Fatal("Error flipping image:", err)
	}
}

func TestOverlayFile(t *testing.T) {
	if err := OverlayFile("test.png", "overlay.png", "overlayed.png", 200, 300); err != nil {
		t.Fatal("Error overlaying image:", err)
	}
}

func TestGrayscaleFile(t *testing.T) {
	if err := GrayscaleFile("test.png", "test_gray.png"); err != nil {
		t.Error(err)
	}
}

func TestGaussianBlurFile(t *testing.T) {
	if err := GaussianBlurFile("test.png", "test_gaussian.png", 3, 1); err != nil {
		t.Error(err)
	}
}

func TestAdjustBrightnessFile(t *testing.T) {
	if err := AdjustBrightnessFile("test.png", "test_brightness.png", 30); err != nil {
		t.Error(err)
	}
}

func TestInvertFile(t *testing.T) {
	if err := InvertFile("test.png", "test_invert.png"); err != nil {
		t.Error(err)
	}
}

func TestBinarizeFile(t *testing.T) {
	if err := BinarizeFile("test.png", "test_binarize.png", 139); err != nil {
		t.Fatal(err)
	}
}

func TestMedianBlurFile(t *testing.T) {
	if err := MedianBlurFile("test.png", "test_median.png", 1); err != nil {
		t.Fatal(err)
	}
}

func TestSobelFile(t *testing.T) {
	if err := SobelFile("test.png", "test_sobel.png", 400); err != nil {
		t.Fatal(err)
	}
}

func TestEqualizeHistFile(t *testing.T) {
	if err := EqualizeHistFile("test.png", "test_equalize_hist.png"); err != nil {
		t.Fatal(err)
	}
}

func TestOtsuThreshold(t *testing.T) {
	img, err := Open("test.png")
	if err != nil {
		t.Fatal(err)
	}
	threshold := OtsuThreshold(Grayscale(img))
	t.Log("Otsu threshold:", threshold)
}

func TestGenerateSolid(t *testing.T) {
	img := GenerateSolid(100, 100, color.RGBA{R: 255, G: 255, B: 0, A: 255})
	Save("solid.png", img, 100)
}
