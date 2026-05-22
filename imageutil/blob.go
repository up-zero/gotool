package imageutil

import (
	"image"
	"image/color"
)

// blobBitset 使用位图压缩访问标记
type blobBitset []uint64

func newBlobBitset(size int) blobBitset {
	return make(blobBitset, (size+63)>>6)
}

func (b blobBitset) Has(idx int) bool {
	return b[idx>>6]&(uint64(1)<<uint(idx&63)) != 0
}

func (b blobBitset) Set(idx int) {
	b[idx>>6] |= uint64(1) << uint(idx&63)
}

// blobSpan 表示同一行上的一个连续前景段。
//
// 相比“每个点都入队”的传统 BFS，按 span 扫描可以显著减少：
//   - 队列操作次数
//   - 8 邻域重复访问次数
//   - 大块连通区域上的无效判断
type blobSpan struct {
	y  int
	x1 int
	x2 int
}

// buildBlobForegroundChecker 根据常见图片类型进行快速前景判断
func buildBlobForegroundChecker(img image.Image, threshold uint8) func(x, y int) bool {
	bounds := img.Bounds()
	minX, minY := bounds.Min.X, bounds.Min.Y

	switch src := img.(type) {
	case *image.Gray:
		return func(x, y int) bool {
			offset := (y-minY)*src.Stride + (x - minX)
			return src.Pix[offset] > threshold
		}
	case *image.Gray16:
		return func(x, y int) bool {
			offset := (y-minY)*src.Stride + (x-minX)*2
			return src.Pix[offset] > threshold
		}
	case *image.Alpha:
		return func(x, y int) bool {
			offset := (y-minY)*src.Stride + (x - minX)
			return src.Pix[offset] > threshold
		}
	case *image.Alpha16:
		return func(x, y int) bool {
			offset := (y-minY)*src.Stride + (x-minX)*2
			return src.Pix[offset] > threshold
		}
	case *image.RGBA:
		return func(x, y int) bool {
			offset := (y-minY)*src.Stride + (x-minX)*4
			grayVal := (uint16(src.Pix[offset]) + uint16(src.Pix[offset+1]) + uint16(src.Pix[offset+2])) / 3
			return uint8(grayVal) > threshold
		}
	case *image.NRGBA:
		return func(x, y int) bool {
			offset := (y-minY)*src.Stride + (x-minX)*4
			grayVal := (uint16(src.Pix[offset]) + uint16(src.Pix[offset+1]) + uint16(src.Pix[offset+2])) / 3
			return uint8(grayVal) > threshold
		}
	case *image.RGBA64:
		return func(x, y int) bool {
			offset := (y-minY)*src.Stride + (x-minX)*8
			grayVal := (uint16(src.Pix[offset]) + uint16(src.Pix[offset+2]) + uint16(src.Pix[offset+4])) / 3
			return uint8(grayVal) > threshold
		}
	case *image.NRGBA64:
		return func(x, y int) bool {
			offset := (y-minY)*src.Stride + (x-minX)*8
			grayVal := (uint16(src.Pix[offset]) + uint16(src.Pix[offset+2]) + uint16(src.Pix[offset+4])) / 3
			return uint8(grayVal) > threshold
		}
	case *image.Paletted:
		var grayTable [256]uint8
		for i := range src.Palette {
			grayTable[i] = color.GrayModel.Convert(src.Palette[i]).(color.Gray).Y
		}
		return func(x, y int) bool {
			offset := (y-minY)*src.Stride + (x - minX)
			return grayTable[src.Pix[offset]] > threshold
		}
	case *image.YCbCr:
		return func(x, y int) bool {
			return src.Y[src.YOffset(x, y)] > threshold
		}
	default:
		return func(x, y int) bool {
			return color.GrayModel.Convert(img.At(x, y)).(color.Gray).Y > threshold
		}
	}
}

func normalizeBlobThreshold(threshold []uint8) uint8 {
	if len(threshold) > 0 {
		return threshold[0]
	}
	return 127
}

// scanBlobs 封装公共连通域扫描逻辑。
//
// collectPoints: 是否收集每个 Blob 的原始像素点
// keepAll:       是否保留全部 Blob 结果；否则仅保留最大 Blob
func scanBlobs(img image.Image, threshold uint8, collectPoints bool, keepAll bool) (*BlobResult, *Blob) {
	bounds := img.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	result := &BlobResult{Width: w, Height: h}
	if w == 0 || h == 0 {
		return result, nil
	}

	visited := newBlobBitset(w * h)
	baseX, baseY := bounds.Min.X, bounds.Min.Y
	isForeground := buildBlobForegroundChecker(img, threshold)
	blobID := 0
	var largest *Blob

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			idx := y*w + x

			if visited.Has(idx) {
				continue
			}

			if !isForeground(x+baseX, y+baseY) {
				visited.Set(idx)
				continue
			}

			blobID++
			currentBlob := Blob{ID: blobID}
			if collectPoints {
				currentBlob.Points = make([]image.Point, 0, 128)
			}

			queue := make([]blobSpan, 0, 64)
			head := 0
			minX, minY, maxX, maxY := x, y, x, y
			sumX, sumY := 0, 0
			count := 0

			// growSpan 从一个种子点向左右扩展，吃掉同一行上的整段前景。
			growSpan := func(seedX, spanY int) blobSpan {
				rowBase := spanY * w

				left := seedX
				for left > 0 {
					leftIdx := rowBase + left - 1
					if visited.Has(leftIdx) {
						break
					}
					if !isForeground(left-1+baseX, spanY+baseY) {
						visited.Set(leftIdx)
						break
					}
					left--
				}

				right := seedX
				for right+1 < w {
					rightIdx := rowBase + right + 1
					if visited.Has(rightIdx) {
						break
					}
					if !isForeground(right+1+baseX, spanY+baseY) {
						visited.Set(rightIdx)
						break
					}
					right++
				}

				for px := left; px <= right; px++ {
					pointIdx := rowBase + px
					if visited.Has(pointIdx) {
						continue
					}

					visited.Set(pointIdx)
					if collectPoints {
						currentBlob.Points = append(currentBlob.Points, image.Point{X: px, Y: spanY})
					}

					count++
					sumX += px
					sumY += spanY

					if px < minX {
						minX = px
					}
					if px > maxX {
						maxX = px
					}
					if spanY < minY {
						minY = spanY
					}
					if spanY > maxY {
						maxY = spanY
					}
				}

				span := blobSpan{y: spanY, x1: left, x2: right}
				queue = append(queue, span)
				return span
			}

			growSpan(x, y)

			for head < len(queue) {
				span := queue[head]
				head++

				for ny := span.y - 1; ny <= span.y+1; ny += 2 {
					if ny < 0 || ny >= h {
						continue
					}

					startX := span.x1 - 1
					if startX < 0 {
						startX = 0
					}
					endX := span.x2 + 1
					if endX >= w {
						endX = w - 1
					}

					for nx := startX; nx <= endX; {
						nIdx := ny*w + nx
						if visited.Has(nIdx) {
							nx++
							continue
						}

						if !isForeground(nx+baseX, ny+baseY) {
							visited.Set(nIdx)
							nx++
							continue
						}

						nextSpan := growSpan(nx, ny)
						nx = nextSpan.x2 + 1
					}
				}
			}

			currentBlob.Area = count
			currentBlob.Bounds = image.Rect(minX, minY, maxX+1, maxY+1)
			if count > 0 {
				currentBlob.Centroid = image.Point{X: sumX / count, Y: sumY / count}
			}

			if keepAll {
				result.Blobs = append(result.Blobs, currentBlob)
				continue
			}

			blobCopy := currentBlob
			if largest == nil || blobCopy.Area > largest.Area {
				largest = &blobCopy
			}
		}
	}

	return result, largest
}

// FindBlobs 查找 Mask 图片的连通区域
//
// # Params:
//
//	img: 输入的图片
//	threshold: 像素值大于此值被视为前景，默认：127
func FindBlobs(img image.Image, threshold ...uint8) *BlobResult {
	result, _ := scanBlobs(img, normalizeBlobThreshold(threshold), true, true)
	return result
}

// FindLargestBlob 查找面积最大的 Blob
//
// 与 FindBlobs(img).GetLargestBlob() 相比，有两个关键优化：
//   - 不构建完整的 BlobResult 列表
//   - 不收集每个像素的 Points
//
// 当只关心最大连通域时，这个方法会更省内存、也更快
//
// # Params:
//
//	img: 输入的图片
//	threshold: 像素值大于此值被视为前景，默认：127
func FindLargestBlob(img image.Image, threshold ...uint8) *Blob {
	_, largest := scanBlobs(img, normalizeBlobThreshold(threshold), false, false)
	return largest
}

// GetLargestBlob 获取最大的 Blob
func (r *BlobResult) GetLargestBlob() *Blob {
	if r == nil || len(r.Blobs) == 0 {
		return nil
	}

	// 默认第一个 Blob 为当前最大值
	maxIdx := 0
	for i := 1; i < len(r.Blobs); i++ {
		if r.Blobs[i].Area > r.Blobs[maxIdx].Area {
			maxIdx = i
		}
	}

	return &r.Blobs[maxIdx]
}
