package mathutil

import (
	"math"
)

// crossProduct 计算 向量 p1p2 和 p1p3 的叉积
//
//	> 0: p3 在 p1->p2 的左侧 (逆时针)
//	< 0: p3 在 p1->p2 的右侧 (顺时针)
//	= 0: p1, p2, p3 共线
func crossProduct(p1, p2, p3 Point) float64 {
	return (p2.X-p1.X)*(p3.Y-p1.Y) - (p2.Y-p1.Y)*(p3.X-p1.X)
}

// ConvexHull 计算凸包，基于 Jarvis 步进算法
func ConvexHull(points []Point) []Point {
	numPoints := len(points)
	if numPoints < 3 {
		return points
	}

	// 找到凸包的起始点 P0
	// 寻找 y 最小的点，如果 y 相同，则找 x 最小的点
	minYIndex := 0
	for i := 1; i < numPoints; i++ {
		if points[i].Y < points[minYIndex].Y ||
			(points[i].Y == points[minYIndex].Y && points[i].X < points[minYIndex].X) {
			minYIndex = i
		}
	}

	hull := make([]Point, 0)
	currentPointIndex := minYIndex

	for {
		// 将当前点加入凸包
		hull = append(hull, points[currentPointIndex])

		// 假设下一个点是第一个点（points[0]）
		nextPointIndex := 0
		if currentPointIndex == 0 {
			nextPointIndex = 1 // 避免选择自身
		}

		// 遍历所有点，寻找下一个顶点
		for i := 0; i < numPoints; i++ {
			// 如果 i 是当前点，则跳过
			if i == currentPointIndex {
				continue
			}

			pCurrent := points[currentPointIndex]
			pNextCandidate := points[nextPointIndex]
			pI := points[i]

			cross := crossProduct(pCurrent, pNextCandidate, pI)

			if cross > 0 {
				// pI 使得角度更大（更左侧），选择 pI 作为下一个点
				nextPointIndex = i
			} else if cross == 0 {
				// 三点共线，选择更远的点
				distSq1 := (pNextCandidate.X-pCurrent.X)*(pNextCandidate.X-pCurrent.X) +
					(pNextCandidate.Y-pCurrent.Y)*(pNextCandidate.Y-pCurrent.Y)
				distSq2 := (pI.X-pCurrent.X)*(pI.X-pCurrent.X) +
					(pI.Y-pCurrent.Y)*(pI.Y-pCurrent.Y)

				if distSq2 > distSq1 {
					nextPointIndex = i
				}
			}
		}

		currentPointIndex = nextPointIndex

		if currentPointIndex == minYIndex {
			break
		}
	}

	return hull
}

// perpendicularDistance 计算点 p 到线段的垂直距离
//
//   - 直线公式：(y2-y1)x + (x1-x2)y + (x2y1 - x1y2) = 0
//   - 点到直线的距离公式： |Ax + By + C| / sqrt(A^2 + B^2)
//
// # Params:
//
//	p: 待计算垂直距离的点
//	lineStart, lineEnd: 线段的两个端点
func perpendicularDistance(p, lineStart, lineEnd Point) float64 {
	p1 := lineStart
	p2 := lineEnd

	// 使用直线的一般方程 Ax + By + C = 0
	// A = y2 - y1
	A := p2.Y - p1.Y
	// B = x1 - x2
	B := p1.X - p2.X
	// C = x2*y1 - x1*y2
	C := p2.X*p1.Y - p1.X*p2.Y

	// 点 (xp, yp) 到直线的距离公式： |A*xp + B*yp + C| / sqrt(A^2 + B^2)
	num := math.Abs(A*p.X + B*p.Y + C)
	den := math.Sqrt(A*A + B*B)

	if den == 0 {
		// 如果分母为 0，说明 lineStart 和 lineEnd 是同一个点
		// 此时，距离为 p 到这个点的欧几里得距离
		dx := p.X - p1.X
		dy := p.Y - p1.Y
		return math.Sqrt(dx*dx + dy*dy)
	}

	return num / den
}

// rdpRecursive Ramer-Douglas-Peucker 算法的递归实现
//
// # Params:
//
//	points: 输入的顶点列表
//	epsilon: 阈值，值越大，简化程度越高
func rdpRecursive(points []Point, epsilon float64) []Point {
	n := len(points)
	if n < 2 {
		return points
	}

	// 取起点和终点
	start := points[0]
	end := points[n-1]

	// 寻找距离线段 (start, end) 最远的点
	maxDist := -1.0
	maxIndex := 0

	// 遍历 1 到 n-2 之间的所有点
	for i := 1; i < n-1; i++ {
		dist := perpendicularDistance(points[i], start, end)
		if dist > maxDist {
			maxDist = dist
			maxIndex = i
		}
	}

	if maxDist > epsilon {
		// 距离 > 阈值，保留这个最远点
		// 递归地简化左右两部分
		left := rdpRecursive(points[0:maxIndex+1], epsilon)
		right := rdpRecursive(points[maxIndex:n], epsilon)

		// 组合结果
		// left 的最后一个点和 right 的第一个点都是 maxIndex
		// 所以我们去掉 left 的最后一个点，然后拼接 right
		return append(left[:len(left)-1], right...)

	} else {
		// 距离 <= 阈值，所有中间点都被丢弃
		return []Point{start, end}
	}
}

// SimplifyPath 简化路径，使用 Ramer-Douglas-Peucker (RDP) 算法进行简化
//
// # Params:
//
//	points: 输入的顶点列表
//	epsilon: 阈值 (点到线段的距离)，值越大，简化程度越高
func SimplifyPath(points []Point, epsilon float64) []Point {
	if len(points) < 3 {
		return points
	}
	return rdpRecursive(points, epsilon)
}

// sub 向量减法
func (p Point) sub(q Point) Point {
	return Point{
		X: p.X - q.X,
		Y: p.Y - q.Y,
	}
}

// add 向量加法
func (p Point) add(q Point) Point {
	return Point{
		X: p.X + q.X,
		Y: p.Y + q.Y,
	}
}

// scale 向量缩放
func (p Point) scale(f float64) Point {
	return Point{
		X: p.X * f,
		Y: p.Y * f,
	}
}

// length 向量的长度
func (p Point) length() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

// normalize 向量归一化
func (p Point) normalize() Point {
	l := p.length()
	if l == 0 {
		return Point{0, 0}
	}
	return p.scale(1.0 / l)
}

// perpendicular 获取向量的左手法线
func (p Point) perpendicular() Point {
	return Point{-p.Y, p.X}
}

// offset 直线平移 margin 距离
func (l Line) offset(margin float64) Line {
	v := l.P2.sub(l.P1)
	normal := v.perpendicular().normalize()
	offsetVec := normal.scale(margin)

	return Line{
		P1: l.P1.add(offsetVec),
		P2: l.P2.add(offsetVec),
	}
}

// lineIntersection 计算两条直线的交点
//
// # Params:
//
//	l2: 另一条直线
//
// # Returns:
//
//	Point: 交点坐标
//	bool: 是否有交点
func (l Line) lineIntersection(l2 Line) (Point, bool) {
	x1, y1 := l.P1.X, l.P1.Y
	x2, y2 := l.P2.X, l.P2.Y
	x3, y3 := l2.P1.X, l2.P1.Y
	x4, y4 := l2.P2.X, l2.P2.Y

	// 标准直线交点公式 (基于行列式)
	den := (x1-x2)*(y3-y4) - (y1-y2)*(x3-x4)
	if den == 0 {
		return Point{}, false // 直线平行
	}

	// 分子
	tNum := (x1-x3)*(y3-y4) - (y1-y3)*(x3-x4)
	// 参数 t
	t := tNum / den

	// 计算交点
	ix := x1 + t*(x2-x1)
	iy := y1 + t*(y2-y1)

	return Point{ix, iy}, true
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
func OffsetPolygon(points []Point, margin float64) []Point {
	n := len(points)
	if n < 3 {
		return points
	}

	// 使用 Shoelace (鞋带) 公式计算多边形有向面积
	area := 0.0
	for i := 0; i < n; i++ {
		p1 := points[i]
		p2 := points[(i+1)%n]
		area += (p1.X * p2.Y) - (p2.X * p1.Y)
	}
	// 在 Y 轴向下的图像坐标系中:
	// area > 0 是顺时针 (Clockwise, CW)
	// area < 0 是逆时针 (Counter-Clockwise, CCW)
	if area > 0 { // 顺时针
		margin = -margin
	}

	// 计算每条边的偏移直线
	offsetLines := make([]Line, n)
	for i := 0; i < n; i++ {
		p1 := points[i]
		p2 := points[(i+1)%n]
		edge := Line{p1, p2}
		offsetLines[i] = edge.offset(margin)
	}

	// 计算偏移直线的交点
	newPoints := make([]Point, n)
	for i := 0; i < n; i++ {
		// 当前边的偏移线
		line1 := offsetLines[i]
		// 上一条边的偏移线
		line0 := offsetLines[(i-1+n)%n]

		// 计算交点
		newPoint, ok := line0.lineIntersection(line1)

		if !ok {
			// 发生平行（通常因为输入多边形有 180° 的共线点）
			newPoint = line1.P1
		}

		newPoints[i] = newPoint
	}

	return newPoints
}

// TranslatePolygon 多边形平移，按指定向量移动多边形
//
// # Params:
//
//	points: 多边形顶点切片
//	offset: 平移向量
func TranslatePolygon(points []Point, offset Point) []Point {
	if len(points) == 0 {
		return []Point{}
	}

	newPoints := make([]Point, len(points))
	for i, p := range points {
		newPoints[i] = p.add(offset)
	}

	return newPoints
}

// PolygonArea 计算多边形的面积，基于鞋带公式
//
//	Sum = (x1*y2 - y1*x2) + (x2*y3 - y2*x3) + ... + (xn*y1 - yn*x1)
//	Area = 0.5 * |Sum|
func PolygonArea(points []Point) float64 {
	n := len(points)
	if n < 3 {
		return 0.0
	}

	var area = 0.0

	for i := 0; i < n; i++ {
		v1 := points[i]
		v2 := points[(i+1)%n]
		area += (v1.X * v2.Y) - (v2.X * v1.Y)
	}

	return math.Abs(area / 2.0)
}
