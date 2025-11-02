package imageutil

import (
	"fmt"
	"image"
	"image/color"
	"math/rand"
	"testing"
	"time"
)

func getRandPoints() ([]image.Point, *image.RGBA) {
	img := image.NewRGBA(image.Rect(0, 0, 1920, 1080))
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	points := make([]image.Point, 0)
	for i := 0; i < 200; i++ {
		point := image.Point{X: r.Intn(1920), Y: r.Intn(1080)}
		points = append(points, point)
		DrawFilledCircle(img, point, 5, color.Black)
	}
	return points, img
}

func TestConvexHull(t *testing.T) {
	points, img := getRandPoints()
	hullPoints := ConvexHull(points)
	DrawThickPolygonOutline(img, hullPoints, 5, color.Black)
	Save("convex_hull.png", img, 100)
}

func TestSimplifyPolygon(t *testing.T) {
	points, img := getRandPoints()
	hullPoints := ConvexHull(points)
	fmt.Printf("Convex hull points length: %v\n", len(hullPoints))

	hullPoints = SimplifyPath(hullPoints, 10)
	fmt.Printf("Simplified convex hull points length: %v\n", len(hullPoints))
	DrawThickPolygonOutline(img, hullPoints, 5, color.Black)

	Save("simplify_path.png", img, 100)
}

func TestOffsetPolygon(t *testing.T) {
	points, img := getRandPoints()
	hullPoints := ConvexHull(points)
	DrawThickPolygonOutline(img, hullPoints, 5, color.Black)

	offsetPoints := OffsetPolygon(hullPoints, -10)
	fmt.Printf("Offset polygon points : %v\n", offsetPoints)
	DrawThickPolygonOutline(img, offsetPoints, 5, color.RGBA{R: 255, G: 0, B: 0, A: 255})

	Save("offset_polygon.png", img, 100)
}
