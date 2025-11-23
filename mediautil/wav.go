package mediautil

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

// 常见错误定义
var (
	ErrInvalidWavData = errors.New("invalid wav data: too short or incorrect format")
	ErrNotWavFile     = errors.New("not a valid RIFF/WAVE file")
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
		return nil, ErrInvalidWavData
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
