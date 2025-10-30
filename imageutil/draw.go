package imageutil

import (
	"github.com/up-zero/gotool/mathutil"
	"image"
	"image/color"
	"image/draw"
	"sort"
)

// DrawFilledCircle 绘制填充的圆形
//
// # Params:
//
//	dst: 目标图片
//	center: 圆心坐标
//	radius: 圆的半径
//	c: 填充颜色
func DrawFilledCircle(dst draw.Image, center image.Point, radius int, c color.Color) {
	// 计算圆形的边界框
	bounds := dst.Bounds()
	x0 := center.X - radius
	y0 := center.Y - radius
	x1 := center.X + radius
	y1 := center.Y + radius

	// 迭代边界框内的每个像素
	for y := y0; y <= y1; y++ {
		for x := x0; x <= x1; x++ {
			// 检查 (x, y) 是否在图像边界内
			if x < bounds.Min.X || x >= bounds.Max.X || y < bounds.Min.Y || y >= bounds.Max.Y {
				continue
			}

			// 计算 (x, y) 到圆心的距离的平方
			dx := x - center.X
			dy := y - center.Y
			distSq := dx*dx + dy*dy

			if distSq <= radius*radius {
				dst.Set(x, y, c)
			}
		}
	}
}

// DrawThickLine 绘制粗线
//
// # Params:
//
//	dst: 目标图片
//	p1, p2: 直线的起点和终点
//	thickness: 线宽
//	c: 颜色
func DrawThickLine(dst draw.Image, p1, p2 image.Point, thickness int, c color.Color) {
	if thickness <= 0 {
		return
	}

	// 笔刷的半径
	radius := thickness / 2

	// Bresenham 算法参数
	x1, y1 := p1.X, p1.Y
	x2, y2 := p2.X, p2.Y

	dx := mathutil.Abs(x2 - x1)
	dy := mathutil.Abs(y2 - y1)
	sx, sy := 1, 1

	if x1 > x2 {
		sx = -1
	}
	if y1 > y2 {
		sy = -1
	}
	err := dx - dy

	for {
		if thickness == 1 {
			dst.Set(x1, y1, c)
		} else {
			DrawFilledCircle(dst, image.Point{X: x1, Y: y1}, radius, c)
		}
		if x1 == x2 && y1 == y2 {
			break
		}
		// 计算下一个点
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x1 += sx
		}
		if e2 < dx {
			err += dx
			y1 += sy
		}
	}
}

// DrawLine 绘制直线，基于布雷森汉姆（Bresenham）直线算法
//
// # Params:
//
//	dst: 目标图片
//	p1, p2: 直线的起点和终点
//	color: 颜色
func DrawLine(dst draw.Image, p1, p2 image.Point, color color.Color) {
	DrawThickLine(dst, p1, p2, 1, color)
}

// DrawRectOutline 绘制矩形边框
//
// # Params:
//
//	dst: 目标图片，必须是可变的（例如 *image.RGBA）
//	r: 要绘制的矩形区域
//	c: 颜色
func DrawRectOutline(dst draw.Image, r image.Rectangle, c color.Color) {
	bounds := dst.Bounds()
	r = r.Intersect(bounds)
	if r.Empty() {
		return
	}

	// 绘制水平线 (Top & Bottom)
	for x := r.Min.X; x < r.Max.X; x++ {
		dst.Set(x, r.Min.Y, c)   // Top
		dst.Set(x, r.Max.Y-1, c) // Bottom
	}

	// 绘制垂直线 (Left & Right)
	for y := r.Min.Y; y < r.Max.Y; y++ {
		dst.Set(r.Min.X, y, c)   // Left
		dst.Set(r.Max.X-1, y, c) // Right
	}
}

// DrawFilledRect 矩形填充
//
// # Params:
//
//	dst: 目标图片
//	r: 要绘制的矩形区域
//	c: 颜色
func DrawFilledRect(dst draw.Image, r image.Rectangle, c color.Color) {
	src := image.NewUniform(c)
	r = r.Intersect(dst.Bounds())

	draw.Draw(dst, r, src, image.Point{}, draw.Src)
}

// DrawThickRectOutline 绘制粗矩形边框
//
// # Params:
//
//	dst: 目标图片
//	r: 要绘制的矩形区域
//	c: 颜色
//	thickness: 边框的粗细，像素
func DrawThickRectOutline(dst draw.Image, r image.Rectangle, c color.Color, thickness int) {
	bounds := dst.Bounds()
	if thickness <= 0 {
		return
	} else if thickness == 1 {
		DrawRectOutline(dst, r, c)
		return
	}
	r = r.Intersect(bounds)
	if r.Empty() {
		return
	}

	// 处理粗细过大的情况
	if thickness*2 >= r.Dx() || thickness*2 >= r.Dy() {
		DrawFilledRect(dst, r, c)
		return
	}

	// Top Bar
	// (Min.X, Min.Y) -> (Max.X, Min.Y + thickness)
	topRect := image.Rect(r.Min.X, r.Min.Y, r.Max.X, r.Min.Y+thickness)
	DrawFilledRect(dst, topRect, c)

	// Bottom Bar
	// (Min.X, Max.Y - thickness) -> (Max.X, Max.Y)
	bottomRect := image.Rect(r.Min.X, r.Max.Y-thickness, r.Max.X, r.Max.Y)
	DrawFilledRect(dst, bottomRect, c)

	// Left Bar
	// (Min.X, Min.Y + thickness) -> (Min.X + thickness, Max.Y - thickness)
	leftRect := image.Rect(r.Min.X, r.Min.Y+thickness, r.Min.X+thickness, r.Max.Y-thickness)
	DrawFilledRect(dst, leftRect, c)

	// Right Bar
	// (Max.X - thickness, Min.Y + thickness) -> (Max.X, Max.Y - thickness)
	rightRect := image.Rect(r.Max.X-thickness, r.Min.Y+thickness, r.Max.X, r.Max.Y-thickness)
	DrawFilledRect(dst, rightRect, c)
}

// DrawPolygonOutline 绘制多边形边框
//
// # Params:
//
//	dst: 目标图片
//	points: 顶点集合
//	c: 颜色
func DrawPolygonOutline(dst draw.Image, points []image.Point, c color.Color) {
	DrawThickPolygonOutline(dst, points, 1, c)
}

// DrawThickPolygonOutline 绘制粗多边形边框
//
// # Params:
//
//	dst: 目标图片
//	points: 顶点集合
//	thickness: 边框的粗细，像素
//	c: 颜色
func DrawThickPolygonOutline(dst draw.Image, points []image.Point, thickness int, c color.Color) {
	numPoints := len(points)
	if numPoints < 2 || thickness <= 0 {
		return
	}

	for i := 0; i < numPoints; i++ {
		p1 := points[i]
		p2 := points[(i+1)%numPoints]

		DrawThickLine(dst, p1, p2, thickness, c)
	}
}

// DrawFilledPolygon 多边形填充，使用扫描线算法填充多边形
//
// # Params:
//
//	dst: 目标图片
//	points: 顶点集合
//	c: 颜色
func DrawFilledPolygon(dst draw.Image, points []image.Point, c color.Color) {
	numPoints := len(points)
	if numPoints < 3 {
		return
	}

	bounds := dst.Bounds()

	// 找到 Y 轴的最小和最大边界
	minY := bounds.Max.Y
	maxY := bounds.Min.Y
	for _, p := range points {
		minY = min(minY, p.Y)
		maxY = max(maxY, p.Y)
	}

	// 裁剪 Y 边界到图像边界内
	minY = max(minY, bounds.Min.Y)
	maxY = min(maxY, bounds.Max.Y)

	// 遍历每一条扫描线 y
	for y := minY; y <= maxY; y++ {
		intersections := make([]int, 0)

		// 遍历多边形的所有边，计算交点坐标
		for i := 0; i < numPoints; i++ {
			p1 := points[i]
			p2 := points[(i+1)%numPoints]

			// p1=(x1, y1), p2=(x2, y2)
			y1, y2 := p1.Y, p2.Y
			x1, x2 := p1.X, p2.X

			// 确保 y1 <= y2
			if y1 > y2 {
				y1, y2 = y2, y1
				x1, x2 = x2, x1
			}

			// 检查边是否跨越了扫描线 y
			if y1 < y && y2 >= y {
				// 特殊处理：如果 y2 == y，且 p2 不是水平边（即 p2.Y > p3.Y），则排除
				// 避免双重计数
				if y2 == y {
					p3 := points[(i+2)%numPoints]
					if p2.Y > p3.Y {
						continue
					}
				}

				// 计算交点的 x 坐标：使用线性插值
				// x = x1 + (y - y1) * (x2 - x1) / (y2 - y1)
				// 使用 float64 进行精确计算
				xIntersect := float64(x1) + float64(y-y1)*float64(x2-x1)/float64(y2-y1)

				// 将交点 x 坐标转换为整数并添加
				intersections = append(intersections, int(xIntersect))
			}
		}

		sort.Ints(intersections)

		// 将交点两两配对并填充
		for i := 0; i+1 < len(intersections); i += 2 {
			xStart := max(intersections[i], bounds.Min.X)
			xEnd := min(intersections[i+1], bounds.Max.X)

			// 绘制水平线段 (从 xStart 到 xEnd - 1)
			for x := xStart; x < xEnd; x++ {
				dst.Set(x, y, c)
			}
		}
	}
}
