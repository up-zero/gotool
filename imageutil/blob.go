package imageutil

import (
	"image"
)

// FindBlobs 查找 Mask 图片的连通区域
//
// # Params:
//
//	img: 输入的图片
//	threshold: 像素值大于此值被视为前景，默认：127
func FindBlobs(img image.Image, threshold ...uint8) *BlobResult {
	bounds := img.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	bThreshold := uint8(127)
	if len(threshold) > 0 {
		bThreshold = threshold[0]
	}

	// 初始化访问标记矩阵
	visited := make([]bool, w*h)

	result := &BlobResult{Width: w, Height: h}
	blobID := 0
	grayImg, isGray := img.(*image.Gray)

	// 判断像素是否为前景
	isForeground := func(x, y int) bool {
		// 边界检查
		if x < 0 || x >= w || y < 0 || y >= h {
			return false
		}

		// 性能分支
		if isGray {
			// 直接访问 Pix 数组，计算偏移量: y * Stride + x
			offset := (y-bounds.Min.Y)*grayImg.Stride + (x - bounds.Min.X)
			return grayImg.Pix[offset] > bThreshold
		} else {
			// 通用接口调用，性能稍慢
			r, g, b, _ := img.At(x+bounds.Min.X, y+bounds.Min.Y).RGBA()
			grayVal := uint8((r + g + b) / 3 >> 8)
			return grayVal > bThreshold
		}
	}

	// 8-邻域方向数组
	dirs := []image.Point{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}

	// 遍历全图
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			idx := y*w + x

			// 如果已访问或者是背景，跳过
			if visited[idx] {
				continue
			}

			// 检查是否为前景点
			if !isForeground(x, y) {
				continue
			}

			// 发现新的 Blob，开始 BFS
			blobID++
			currentBlob := Blob{ID: blobID}

			// 队列初始化
			queue := []image.Point{{x, y}}
			visited[idx] = true

			// 统计变量初始化
			minX, minY, maxX, maxY := x, y, x, y
			sumX, sumY := 0, 0
			count := 0

			for len(queue) > 0 {
				p := queue[0]
				queue = queue[1:]

				// 记录点信息
				currentBlob.Points = append(currentBlob.Points, p)

				count++
				sumX += p.X
				sumY += p.Y

				// 更新 AABB 边界
				if p.X < minX {
					minX = p.X
				}
				if p.X > maxX {
					maxX = p.X
				}
				if p.Y < minY {
					minY = p.Y
				}
				if p.Y > maxY {
					maxY = p.Y
				}

				// 搜索 8 邻域
				for _, d := range dirs {
					nx, ny := p.X+d.X, p.Y+d.Y
					nIdx := ny*w + nx

					// 边界检查 + 访问检查
					if nx >= 0 && nx < w && ny >= 0 && ny < h && !visited[nIdx] {
						if isForeground(nx, ny) {
							visited[nIdx] = true
							queue = append(queue, image.Point{X: nx, Y: ny})
						}
					}
				}
			}

			// 填充 Blob 统计特征
			currentBlob.Area = count
			currentBlob.Bounds = image.Rect(minX, minY, maxX+1, maxY+1)
			if count > 0 {
				currentBlob.Centroid = image.Point{X: sumX / count, Y: sumY / count}
			}

			result.Blobs = append(result.Blobs, currentBlob)
		}
	}

	return result
}
