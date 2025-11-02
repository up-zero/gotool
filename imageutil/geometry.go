package imageutil

import (
	"github.com/up-zero/gotool/mathutil"
	"image"
)

func toImagePoint(p mathutil.Point) image.Point {
	round := func(f float64) int {
		if f < 0 {
			return 1
		}
		return int(f + 0.5)
	}
	return image.Point{
		X: round(p.X),
		Y: round(p.Y),
	}
}

func toImagePoints(points []mathutil.Point) []image.Point {
	result := make([]image.Point, len(points))
	for i, pt := range points {
		result[i] = toImagePoint(pt)
	}
	return result
}

func toMathPoints(points []image.Point) []mathutil.Point {
	result := make([]mathutil.Point, len(points))
	for i, pt := range points {
		result[i] = mathutil.Point{
			X: float64(pt.X),
			Y: float64(pt.Y),
		}
	}
	return result
}

// ConvexHull 计算凸包，基于 Jarvis 步进算法
func ConvexHull(points []image.Point) []image.Point {
	return toImagePoints(mathutil.ConvexHull(toMathPoints(points)))
}

// SimplifyPath 简化路径，使用 Ramer-Douglas-Peucker (RDP) 算法进行简化
//
// # Params:
//
//	points: 输入的顶点列表
//	epsilon: 阈值 (点到线段的距离)，值越大，简化程度越高
func SimplifyPath(points []image.Point, epsilon float64) []image.Point {
	return toImagePoints(mathutil.SimplifyPath(toMathPoints(points), epsilon))
}

// OffsetPolygon 多边形偏移 (内缩/外扩)
//
//   - margin > 0: 外扩
//   - margin < 0: 内缩
//
// # Params:
//
//	points: 输入的顶点列表
//	margin: 内外缩的距离
func OffsetPolygon(points []image.Point, margin float64) []image.Point {
	return toImagePoints(mathutil.OffsetPolygon(toMathPoints(points), margin))
}
