package imageutil

import (
	"image"
	"math"
)

// crossProduct 计算 向量 p1p2 和 p1p3 的叉积
//
//	> 0: p3 在 p1->p2 的左侧 (逆时针)
//	< 0: p3 在 p1->p2 的右侧 (顺时针)
//	= 0: p1, p2, p3 共线
func crossProduct(p1, p2, p3 image.Point) int {
	return (p2.X-p1.X)*(p3.Y-p1.Y) - (p2.Y-p1.Y)*(p3.X-p1.X)
}

// ConvexHull 计算凸包，基于 Jarvis 步进算法
func ConvexHull(points []image.Point) []image.Point {
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

	hull := make([]image.Point, 0)
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
func perpendicularDistance(p, lineStart, lineEnd image.Point) float64 {
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
	num := math.Abs(float64(A*p.X + B*p.Y + C))
	den := math.Sqrt(float64(A*A + B*B))

	if den == 0 {
		// 如果分母为 0，说明 lineStart 和 lineEnd 是同一个点
		// 此时，距离为 p 到这个点的欧几里得距离
		dx := p.X - p1.X
		dy := p.Y - p1.Y
		return math.Sqrt(float64(dx*dx + dy*dy))
	}

	return num / den
}

// rdpRecursive Ramer-Douglas-Peucker 算法的递归实现
//
// # Params:
//
//	points: 输入的顶点列表
//	epsilon: 阈值，值越大，简化程度越高
func rdpRecursive(points []image.Point, epsilon float64) []image.Point {
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
		return []image.Point{start, end}
	}
}

// SimplifyPath 简化路径，使用 Ramer-Douglas-Peucker (RDP) 算法进行简化
//
// # Params:
//
//	points: 输入的顶点列表
//	epsilon: 阈值 (点到线段的距离)，值越大，简化程度越高
func SimplifyPath(points []image.Point, epsilon float64) []image.Point {
	if len(points) < 3 {
		return points
	}

	return rdpRecursive(points, epsilon)
}
