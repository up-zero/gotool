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
	"sort"
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
//	src: 源图片
//	cropRect: 裁剪区域
func Crop(src image.Image, cropRect image.Rectangle) (image.Image, error) {
	subImager, ok := src.(interface {
		SubImage(r image.Rectangle) image.Image
	})
	if !ok {
		return nil, fmt.Errorf("unsupported image format")
	}
	dstImage := subImager.SubImage(cropRect)
	return dstImage, nil
}

// CropFile 图片文件裁剪
//
// # Params:
//
//	srcFile: 源图片路径
//	dstFile: 目标图片路径
//	cropRect: 裁剪区域
func CropFile(srcFile, dstFile string, cropRect image.Rectangle) error {
	// 打开文件
	img, err := Open(srcFile)
	if err != nil {
		return err
	}
	dstImage, err := Crop(img, cropRect)
	if err != nil {
		return err
	}
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

// Overlay 图片叠加，将 overlay 图片叠加到 base 图片的指定位置 (x, y)，支持透明通道 (alpha blending)
//
// # Params:
//
//	base: 基础图片
//	overlay: 叠加图片
//	x: 基础图片的 X 坐标
//	y: 基础图片的 Y 坐标
//
// # Example:
//
//	Overlay(base, overlay, 100, 100) // 将 overlay 图片叠加到 base 图片的 (100, 100) 位置
func Overlay(base, overlay image.Image, x, y int) image.Image {
	// 基础图片
	dst := image.NewRGBA(base.Bounds())
	draw.Draw(dst, dst.Bounds(), base, image.Point{X: 0, Y: 0}, draw.Src)

	// 叠加图片
	overlayRect := image.Rect(x, y, x+overlay.Bounds().Dx(), y+overlay.Bounds().Dy())
	draw.Draw(dst, overlayRect, overlay, image.Point{X: 0, Y: 0}, draw.Over)

	return dst
}

// OverlayFile 图片文件叠加，将 overlay 图片文件叠加到 base 图片文件指定位置 (x, y)，支持透明通道 (alpha blending)
//
// # Params:
//
//	baseFile: 基础图片文件
//	overlayFile: 叠加图片文件
//	dstFile: 目标图片文件
//	x: 基础图片的 X 坐标
//	y: 基础图片的 Y 坐标
func OverlayFile(baseFile, overlayFile, dstFile string, x, y int) error {
	baseImg, err := Open(baseFile)
	if err != nil {
		return err
	}
	overlayImg, err := Open(overlayFile)
	if err != nil {
		return err
	}
	dstImg := Overlay(baseImg, overlayImg, x, y)
	return Save(dstFile, dstImg, 100)
}

// Grayscale 图片灰度化
//
// # Params:
//
//	src: 源图片
func Grayscale(src image.Image) *image.Gray {
	bounds := src.Bounds()
	dst := image.NewGray(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := src.At(x, y)
			// 使用默认的灰阶模型转换
			// 亮度公式：Y = 0.299R + 0.587G + 0.114B
			grayY := color.GrayModel.Convert(c).(color.Gray).Y
			dst.SetGray(x, y, color.Gray{Y: grayY})
		}
	}
	return dst
}

// GrayscaleFile 图片文件灰度化
//
// # Params:
//
//	srcFile: 源图片文件
//	dstFile: 目标图片文件
func GrayscaleFile(srcFile, dstFile string) error {
	img, err := Open(srcFile)
	if err != nil {
		return err
	}
	return Save(dstFile, Grayscale(img), 100)
}

// gaussianKernel 生成高斯核
func gaussianKernel(radius int, sigma float64) [][]float64 {
	size := 2*radius + 1
	kernel := make([][]float64, size)
	sum := 0.0

	for i := 0; i < size; i++ {
		kernel[i] = make([]float64, size)
		for j := 0; j < size; j++ {
			x := float64(i - radius)
			y := float64(j - radius)
			// 高斯公式：G(x,y) = (1/(2πσ²)) * exp(-(x²+y²)/(2σ²))
			value := math.Exp(-(x*x+y*y)/(2*sigma*sigma)) / (2 * math.Pi * sigma * sigma)
			kernel[i][j] = value
			sum += value
		}
	}

	// 归一化
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			kernel[i][j] /= sum
		}
	}
	return kernel
}

// GaussianBlur 图片高斯模糊
//
//   - 高斯核半径和标准差越大，越模糊
//   - 建议 radius > 3*sigma
//   - 小值（radius=1-3, sigma=0.5-1.0）：轻微模糊，适合细微效果
//   - 中等值（radius=3-5, sigma=1.0-2.0）：明显模糊，适合背景虚化
//   - 大值（radius>5, sigma>2.0）：强烈模糊，适合艺术效果，但计算慢
//
// # Params:
//
//	src: 源图片
//	radius: 高斯核半径
//	sigma: 高斯核标准差
//
// # Example:
//
//	GaussianBlur(img, 3, 1.0)
func GaussianBlur(src image.Image, radius int, sigma float64) image.Image {
	bounds := src.Bounds()
	width, height := bounds.Dx(), bounds.Dy()
	dst := image.NewRGBA(bounds)

	kernel := gaussianKernel(radius, sigma)
	kernelSize := 2*radius + 1

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			var r, g, b, a float64
			for ky := 0; ky < kernelSize; ky++ {
				for kx := 0; kx < kernelSize; kx++ {
					// 计算源图像坐标，边界处理：镜像反射
					px := x + kx - radius
					py := y + ky - radius
					if px < 0 {
						px = -px
					}
					if py < 0 {
						py = -py
					}
					if px >= width {
						px = 2*width - px - 1
					}
					if py >= height {
						py = 2*height - py - 1
					}

					// 获取像素颜色
					cr, cg, cb, ca := src.At(px, py).RGBA()
					weight := kernel[ky][kx]
					r += float64(cr>>8) * weight
					g += float64(cg>>8) * weight
					b += float64(cb>>8) * weight
					a += float64(ca>>8) * weight
				}
			}
			dst.Set(x, y, color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)})
		}
	}
	return dst
}

// GaussianBlurFile 图片文件高斯模糊
//
// # Params:
//
//	srcFile: 源图片文件
//	dstFile: 目标图片文件
//	radius: 高斯核半径
//	sigma: 高斯核标准差
//
// # Example:
//
//	GaussianBlurFile("test.png", "test_gaussian.png", 3, 1.0)
func GaussianBlurFile(srcFile, dstFile string, radius int, sigma float64) error {
	img, err := Open(srcFile)
	if err != nil {
		return err
	}
	return Save(dstFile, GaussianBlur(img, radius, sigma), 100)
}

// AdjustBrightness 图片亮度调整
//
// # Params:
//
//	src: 源图片
//	brightness: 亮度调整值
//
// # Example:
//
//	AdjustBrightness(img, 50)
func AdjustBrightness(src image.Image, brightness float64) image.Image {
	bounds := src.Bounds()
	dst := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := src.At(x, y).RGBA()
			// 转换为 0-255 范围
			r8, g8, b8 := float64(r>>8), float64(g>>8), float64(b>>8)
			// 调整亮度
			r8 = r8 + brightness
			g8 = g8 + brightness
			b8 = b8 + brightness
			// 限制在 [0, 255]
			r8 = max(0, min(255, r8))
			g8 = max(0, min(255, g8))
			b8 = max(0, min(255, b8))
			dst.Set(x, y, color.RGBA{R: uint8(r8), G: uint8(g8), B: uint8(b8), A: uint8(a >> 8)})
		}
	}
	return dst
}

// AdjustBrightnessFile 图片文件亮度调整
//
// # Params:
//
//	srcFile: 源图片文件
//	dstFile: 目标图片文件
//	brightness: 亮度调整值
func AdjustBrightnessFile(srcFile, dstFile string, brightness float64) error {
	img, err := Open(srcFile)
	if err != nil {
		return err
	}
	return Save(dstFile, AdjustBrightness(img, brightness), 100)
}

// Invert 图片反转颜色
//
// # Params:
//
//	src: 源图片
func Invert(src image.Image) image.Image {
	bounds := src.Bounds()
	dst := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := src.At(x, y)
			r, g, b, a := c.RGBA()
			// 反色：255 - 原值
			dst.Set(x, y, color.RGBA{
				R: uint8(255 - (r >> 8)),
				G: uint8(255 - (g >> 8)),
				B: uint8(255 - (b >> 8)),
				A: uint8(a >> 8),
			})
		}
	}
	return dst
}

// InvertFile 图片文件反转颜色
//
// # Params:
//
//	srcFile: 源图片文件
//	dstFile: 目标图片文件
func InvertFile(srcFile, dstFile string) error {
	img, err := Open(srcFile)
	if err != nil {
		return err
	}
	return Save(dstFile, Invert(img), 100)
}

// Binarize 图片二值化
//
// # Params:
//
//	src: 源图片
//	threshold: 阈值，推荐为 128
func Binarize(src image.Image, threshold uint8) image.Image {
	bounds := src.Bounds()
	dst := image.NewGray(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			oldColor := src.At(x, y)
			// 将原始图片转为 8-bit 灰阶值
			var grayY uint8
			switch c := oldColor.(type) {
			case color.Gray:
				grayY = c.Y
			case color.Gray16:
				grayY = uint8(c.Y >> 8)
			default:
				// 使用默认的灰阶模型转换
				// 亮度公式：Y = 0.299R + 0.587G + 0.114B
				grayY = color.GrayModel.Convert(oldColor).(color.Gray).Y
			}
			// 根据阈值进行二值化
			if grayY >= threshold {
				dst.Set(x, y, color.Gray{Y: 255}) // 白色
			} else {
				dst.Set(x, y, color.Gray{Y: 0}) // 黑色
			}
		}
	}
	return dst
}

// BinarizeFile 图片文件二值化
//
// # Params:
//
//	srcFile: 源图片文件
//	dstFile: 目标图片文件
//	threshold: 阈值，推荐为 128
func BinarizeFile(srcFile, dstFile string, threshold uint8) error {
	img, err := Open(srcFile)
	if err != nil {
		return err
	}
	return Save(dstFile, Binarize(img, threshold), 100)
}

// MedianBlur 图片中值滤波
//
// # Params
//
//	src: 源图像
//	radius: 滤波半径，radius=1 表示 3x3 窗口, radius=2 表示 5x5 窗口
func MedianBlur(src image.Image, radius int) image.Image {
	bounds := src.Bounds()
	dst := image.NewRGBA(bounds)

	// 窗口的边长 (e.g., radius=1 -> size=3)
	size := radius*2 + 1
	// 窗口内的像素总数 (e.g., radius=1 -> 3x3=9)
	windowSize := size * size
	// 中位数在排序后切片中的索引
	medianIndex := windowSize / 2

	// 为 R, G, B, A 四个通道分别创建切片，用于存放窗口内的像素值
	rNeighbors := make([]uint8, windowSize)
	gNeighbors := make([]uint8, windowSize)
	bNeighbors := make([]uint8, windowSize)
	aNeighbors := make([]uint8, windowSize)

	// 遍历图像的每一个像素 (x, y)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {

			// 遍历 (x, y) 像素周围的窗口
			i := 0
			for ky := -radius; ky <= radius; ky++ {
				for kx := -radius; kx <= radius; kx++ {
					// 计算邻居像素的坐标
					nx := x + kx
					ny := y + ky

					// 边界处理：如果邻居在图像外，则复制最近的边缘像素 (Clamping)
					if nx < bounds.Min.X {
						nx = bounds.Min.X
					} else if nx >= bounds.Max.X {
						nx = bounds.Max.X - 1
					}
					if ny < bounds.Min.Y {
						ny = bounds.Min.Y
					} else if ny >= bounds.Max.Y {
						ny = bounds.Max.Y - 1
					}

					// 获取邻居像素的颜色
					r, g, b, a := src.At(nx, ny).RGBA()

					// 存入切片
					rNeighbors[i] = uint8(r >> 8)
					gNeighbors[i] = uint8(g >> 8)
					bNeighbors[i] = uint8(b >> 8)
					aNeighbors[i] = uint8(a >> 8)
					i++
				}
			}

			// 排序四个通道的切片
			sort.Slice(rNeighbors, func(i, j int) bool { return rNeighbors[i] < rNeighbors[j] })
			sort.Slice(gNeighbors, func(i, j int) bool { return gNeighbors[i] < gNeighbors[j] })
			sort.Slice(bNeighbors, func(i, j int) bool { return bNeighbors[i] < bNeighbors[j] })
			sort.Slice(aNeighbors, func(i, j int) bool { return aNeighbors[i] < aNeighbors[j] })

			// 找出中位数
			medianR := rNeighbors[medianIndex]
			medianG := gNeighbors[medianIndex]
			medianB := bNeighbors[medianIndex]
			medianA := aNeighbors[medianIndex]

			// 将中位数作为新颜色设置到目标图像
			dst.Set(x, y, color.RGBA{
				R: medianR,
				G: medianG,
				B: medianB,
				A: medianA,
			})
		}
	}

	return dst
}

// MedianBlurFile 图片文件中值滤波
//
// # Params:
//
//	srcFile: 源图片文件
//	dstFile: 目标图片文件
//	radius: 滤波半径，radius=1 表示 3x3 窗口, radius=2 表示 5x5 窗口
func MedianBlurFile(srcFile, dstFile string, radius int) error {
	src, err := Open(srcFile)
	if err != nil {
		return err
	}
	return Save(dstFile, MedianBlur(src, radius), 100)
}

// Sobel 索贝尔边缘检测
//
// # Sobel 算子说明：
//   - 使用 3x3 Gx/Gy 卷积核分别计算水平与垂直方向梯度
//   - 边缘强度 edge = sqrt(gx^2 + gy^2)
//
// # Params:
//
//	src: 源图片
//	threshold: 阈值 [0, 1442] 推荐值：400
func Sobel(src image.Image, threshold float64) *image.Gray {
	// 图片灰阶
	gray := Grayscale(src)
	bounds := src.Bounds()

	// 定义 Sobel 核
	kernelX := [3][3]int{
		{-1, 0, 1},
		{-2, 0, 2},
		{-1, 0, 1},
	}
	kernelY := [3][3]int{
		{-1, -2, -1},
		{0, 0, 0},
		{1, 2, 1},
	}

	dst := image.NewGray(bounds)

	for y := bounds.Min.Y + 1; y < bounds.Max.Y-1; y++ {
		for x := bounds.Min.X + 1; x < bounds.Max.X-1; x++ {
			var gx, gy int // 水平和垂直梯度

			// 应用 3x3 卷积核
			for ky := -1; ky <= 1; ky++ {
				for kx := -1; kx <= 1; kx++ {
					// 获取 3x3 邻域内的像素灰阶值
					// .GrayAt(x, y) 返回 color.Gray, .Y 即为 8 位灰阶值
					pixelVal := int(gray.GrayAt(x+kx, y+ky).Y)

					// 累加 Gx 和 Gy
					gx += pixelVal * kernelX[ky+1][kx+1]
					gy += pixelVal * kernelY[ky+1][kx+1]
				}
			}

			// 计算梯度幅值
			// G = Sqrt(Gx^2 + Gy^2)
			gradient := math.Sqrt(float64(gx*gx) + float64(gy*gy))

			if gradient > threshold {
				dst.SetGray(x, y, color.Gray{Y: 255}) // 白色
			} else {
				dst.SetGray(x, y, color.Gray{Y: 0}) // 黑色
			}

		}
	}

	return dst
}

// SobelFile 图片文件索贝尔边缘检测
//
// # Params:
//
//	srcFile: 源图片文件
//	dstFile: 目标图片文件
//	threshold: 阈值 [0, 1442] 推荐值：400
func SobelFile(srcFile, dstFile string, threshold float64) error {
	img, err := Open(srcFile)
	if err != nil {
		return err
	}
	return Save(dstFile, Sobel(img, threshold), 100)
}
