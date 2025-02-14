package gotool

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"time"
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

// GenerateCaptcha 验证码图片生成（目前只支持数字）
//
// text 验证码文本
//
// # Example 1:
//
//	// image.Image to bytes
//	img, _ := GenerateCaptcha("5679")
//	var buf bytes.Buffer
//	err := png.Encode(&buf, img)
//	if err != nil {
//		t.Fatal("Error encoding image:", err)
//	}
//	buf.Bytes()
func GenerateCaptcha(text string) (image.Image, error) {
	var (
		charAreaWidth = 25
		width         = len(text) * charAreaWidth
		height        = 40
		charWidth     = 10
		charHeight    = 20
		marginWidth   = 5
		marginHeight  = 5
	)
	for _, c := range text {
		if _, ok := simpleFontMap[c]; !ok {
			return nil, ErrNotSupportFormat
		}
	}

	// 创建一个白底的 RGBA 图像
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	bgColor := color.RGBA{R: 255, G: 255, B: 255, A: 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{C: bgColor}, image.Point{}, draw.Src)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	textColor := color.RGBA{A: 255}

	for i, char := range text {
		x := i*charAreaWidth + marginWidth + r.Intn(charAreaWidth-marginWidth-charWidth)
		y := marginHeight + r.Intn(height-marginHeight-charHeight)
		drawChar(img, x, y, char, textColor)
	}

	// 添加一些干扰线
	for i := 0; i < 5; i++ {
		x1 := r.Intn(width)
		y1 := r.Intn(height)
		x2 := r.Intn(width)
		y2 := r.Intn(height)
		drawLine(img, x1, y1, x2, y2, color.RGBA{R: uint8(r.Intn(255)), G: uint8(r.Intn(255)), B: uint8(r.Intn(255)), A: 255})
	}

	return img, nil
}

func drawChar(img *image.RGBA, x, y int, char rune, color color.RGBA) {
	if pixels, ok := simpleFontMap[char]; ok {
		for py, row := range pixels {
			for px, value := range row {
				if value == 1 {
					img.Set(x+px, y+py, color)
				}
			}
		}
	}
}

// drawLine 在图像上绘制一条线
func drawLine(img *image.RGBA, x1, y1, x2, y2 int, color color.RGBA) {
	dx := int(math.Abs(float64(x2 - x1)))
	dy := int(math.Abs(float64(y2 - y1)))
	sx := 1
	sy := 1

	if x1 > x2 {
		sx = -1
	}
	if y1 > y2 {
		sy = -1
	}

	err := dx - dy

	for {
		img.Set(x1, y1, color)
		if x1 == x2 && y1 == y2 {
			break
		}
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x1 += sx
		}
		if e2 < dx {
			err += dx
			y1 += sy
		}
	}
}
