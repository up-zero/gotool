package mediautil

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
