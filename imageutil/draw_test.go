package imageutil

import (
	"image"
	"image/color"
	"math/rand"
	"testing"
	"time"
)

func TestDrawFilledCircle(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 1920, 1080))
	DrawFilledCircle(img, image.Point{X: 960, Y: 540}, 600, color.Black)
	Save("test_filled_circle.png", img, 100)
}

func TestDrawThickLine(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	img := image.NewRGBA(image.Rect(0, 0, 1980, 1024))
	for i := 0; i < 100; i++ {
		p1 := image.Point{X: r.Intn(img.Bounds().Dx()), Y: r.Intn(img.Bounds().Dy())}
		p2 := image.Point{X: r.Intn(img.Bounds().Dx()), Y: r.Intn(img.Bounds().Dy())}
		DrawThickLine(img, p1, p2, 10, color.RGBA{R: uint8(r.Intn(255)), G: uint8(r.Intn(255)), B: uint8(r.Intn(255)), A: 255})
	}
	Save("test_draw_thick_line.png", img, 100)
}

func TestDrawLine(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	img := image.NewRGBA(image.Rect(0, 0, 1980, 1024))
	for i := 0; i < 100; i++ {
		p1 := image.Point{X: r.Intn(img.Bounds().Dx()), Y: r.Intn(img.Bounds().Dy())}
		p2 := image.Point{X: r.Intn(img.Bounds().Dx()), Y: r.Intn(img.Bounds().Dy())}
		DrawLine(img, p1, p2, color.RGBA{R: uint8(r.Intn(255)), G: uint8(r.Intn(255)), B: uint8(r.Intn(255)), A: 255})
	}
	Save("test_draw_line.png", img, 100)
}
