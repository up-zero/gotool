package imageutil

import (
	"image"
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
