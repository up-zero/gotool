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
