package gotool

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

// ImageCompression 图片压缩
//
// srcFile 源图片路径
// dstFile 目标图片路径
// quality 压缩质量，范围 1-100（值越低，压缩率越高，质量越低），对于 PNG 图片，映射到 0-9 的压缩级别（0：无压缩，9：最大压缩）
func ImageCompression(srcFile, dstFile string, quality int) error {
	file, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}
	if err = os.MkdirAll(filepath.Dir(dstFile), os.ModePerm); err != nil {
		return err
	}

	outFile, err := os.Create(dstFile)
	if err != nil {
		return err
	}
	defer outFile.Close()

	compressionLevel := int((100 - quality) / 10) // 从质量 100 映射到 0（无压缩），质量 0 映射到 9（最大压缩）

	imgType := filepath.Ext(dstFile)
	switch imgType {
	case ".jpeg", ".jpg":
		err = jpeg.Encode(outFile, img, &jpeg.Options{Quality: quality})
		if err != nil {
			return err
		}
	case ".png":
		encoder := png.Encoder{CompressionLevel: png.CompressionLevel(compressionLevel)}
		err = encoder.Encode(outFile, img)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported image format: %s", imgType)
	}

	return nil
}

// ImageSize 图片尺寸
// 说明：当图片类型不是标准库提供的，需要导入扩展库中的image golang.org/x/image
//
// imagePath 图片路径
func ImageSize(imagePath string) (*ImageSizeReply, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	imgConfig, _, err := image.DecodeConfig(file)
	if err != nil {
		return nil, err
	}
	return &ImageSizeReply{
		Height: imgConfig.Height,
		Width:  imgConfig.Width,
	}, nil
}
