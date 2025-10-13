package imageutil

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

	"github.com/up-zero/gotool"
)

// Open 打开图片
//
// # Params:
//
//	imagePath: 图片路径
func Open(imagePath string) (image.Image, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return nil, fmt.Errorf("open file error: %v", err)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("decode image error: %v", err)
	}
	return img, nil
}

// Save 保存图片
//
// # Params:
//
//	imagePath: 图片路径
//	img: 图片
//	quality: 压缩质量，范围 1-100（值越低，压缩率越高，质量越低）
func Save(imagePath string, img image.Image, quality int) error {
	if quality < 1 {
		quality = 1
	}
	if quality > 100 {
		quality = 100
	}

	// 创建文件
	if err := os.MkdirAll(filepath.Dir(imagePath), os.ModePerm); err != nil {
		return err
	}
	imageFile, err := os.Create(imagePath)
	if err != nil {
		return err
	}
	defer imageFile.Close()

	// 保存文件
	imgType := filepath.Ext(imagePath)
	switch imgType {
	case ".jpeg", ".jpg":
		return jpeg.Encode(imageFile, img, &jpeg.Options{Quality: quality})
	case ".png":
		// 将 1-100 的 quality 映射到 png.CompressionLevel
		// quality 100 -> DefaultCompression (0)
		// quality 1   -> BestCompression (-3)
		var level png.CompressionLevel
		switch {
		case quality == 100:
			level = png.DefaultCompression
		case quality >= 90:
			level = png.BestSpeed
		case quality >= 50:
			level = png.DefaultCompression
		default:
			level = png.BestCompression
		}
		encoder := png.Encoder{CompressionLevel: level}
		return encoder.Encode(imageFile, img)
	default:
		return fmt.Errorf("unsupported image format: %s", imgType)
	}
}

// Compression 图片压缩
//
// # Params:
//
//	srcFile: 源图片路径
//	dstFile: 目标图片路径
//	quality: 压缩质量，范围 1-100（值越低，压缩率越高，质量越低）
func Compression(srcFile, dstFile string, quality int) error {
	img, err := Open(srcFile)
	if err != nil {
		return err
	}
	return Save(dstFile, img, quality)
}

// Size 图片尺寸
// 说明：当图片类型不是标准库提供的，需要导入扩展库中的image golang.org/x/image
//
// # Params:
//
//	imagePath: 图片路径
func Size(imagePath string) (*SizeReply, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	imgConfig, _, err := image.DecodeConfig(file)
	if err != nil {
		return nil, err
	}
	return &SizeReply{
		Height: imgConfig.Height,
		Width:  imgConfig.Width,
	}, nil
}

// GenerateCaptcha 验证码图片生成（目前只支持数字）
//
// # Params:
//
// text: 验证码文本
//
// # Example:
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
		if _, ok := gotool.SimpleFontMap[c]; !ok {
			return nil, gotool.ErrNotSupportFormat
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
	if pixels, ok := gotool.SimpleFontMap[char]; ok {
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

// Crop 图片裁剪
//
// # Params:
//
//	srcFile: 源图片路径
//	dstFile: 目标图片路径
//	cropRect: 裁剪区域
func Crop(srcFile, dstFile string, cropRect image.Rectangle) error {
	// 打开文件
	img, err := Open(srcFile)
	outFile, err := os.Create(dstFile)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// 裁剪图片
	subImager, ok := img.(interface {
		SubImage(r image.Rectangle) image.Image
	})
	if !ok {
		return fmt.Errorf("unsupported image format: %s", filepath.Ext(srcFile))
	}
	dstImage := subImager.SubImage(cropRect)

	// 保存图片
	return Save(dstFile, dstImage, 100)
}

// getRGBA 将 color.Color 转换为 uint8 的 RGBA 值
func getRGBA(c color.Color) (uint8, uint8, uint8, uint8) {
	r, g, b, a := c.RGBA()
	return uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)
}

// bilinearInterpolate 执行双线性插值
func bilinearInterpolate(c00, c10, c01, c11 color.Color, dx, dy float64) color.Color {
	r00, g00, b00, a00 := getRGBA(c00)
	r10, g10, b10, a10 := getRGBA(c10)
	r01, g01, b01, a01 := getRGBA(c01)
	r11, g11, b11, a11 := getRGBA(c11)

	// 水平插值
	r0 := (1-dx)*float64(r00) + dx*float64(r10)
	g0 := (1-dx)*float64(g00) + dx*float64(g10)
	b0 := (1-dx)*float64(b00) + dx*float64(b10)
	a0 := (1-dx)*float64(a00) + dx*float64(a10)

	r1 := (1-dx)*float64(r01) + dx*float64(r11)
	g1 := (1-dx)*float64(g01) + dx*float64(g11)
	b1 := (1-dx)*float64(b01) + dx*float64(b11)
	a1 := (1-dx)*float64(a01) + dx*float64(a11)

	// 垂直插值
	r := uint8((1-dy)*r0 + dy*r1)
	g := uint8((1-dy)*g0 + dy*g1)
	b := uint8((1-dy)*b0 + dy*b1)
	a := uint8((1-dy)*a0 + dy*a1)

	return color.RGBA{R: r, G: g, B: b, A: a}
}

// Resize 图片缩放
//
//   - 如果 newWidth > 0 && newHeight == 0：按比例基于宽度缩放
//   - 如果 newWidth == 0 && newHeight > 0：按比例基于高度缩放
//   - 如果 newWidth > 0 && newHeight > 0：固定宽高缩放（可能扭曲）
//   - 如果两者均为 0：返回原图
//
// # Params:
//
//	src: 源图片
//	newWidth: 新宽度
//	newHeight: 新高度
func Resize(src image.Image, newWidth, newHeight int) image.Image {
	bounds := src.Bounds()
	srcWidth := bounds.Dx()
	srcHeight := bounds.Dy()

	// 返回原图
	if newWidth == 0 && newHeight == 0 {
		return src
	}
	// 基于高度按比例计算宽度
	if newWidth == 0 {
		newWidth = srcWidth * newHeight / srcHeight
	}
	// 基于宽度按比例计算高度
	if newHeight == 0 {
		newHeight = srcHeight * newWidth / srcWidth
	}

	// 创建目标图像
	dst := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

	// 循环每个像素，进行双线性插值
	for y := 0; y < newHeight; y++ {
		for x := 0; x < newWidth; x++ {
			// 计算源图像对应位置（浮点）
			fx := float64(x) * float64(srcWidth) / float64(newWidth)
			fy := float64(y) * float64(srcHeight) / float64(newHeight)

			// 源图像整数坐标
			ix := int(math.Floor(fx))
			iy := int(math.Floor(fy))

			// 小数部分
			dx := fx - float64(ix)
			dy := fy - float64(iy)

			// 边界处理：使用最近像素
			ix = max(0, min(ix, srcWidth-1))
			iy = max(0, min(iy, srcHeight-1))
			ix1 := max(0, min(ix+1, srcWidth-1))
			iy1 := max(0, min(iy+1, srcHeight-1))

			// 获取四个邻近像素
			c00 := src.At(ix, iy)
			c10 := src.At(ix1, iy)
			c01 := src.At(ix, iy1)
			c11 := src.At(ix1, iy1)

			// 插值并设置像素
			dst.Set(x, y, bilinearInterpolate(c00, c10, c01, c11, dx, dy))
		}
	}

	return dst
}

// ResizeFile 图片文件缩放
//
// # Params:
//
//	srcFile: 源图片文件
//	dstFile: 目标图片文件
//	newWidth: 新宽度
//	newHeight: 新高度
func ResizeFile(srcFile, dstFile string, newWidth, newHeight int) error {
	img, err := Open(srcFile)
	if err != nil {
		return err
	}
	dstImg := Resize(img, newWidth, newHeight)
	return Save(dstFile, dstImg, 100)
}

const (
	// RotateAngle90 旋转角度90°
	RotateAngle90 = 90
	// RotateAngle180 旋转角度180°
	RotateAngle180 = 180
	// RotateAngle270 转角度270°
	RotateAngle270 = 270
)

// Rotate 旋转图片（顺时针 90°、180°、270°）
//
// # Params:
//
//	src: 源图片
//	angle: 旋转角度
//
// # Example:
//
//	Rotate(img, RotateAngle90) // 旋转90°
func Rotate(src image.Image, angle int) (image.Image, error) {
	bounds := src.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	var dst *image.RGBA

	switch angle {
	case RotateAngle90:
		dst = image.NewRGBA(image.Rect(0, 0, height, width))
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				dst.Set(height-1-y, x, src.At(x, y))
			}
		}
	case RotateAngle180:
		dst = image.NewRGBA(image.Rect(0, 0, width, height))
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				dst.Set(width-1-x, height-1-y, src.At(x, y))
			}
		}
	case RotateAngle270:
		dst = image.NewRGBA(image.Rect(0, 0, height, width))
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				dst.Set(y, width-1-x, src.At(x, y))
			}
		}
	default:
		return nil, fmt.Errorf("unsupported rotation angle: %d", angle)
	}

	return dst, nil
}

// RotateFile 旋转图片文件（顺时针 90°、180°、270°）
//
// # Params:
//
//	srcFile: 源图片文件
//	dstFile: 目标图片文件
//	angle: 旋转角度
func RotateFile(srcFile, dstFile string, angle int) error {
	img, err := Open(srcFile)
	if err != nil {
		return err
	}
	dstImg, err := Rotate(img, angle)
	if err != nil {
		return err
	}
	return Save(dstFile, dstImg, 100)
}

const (
	// FlipModeHorizontal 水平翻转（左右镜像）
	FlipModeHorizontal = "horizontal"
	// FlipModeVertical 垂直翻转（上下镜像）
	FlipModeVertical = "vertical"
)

// Flip 翻转图片（水平或垂直）
//
// # Params:
//
//	src: 源图片
//	mode: 翻转模式
//
// # Example:
//
//	Flip(img, FlipModeHorizontal) // 水平翻转
func Flip(src image.Image, mode string) (image.Image, error) {
	bounds := src.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	switch mode {
	case FlipModeHorizontal: // 水平翻转（左右镜像）
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				dst.Set(width-1-x, y, src.At(x, y))
			}
		}
	case FlipModeVertical: // 垂直翻转（上下镜像）
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				dst.Set(x, height-1-y, src.At(x, y))
			}
		}
	default:
		return nil, fmt.Errorf("unsupported flip mode: %s", mode)
	}

	return dst, nil
}

// FlipFile 翻转图片文件（水平或垂直）
//
// # Params:
//
//	srcFile: 源图片文件
//	dstFile: 目标图片文件
//	mode: 翻转模式
func FlipFile(srcFile, dstFile string, mode string) error {
	img, err := Open(srcFile)
	if err != nil {
		return err
	}
	dstImg, err := Flip(img, mode)
	if err != nil {
		return err
	}
	return Save(dstFile, dstImg, 100)
}
