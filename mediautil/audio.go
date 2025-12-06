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

// MelFilters 生成 Mel 滤波器组权重矩阵，将线性频率映射到Mel刻度
//
// # Params:
//
//	sampleRate: 采样率
//	fftSize: FFT 窗口大小
//	melBinCount: Mel 频带数量
func MelFilters(sampleRate, fftSize, melBinCount int) [][]float32 {
	fMin := 20.0
	fMax := float64(sampleRate) / 2.0 // Nyquist 频率

	// 内部闭包：Hz 转 Mel
	hzToMel := func(hz float64) float64 {
		return 1127.0 * math.Log(1.0+hz/700.0)
	}
	// 内部闭包：Mel 转 Hz
	melToHz := func(mel float64) float64 {
		return 700.0 * (math.Exp(mel/1127.0) - 1.0)
	}

	melMin := hzToMel(fMin)
	melMax := hzToMel(fMax)

	// 计算所有滤波器的中心频率
	melPoints := make([]float64, melBinCount+2)
	binPoints := make([]int, melBinCount+2)

	step := (melMax - melMin) / float64(melBinCount+1)

	for i := 0; i < melBinCount+2; i++ {
		melPoints[i] = melMin + float64(i)*step
		hz := melToHz(melPoints[i])
		// 将 Hz 映射到 FFT bin 索引
		// bin = floor((N+1) * hz / sampleRate)
		binPoints[i] = int(math.Floor((float64(fftSize) + 1) * hz / float64(sampleRate)))
	}

	filters := make([][]float32, melBinCount)
	for i := 0; i < melBinCount; i++ {
		filters[i] = make([]float32, fftSize/2+1)
		start := binPoints[i]
		center := binPoints[i+1]
		end := binPoints[i+2]

		// 构建三角形滤波器
		for j := start; j < center; j++ {
			filters[i][j] = float32(j-start) / float32(center-start)
		}
		for j := center; j < end; j++ {
			filters[i][j] = 1.0 - float32(j-center)/float32(end-center)
		}
	}
	return filters
}

// ApplyCMVN 倒谱均值方差归一化 (Cepstral Mean and Variance Normalization)
//
// 公式: result = (x + negMean) * invStd
//
// # Params:
//
// features: 特征矩阵
// negMean: 负均值向量
// invStd: 逆标准差向量
func ApplyCMVN(features [][]float32, negMean []float32, invStd []float32) {
	for i := 0; i < len(features); i++ {
		dim := len(features[i])
		// 安全检查，防止维度不匹配越界
		checkLen := len(negMean)
		if checkLen > dim {
			checkLen = dim
		}
		if len(invStd) < checkLen {
			checkLen = len(invStd)
		}

		for j := 0; j < checkLen; j++ {
			features[i][j] = (features[i][j] + negMean[j]) * invStd[j]
		}
	}
}
