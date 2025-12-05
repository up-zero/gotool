package mediautil

import (
	"math"
	"math/cmplx"
)

// PreEmphasis 预加重滤波器，提升高频部分，平衡频谱能量
//
// 差分公式: y[t] = x[t] - alpha * x[t-1]
//
// # Params:
//
//	samples: 输入的音频数据
//	alpha: 预加重系数，推荐 0.97
func PreEmphasis(samples []float32, alpha float32) []float32 {
	if len(samples) == 0 {
		return nil
	}
	output := make([]float32, len(samples))
	output[0] = samples[0]
	for i := 1; i < len(samples); i++ {
		output[i] = samples[i] - alpha*samples[i-1]
	}
	return output
}

// HammingWindow 生成汉明窗，减少频谱泄漏
//
// Hamming 公式: 0.54 - 0.46 * cos(2πn / (N-1))
//
// # Params:
//
//	size: 窗口大小
func HammingWindow(size int) []float32 {
	window := make([]float32, size)
	for i := 0; i < size; i++ {
		window[i] = float32(0.54 - 0.46*math.Cos(2*math.Pi*float64(i)/float64(size-1)))
	}
	return window
}

// FFT 快速傅里叶变换，时域转频域
//
// # Params:
//
//	x: 时域数据（波形）
func FFT(x []complex128) []complex128 {
	n := len(x)
	if n <= 1 {
		return x
	}

	// 分治
	even := make([]complex128, n/2)
	odd := make([]complex128, n/2)
	for i := 0; i < n/2; i++ {
		even[i] = x[2*i]
		odd[i] = x[2*i+1]
	}

	even = FFT(even)
	odd = FFT(odd)

	// 合并
	t := make([]complex128, n)
	for k := 0; k < n/2; k++ {
		angle := -2 * math.Pi * float64(k) / float64(n)
		w := cmplx.Exp(complex(0, angle)) * odd[k]
		t[k] = even[k] + w
		t[k+n/2] = even[k] - w
	}
	return t
}
