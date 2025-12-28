package imageutil

import (
	"image"
	"math"
)

// AHash 平均哈希
func AHash(img image.Image) uint64 {
	grayImg := Grayscale(Resize(img, 8, 8))

	var sum uint64
	pixels := make([]uint8, 64)

	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			r := grayImg.GrayAt(x, y).Y
			pixels[y*8+x] = r
			sum += uint64(r)
		}
	}
	avg := uint8(sum / 64)

	// 生成 64 位哈希：大于平均值为 1，小于为 0
	var hash uint64
	for i, p := range pixels {
		if p >= avg {
			hash |= 1 << uint(i)
		}
	}
	return hash
}

// DHash 差异哈希
func DHash(img image.Image) uint64 {
	grayImg := Grayscale(Resize(img, 9, 8))

	var hash uint64
	bitPos := 0

	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			left := grayImg.GrayAt(x, y).Y
			right := grayImg.GrayAt(x+1, y).Y

			// 如果左边亮度 >= 右边，则将对应位置的 bit 设为 1
			if left >= right {
				hash |= 1 << uint(bitPos)
			}
			bitPos++
		}
	}

	return hash
}

// PHash 感知哈希
func PHash(img image.Image) uint64 {
	gray := Grayscale(Resize(img, 32, 32))

	// 转换为二维 float64 矩阵
	pixels := make([][]float64, 32)
	for i := 0; i < 32; i++ {
		pixels[i] = make([]float64, 32)
		for j := 0; j < 32; j++ {
			pixels[i][j] = float64(gray.GrayAt(i, j).Y)
		}
	}

	// 执行 DCT 变换
	dctMatrix := applyDCT(pixels, 32)

	// 计算 8x8 区域的低频分量均值
	var sum float64
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			sum += dctMatrix[y][x]
		}
	}
	avg := sum / 64.0

	// 生成哈希
	var hash uint64
	bitPos := 0
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			if dctMatrix[y][x] > avg {
				hash |= 1 << uint(bitPos)
			}
			bitPos++
		}
	}

	return hash
}

// applyDCT 二维离散余弦变换
func applyDCT(f [][]float64, n int) [][]float64 {
	F := make([][]float64, n)
	for i := range F {
		F[i] = make([]float64, n)
	}

	pi := math.Pi
	for u := 0; u < n; u++ {
		for v := 0; v < n; v++ {
			var sum float64
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					sum += f[i][j] *
						math.Cos(float64(2*i+1)*float64(u)*pi/float64(2*n)) *
						math.Cos(float64(2*j+1)*float64(v)*pi/float64(2*n))
				}
			}
			// 归一化系数
			cu := 1.0
			if u == 0 {
				cu = 1.0 / math.Sqrt(2)
			}
			cv := 1.0
			if v == 0 {
				cv = 1.0 / math.Sqrt(2)
			}

			F[u][v] = 0.25 * cu * cv * sum
		}
	}
	return F
}
