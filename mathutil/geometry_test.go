package mathutil

import (
	"github.com/up-zero/gotool/testutil"
	"testing"
)

func TestTranslatePolygon(t *testing.T) {
	points := []Point{
		{0, 0},
		{1, 0},
		{1, 1},
		{0, 1},
	}
	offset := Point{2, 3}
	result := TranslatePolygon(points, offset)
	testutil.Equal(t, result, []Point{{2, 3}, {3, 3}, {3, 4}, {2, 4}})
}

func TestPolygonArea(t *testing.T) {
	testutil.Equal(t, PolygonArea([]Point{{0, 0}, {1, 0}, {1, 1}, {0, 1}}), 1.0)
}

func TestGetAABB(t *testing.T) {
	aabb, _ := GetAABB([]Point{{0, 0}, {1, 0}, {1, 1}, {0.5, 1.5}, {0, 1}})
	testutil.Equal(t, aabb, Rectangle{Min: Point{0, 0}, Max: Point{1, 1.5}})
}
