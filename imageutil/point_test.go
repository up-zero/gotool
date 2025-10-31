package imageutil

import (
	"image"
	"image/color"
	"math/rand"
	"testing"
	"time"
)

func TestConvexHull(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 1920, 1080))
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	points := make([]image.Point, 0)
	for i := 0; i < 200; i++ {
		point := image.Point{X: r.Intn(1920), Y: r.Intn(1080)}
		points = append(points, point)
		img.Set(point.X, point.Y, image.Black)
		DrawFilledCircle(img, point, 5, color.Black)
	}
	hullPoints := ConvexHull(points)
	DrawThickPolygonOutline(img, hullPoints, 5, color.Black)
	Save("convex_hull.png", img, 100)
}
