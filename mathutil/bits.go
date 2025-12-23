package mathutil

import "math/bits"

// HammingDistance 汉明距离，计算两个 64 位无符号整数之间二进制位不同的个数
func HammingDistance(h1, h2 uint64) int {
	return bits.OnesCount64(h1 ^ h2)
}
