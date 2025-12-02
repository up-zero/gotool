package mediautil

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/up-zero/gotool"
	"io"
	"math"
	"os"
	"time"
)

// 常见错误定义
var (
	// ErrNotWavFile 不是有效的 WAV 文件
	ErrNotWavFile = errors.New("not a valid RIFF/WAVE file")
	// ErrUnsupportedBitDepth 不支持的位深
	ErrUnsupportedBitDepth = errors.New("unsupported bit depth: only 16, 24, 32 are supported")
)

const (
	// SampleRate8K 采样率 8KHz
	SampleRate8K = 8000
	// SampleRate16K 采样率 16KHz
	SampleRate16K = 16000
	// SampleRate44K 采样率 44.1KHz
	SampleRate44K = 44100
	// SampleRate48K 采样率 48KHz
	SampleRate48K = 48000
	// SampleRate96K 采样率 96KHz
	SampleRate96K = 96000

	// BitsPerSample16 位深 16
	BitsPerSample16 = 16
	// BitsPerSample24 位深 24
	BitsPerSample24 = 24
	// BitsPerSample32 位深 32
	BitsPerSample32 = 32
)

// WavHeader 定义了标准的 WAV 文件头 (44 bytes)
// 对应 RIFF WAVE 格式标准
type WavHeader struct {
	// RIFF Chunk
	ChunkID   [4]byte // "RIFF"
	ChunkSize uint32  // 文件总大小 - 8 字节
	Format    [4]byte // "WAVE"

	// fmt Chunk
	Subchunk1ID   [4]byte // "fmt "
	Subchunk1Size uint32  // 通常为 16 (针对 PCM)
	AudioFormat   uint16  // 音频格式: 1 = PCM (无损), 3 = IEEE Float
	NumChannels   uint16  // 声道数: 1 = Mono, 2 = Stereo
	SampleRate    uint32  // 采样率: e.g., 44100, 16000
	ByteRate      uint32  // 传输速率: SampleRate * NumChannels * BitsPerSample / 8
	BlockAlign    uint16  // 块对齐: NumChannels * BitsPerSample / 8
	BitsPerSample uint16  // 位深: e.g., 16, 24, 32

	// data Chunk
	Subchunk2ID   [4]byte // "data"
	Subchunk2Size uint32  // 音频数据的字节大小
}

// GetDuration 根据头部信息计算音频时长
func (h *WavHeader) GetDuration() time.Duration {
	if h.ByteRate == 0 {
		return 0
	}
	// 时长 = 数据总字节数 / 每秒字节传输率
	seconds := float64(h.Subchunk2Size) / float64(h.ByteRate)
	return time.Duration(seconds * float64(time.Second))
}

// String 格式化输出头部关键信息
func (h *WavHeader) String() string {
	return fmt.Sprintf(
		"WAV Header: [Format: %d, Channels: %d, Rate: %dHz, Bits: %d, DataSize: %d bytes, Duration: %s]",
		h.AudioFormat, h.NumChannels, h.SampleRate, h.BitsPerSample, h.Subchunk2Size, h.GetDuration(),
	)
}

// ParseWavHeader 从字节切片中解析 WAV 头部
//
// # Params:
//
//	data: 包含 WAV 头部信息的字节切片
func ParseWavHeader(data []byte) (*WavHeader, error) {
	if len(data) < 44 {
		return nil, fmt.Errorf("%w, data length must be at least 44 bytes", gotool.ErrInvalidParam)
	}

	// 使用 LittleEndian 解析
	reader := bytes.NewReader(data[:44])
	header := &WavHeader{}
	if err := binary.Read(reader, binary.LittleEndian, header); err != nil {
		return nil, err
	}

	// 校验标志位
	if string(header.ChunkID[:]) != "RIFF" || string(header.Format[:]) != "WAVE" {
		return nil, ErrNotWavFile
	}

	return header, nil
}

// ReadWavHeader 从文件中读取 WAV 头部
//
// # Params:
//
//	filePath: wav文件路径
func ReadWavHeader(filePath string) (*WavHeader, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 只读取前 44 个字节，避免加载整个大文件
	headerBytes := make([]byte, 44)
	if _, err := io.ReadFull(file, headerBytes); err != nil {
		return nil, err
	}

	return ParseWavHeader(headerBytes)
}

// WriteWav 将 PCM 数据封装为 WAV 格式写入 io.Writer
//
// # Params:
//
//	w: 写入目标
//	pcmData: 原始 PCM 数据
//	sampleRate: 采样率
//	channels: 声道数
//	bitsPerSample: 位深
func WriteWav(w io.Writer, pcmData []byte, sampleRate, channels, bitsPerSample int) error {
	if sampleRate <= 0 || channels <= 0 || bitsPerSample <= 0 {
		return fmt.Errorf("%w, rate=%d, chan=%d, bit=%d", gotool.ErrInvalidParam, sampleRate, channels, bitsPerSample)
	}

	dataSize := uint32(len(pcmData))

	// 计算相关的速率参数
	byteRate := uint32(sampleRate * channels * bitsPerSample / 8)
	blockAlign := uint16(channels * bitsPerSample / 8)

	// 构建 WAV 头部
	// ChunkSize = 36 + Subchunk2Size
	header := WavHeader{
		// RIFF Chunk
		ChunkID:   [4]byte{'R', 'I', 'F', 'F'},
		ChunkSize: 36 + dataSize,
		Format:    [4]byte{'W', 'A', 'V', 'E'},

		// fmt Chunk
		Subchunk1ID:   [4]byte{'f', 'm', 't', ' '},
		Subchunk1Size: 16, // PCM 格式固定为 16
		AudioFormat:   1,  // 1 表示 PCM (Linear Quantization)
		NumChannels:   uint16(channels),
		SampleRate:    uint32(sampleRate),
		ByteRate:      byteRate,
		BlockAlign:    blockAlign,
		BitsPerSample: uint16(bitsPerSample),

		// data Chunk
		Subchunk2ID:   [4]byte{'d', 'a', 't', 'a'},
		Subchunk2Size: dataSize,
	}

	// 写入头部
	if err := binary.Write(w, binary.LittleEndian, header); err != nil {
		return err
	}

	// 写入音频数据本体
	// 直接写入 []byte，不需要字节序转换 (本身就是字节流)
	if _, err := w.Write(pcmData); err != nil {
		return err
	}

	return nil
}

// SaveWav 将 PCM 数据保存为本地 WAV 文件
//
// # Params:
//
//	filePath: 文件路径
//	pcmData: 原始 PCM 数据
//	sampleRate: 采样率,例如: 16000(16KHz), 44100(44.1KHz)
//	channels: 声道数
//	bitsPerSample: 位深,例如: 16(CD音质), 24(专业录音), 32
func SaveWav(filePath string, pcmData []byte, sampleRate, channels, bitsPerSample int) error {
	// 创建文件
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	return WriteWav(f, pcmData, sampleRate, channels, bitsPerSample)
}

// Float32ToPcmBytes 将标准浮点音频数据转换为指定位深的 PCM 字节流
//
// # Params:
//
//	data:
//	 - 音频采样点数组 (Amplitudes)
//	 - 值域理论上应在 -1.0 到 +1.0 之间 (0.0 表示静音)
//	 - 超出范围的值会被削波 (Clipping) 处理
//	bitsPerSample: 位深,例如: 16(CD音质), 24(专业录音), 32
func Float32ToPcmBytes(data []float32, bitsPerSample int) ([]byte, error) {
	if len(data) == 0 {
		return []byte{}, nil
	}

	// 计算每个采样点的字节数
	bytesPerSample := bitsPerSample / 8
	output := make([]byte, len(data)*bytesPerSample)

	// 定义量化所需的缩放因子 (Scale Factor)
	// 16-bit: 32767
	// 24-bit: 8388607
	// 32-bit: 2147483647
	var scale float64
	switch bitsPerSample {
	case 16:
		scale = 32767.0
	case 24:
		scale = 8388607.0 // 2^23 - 1
	case 32:
		scale = 2147483647.0 // 2^31 - 1
	default:
		return nil, ErrUnsupportedBitDepth
	}

	offset := 0
	for _, sample := range data {
		// 处理特殊值和削波 (Clipping)
		if math.IsNaN(float64(sample)) {
			sample = 0
		}
		if sample > 1.0 {
			sample = 1.0
		} else if sample < -1.0 {
			sample = -1.0
		}

		// 量化并根据位深写入字节
		switch bitsPerSample {
		case 16:
			// 转换逻辑：int16
			val := int16(float64(sample) * scale)
			binary.LittleEndian.PutUint16(output[offset:], uint16(val))
			offset += 2

		case 24:
			// 转换逻辑：将值放大到 int32 范围，然后只取低 3 个字节
			// 24-bit Little Endian: [Low, Mid, High]
			val := int32(float64(sample) * scale)
			output[offset] = byte(val)         // 低位
			output[offset+1] = byte(val >> 8)  // 中位
			output[offset+2] = byte(val >> 16) // 高位
			offset += 3

		case 32:
			// 转换逻辑：int32
			val := int32(float64(sample) * scale)
			binary.LittleEndian.PutUint32(output[offset:], uint32(val))
			offset += 4
		}
	}

	return output, nil
}

// Float32ToWavBytes 将标准浮点音频数据转换为完整的 WAV 文件字节流
// 包含 WAV 头部 (Header) 和 PCM 数据体
//
// # Params:
//
//	data: 原始音频数据
//	 - 单声道 (Mono): [样本1, 样本2, 样本3, ...]
//	 - 双声道 (Stereo): [左1, 右1, 左2, 右2, 左3, 右3, ...]
//	sampleRate: 采样率
//	channels: 声道数
//	bitsPerSample: 位深
func Float32ToWavBytes(data []float32, sampleRate, channels, bitsPerSample int) ([]byte, error) {
	if sampleRate <= 0 || channels <= 0 || bitsPerSample <= 0 {
		return nil, fmt.Errorf("%w, rate=%d, chan=%d, bit=%d", gotool.ErrInvalidParam, sampleRate, channels, bitsPerSample)
	}

	// 将 Float32 数据转换为裸 PCM 字节流
	pcmBody, err := Float32ToPcmBytes(data, bitsPerSample)
	if err != nil {
		return nil, fmt.Errorf("pcm convert failed: %w", err)
	}

	// 写入缓冲区
	buf := new(bytes.Buffer)
	err = WriteWav(buf, pcmBody, sampleRate, channels, bitsPerSample)
	if err != nil {
		return nil, fmt.Errorf("write wav failed: %w", err)
	}

	return buf.Bytes(), nil
}

// PcmBytesToFloat32 PCM 字节流转 float32 数组
//
// # Params:
//
//	data: 原始 PCM 数据
//	bitPerSample: 位深,支持 16, 24, 32 bit
func PcmBytesToFloat32(data []byte, bitPerSample int) ([]float32, error) {
	if bitPerSample != 16 &&
		bitPerSample != 24 &&
		bitPerSample != 32 {
		return nil, fmt.Errorf("%w, unsupported source bit per sample", gotool.ErrInvalidParam)
	}

	bytesPerSample := bitPerSample / 8
	if len(data)%bytesPerSample != 0 {
		return nil, fmt.Errorf("%w, pcm data length is not aligned with bit per sample", gotool.ErrInvalidParam)
	}

	numSamples := len(data) / bytesPerSample
	output := make([]float32, numSamples)

	for i := 0; i < numSamples; i++ {
		offset := i * bytesPerSample
		var valFloat float64

		switch bitPerSample {
		case 16:
			valInt := int16(binary.LittleEndian.Uint16(data[offset:]))
			valFloat = float64(valInt) / 32767.0
		case 24:
			b0 := int32(data[offset])
			b1 := int32(data[offset+1])
			b2 := int32(data[offset+2])
			valInt := (b2 << 16) | (b1 << 8) | b0
			// 处理符号位
			if valInt&0x800000 != 0 {
				valInt |= ^0xFFFFFF // 扩展符号位到 32位
			}
			valFloat = float64(valInt) / 8388607.0
		case 32:
			valInt := int32(binary.LittleEndian.Uint32(data[offset:]))
			valFloat = float64(valInt) / 2147483647.0
		}
		output[i] = float32(valFloat)
	}
	return output, nil
}

// ReformatWavBytes WAV 字节流格式转换
//
// 支持：位深转换、采样率转换、声道转换
//
// # Params:
//
//	wavData: 原始 WAV 文件数据
//	targetRate: 目标采样率
//	targetChannels: 目标声道数
//	targetBitPerSample: 目标位深
func ReformatWavBytes(wavData []byte, targetRate, targetChannels, targetBitPerSample int) ([]byte, error) {
	// 解析原始头部
	header, err := ParseWavHeader(wavData)
	if err != nil {
		return nil, err
	}

	// 当前的参数状态
	currentRate := int(header.SampleRate)
	currentChannels := int(header.NumChannels)
	currentBitPerSample := int(header.BitsPerSample)

	// 不存在转换,直接返回
	if targetRate == currentRate &&
		targetChannels == currentChannels &&
		targetBitPerSample == currentBitPerSample {
		return wavData, nil
	}

	// 提取 PCM 数据并转为 float32
	const headerSize = 44
	pcmRaw := wavData[headerSize:]
	samples, err := PcmBytesToFloat32(pcmRaw, int(header.BitsPerSample))
	if err != nil {
		return nil, fmt.Errorf("decode pcm failed: %w", err)
	}

	// 声道转换
	// 优先处理声道，如果转为单声道，可以减少后续重采样 50% 的计算量
	if targetChannels > 0 && targetChannels != currentChannels {
		samples, err = changeChannels(samples, currentChannels, targetChannels)
		if err != nil {
			return nil, err
		}
		currentChannels = targetChannels
	}

	// 重采样
	if targetRate > 0 && targetRate != currentRate {
		samples = resampleSafe(samples, currentRate, targetRate, currentChannels)
		currentRate = targetRate
	}

	// 目标位深
	if targetBitPerSample <= 0 {
		targetBitPerSample = int(header.BitsPerSample)
	}

	// 编码回 WAV
	return Float32ToWavBytes(samples, currentRate, currentChannels, targetBitPerSample)
}

// changeChannels 音频数据声道数转换
//
// 支持: Stereo(2) -> Mono(1), Mono(1) -> Stereo(2)
//
// # Params:
//
//	data: 原始音频数据
//	srcChannel: 源声道数
//	dstChannel: 目标声道数
func changeChannels(data []float32, srcChannel, dstChannel int) ([]float32, error) {
	if srcChannel == dstChannel {
		return data, nil
	}

	// 双声道转单声道 (Stereo -> Mono)
	// 算法: (L + R) / 2
	if srcChannel == 2 && dstChannel == 1 {
		output := make([]float32, len(data)/2)
		for i := 0; i < len(output); i++ {
			// data[2*i] 是 L, data[2*i+1] 是 R
			l := data[2*i]
			r := data[2*i+1]
			output[i] = (l + r) / 2
		}
		return output, nil
	}

	// 单声道转双声道 (Mono -> Stereo)
	// 算法: L = M, R = M (复制)
	if srcChannel == 1 && dstChannel == 2 {
		output := make([]float32, len(data)*2)
		for i := 0; i < len(data); i++ {
			output[2*i] = data[i]   // L
			output[2*i+1] = data[i] // R
		}
		return output, nil
	}

	return nil, fmt.Errorf("%w, unsupported channel conversion: %d -> %d",
		gotool.ErrInvalidParam, srcChannel, dstChannel)
}

// resampleSafe 安全的重采样函数，支持单/双声道
//
// # Params:
//
//	data: 原始数据
//	oldRate: 源采样率
//	newRate: 目标采样率
//	channels: 声道数
func resampleSafe(data []float32, oldRate, newRate, channels int) []float32 {
	if oldRate == newRate {
		return data
	}

	// 单声道：直接线性插值
	if channels == 1 {
		return resampleLinear(data, oldRate, newRate)
	}

	// 双声道：拆分 -> 重采样 -> 合并
	if channels == 2 {
		// Split
		length := len(data) / 2
		left := make([]float32, length)
		right := make([]float32, length)
		for i := 0; i < length; i++ {
			left[i] = data[i*2]
			right[i] = data[i*2+1]
		}

		// Resample
		leftResampled := resampleLinear(left, oldRate, newRate)
		rightResampled := resampleLinear(right, oldRate, newRate)

		// Merge
		newLen := len(leftResampled)
		output := make([]float32, newLen*2)
		for i := 0; i < newLen; i++ {
			output[i*2] = leftResampled[i]
			output[i*2+1] = rightResampled[i]
		}
		return output
	}

	return data
}

// resampleLinear 线性插值重采样，适用于单轨道连续数据
//
// # Params:
//
//	data: 原始数据
//	oldRate: 源采样率
//	newRate: 目标采样率
func resampleLinear(data []float32, oldRate, newRate int) []float32 {
	if oldRate == newRate {
		return data
	}
	ratio := float64(oldRate) / float64(newRate)
	newLen := int(float64(len(data)) / ratio)
	output := make([]float32, newLen)

	for i := 0; i < newLen; i++ {
		srcIdxFloat := float64(i) * ratio
		srcIdxInt := int(srcIdxFloat)

		if srcIdxInt >= len(data)-1 {
			output[i] = data[len(data)-1]
			continue
		}

		fraction := srcIdxFloat - float64(srcIdxInt)
		sample1 := data[srcIdxInt]
		sample2 := data[srcIdxInt+1]

		output[i] = sample1 + float32(fraction)*(sample2-sample1)
	}
	return output
}
