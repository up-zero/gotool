package imageutil

import "image"

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
