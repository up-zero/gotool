package imageutil

import (
	"image"
	"image/color"
)

// morphologyOp 形态学基础操作，支持图片的腐蚀、膨胀
//
// # Params:
//
//	src: 源图片
//	se: 结构元素
func morphologyOp(src image.Image, se StructuringElement, initialVal uint8,
	compare func(newVal, bestVal uint8) bool) image.Image {
	bounds := src.Bounds()
	gray := Grayscale(src)
	dst := image.NewGray(bounds)

	// 核判空
	kernelHeight := len(se.Kernel)
	if kernelHeight == 0 {
		return src
	}
	kernelWidth := len(se.Kernel[0])
	if kernelWidth == 0 {
		return src
	}

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			var bestVal = initialVal

			for ky := 0; ky < kernelHeight; ky++ {
				for kx := 0; kx < kernelWidth; kx++ {
					if !se.Kernel[ky][kx] {
						continue
					}

					srcX := x + (kx - se.Anchor.X)
					srcY := y + (ky - se.Anchor.Y)
					safeX := max(bounds.Min.X, min(srcX, bounds.Max.X-1))
					safeY := max(bounds.Min.Y, min(srcY, bounds.Max.Y-1))

					val := gray.GrayAt(safeX, safeY).Y
					if compare(val, bestVal) {
						bestVal = val
					}
				}
			}

			dst.SetGray(x, y, color.Gray{Y: bestVal})
		}
	}

	return dst
}

// Erode 图片腐蚀
//
// # Params:
//
//	src: 源图片
//	se: 结构元素
//
// # Example:
//
//	Erode(img, NewRectKernel(3, 3))
func Erode(src image.Image, se StructuringElement) image.Image {
	return morphologyOp(src, se, 255, func(newVal, bestVal uint8) bool {
		return newVal < bestVal
	})
}

// ErodeFile 图片文件腐蚀
//
// # Params:
//
//	srcFile: 源图片文件
//	dstFile: 目标图片文件
//	se: 结构元素
func ErodeFile(srcFile, dstFile string, se StructuringElement) error {
	img, err := Open(srcFile)
	if err != nil {
		return err
	}
	return Save(dstFile, Erode(img, se), 100)
}

// Dilate 图片膨胀
//
// # Params:
//
//	src: 源图片
//	se: 结构元素
//
// # Example:
//
//	Dilate(img, NewRectKernel(3, 3))
func Dilate(src image.Image, se StructuringElement) image.Image {
	return morphologyOp(src, se, 0, func(newVal, bestVal uint8) bool {
		return newVal > bestVal
	})
}

// DilateFile 图片文件膨胀
//
// # Params:
//
//	srcFile: 源图片文件
//	dstFile: 目标图片文件
//	se: 结构元素
func DilateFile(srcFile, dstFile string, se StructuringElement) error {
	img, err := Open(srcFile)
	if err != nil {
		return err
	}
	return Save(dstFile, Dilate(img, se), 100)
}

// MorphologyOpen 开运算，先腐蚀，再膨胀
//
// # Params:
//
//	src: 源图片
//	se: 结构元素
//
// # Example:
//
//	MorphologyOpen(img, NewRectKernel(3, 3))
func MorphologyOpen(src image.Image, se StructuringElement) image.Image {
	erodedImage := Erode(src, se)
	openedImage := Dilate(erodedImage, se)

	return openedImage
}

// MorphologyOpenFile 对图片文件进行开运算
//
// # Params:
//
//	srcFile: 源图片文件
//	dstFile: 目标图片文件
//	se: 结构元素
func MorphologyOpenFile(srcFile, dstFile string, se StructuringElement) error {
	img, err := Open(srcFile)
	if err != nil {
		return err
	}
	return Save(dstFile, MorphologyOpen(img, se), 100)
}

// MorphologyClose 闭运算，先膨胀，再腐蚀
//
// # Params:
//
//	src: 源图片
//	se: 结构元素
//
// # Example:
//
//	MorphologyClose(img, NewRectKernel(3, 3))
func MorphologyClose(src image.Image, se StructuringElement) image.Image {
	dilatedImage := Dilate(src, se)
	closedImage := Erode(dilatedImage, se)

	return closedImage
}

// MorphologyCloseFile 对图片文件进行闭运算
//
// # Params:
//
//	srcFile: 源图片文件
//	dstFile: 目标图片文件
//	se: 结构元素
func MorphologyCloseFile(srcFile, dstFile string, se StructuringElement) error {
	img, err := Open(srcFile)
	if err != nil {
		return err
	}
	return Save(dstFile, MorphologyClose(img, se), 100)
}
