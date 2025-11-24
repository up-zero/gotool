package mediautil

import (
	"math/rand"
	"testing"
)

func TestReadWavHeader(t *testing.T) {
	header, err := ReadWavHeader("test.wav")
	if err != nil {
		t.Fatalf("read wav header error: %v", err)
		return
	}
	t.Logf("wav header: %+v", header)
}

func TestSaveWav(t *testing.T) {
	// 2s秒的音频数据
	sampleRate := 16000
	seconds := 2

	// 16bit = 2 bytes per sample
	dataSize := sampleRate * seconds * 2
	pcmData := make([]byte, dataSize)

	// 填充随机数据模拟声音
	for i := 0; i < len(pcmData); i++ {
		pcmData[i] = byte(rand.Intn(255))
	}

	// 保存为 WAV 文件
	fileName := "output_save.wav"
	err := SaveWav(fileName, pcmData, sampleRate, 1, 16)
	if err != nil {
		t.Fatalf("保存失败: %v", err)
	}
	t.Logf("保存成功: %s", fileName)
}
