package mediautil

import "math"

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
