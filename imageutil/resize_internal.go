package imageutil

import (
	"image"
	"image/color"
	"runtime"
	"sync"
)

const resizeParallelThresholdPixels = 512 * 512
const resizeWeightScale = 1 << 16
const resizeBilinearRound = 1 << 31

type resizeAxisSample struct {
	i0 int
	i1 int
	w  uint16
}

func normalizeResizeTargetSize(srcWidth, srcHeight, newWidth, newHeight int) (int, int) {
	if srcWidth <= 0 || srcHeight <= 0 {
		return 0, 0
	}
	if newWidth < 0 {
		newWidth = 0
	}
	if newHeight < 0 {
		newHeight = 0
	}
	if newWidth == 0 && newHeight == 0 {
		return srcWidth, srcHeight
	}
	if newWidth == 0 {
		newWidth = srcWidth * newHeight / srcHeight
	}
	if newHeight == 0 {
		newHeight = srcHeight * newWidth / srcWidth
	}
	if newWidth < 1 {
		newWidth = 1
	}
	if newHeight < 1 {
		newHeight = 1
	}
	return newWidth, newHeight
}

func precomputeResizeAxis(srcSize, dstSize int) []resizeAxisSample {
	samples := make([]resizeAxisSample, dstSize)
	if srcSize <= 0 || dstSize <= 0 {
		return samples
	}

	last := srcSize - 1

	for i := 0; i < dstSize; i++ {
		numerator := i * srcSize
		i0 := numerator / dstSize
		if i0 > last {
			i0 = last
		}
		w := (numerator % dstSize) * resizeWeightScale / dstSize
		i1 := i0 + 1
		if i1 > last {
			i1 = last
		}
		samples[i] = resizeAxisSample{i0: i0, i1: i1, w: uint16(w)}
	}

	return samples
}

func shouldParallelResize(width, height int) bool {
	return width > 0 && height > 0 && width*height >= resizeParallelThresholdPixels && runtime.GOMAXPROCS(0) > 1
}

func resizeRGBA(src *image.RGBA, newWidth, newHeight int) *image.RGBA {
	bounds := src.Bounds()
	srcWidth := bounds.Dx()
	srcHeight := bounds.Dy()
	dst := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

	xSamples := precomputeResizeAxis(srcWidth, newWidth)
	ySamples := precomputeResizeAxis(srcHeight, newHeight)

	renderRows := func(startY, endY int) {
		for y := startY; y < endY; y++ {
			ys := ySamples[y]
			row0 := ys.i0 * src.Stride
			row1 := ys.i1 * src.Stride
			wy := uint64(ys.w)
			invWy := uint64(resizeWeightScale) - wy
			dstRow := y * dst.Stride

			for x := 0; x < newWidth; x++ {
				xs := xSamples[x]
				wx := uint64(xs.w)
				invWx := uint64(resizeWeightScale) - wx

				p00 := row0 + xs.i0*4
				p10 := row0 + xs.i1*4
				p01 := row1 + xs.i0*4
				p11 := row1 + xs.i1*4

				r00 := float64(src.Pix[p00])
				g00 := float64(src.Pix[p00+1])
				b00 := float64(src.Pix[p00+2])
				a00 := float64(src.Pix[p00+3])

				r10 := float64(src.Pix[p10])
				g10 := float64(src.Pix[p10+1])
				b10 := float64(src.Pix[p10+2])
				a10 := float64(src.Pix[p10+3])

				r01 := float64(src.Pix[p01])
				g01 := float64(src.Pix[p01+1])
				b01 := float64(src.Pix[p01+2])
				a01 := float64(src.Pix[p01+3])

				r11 := float64(src.Pix[p11])
				g11 := float64(src.Pix[p11+1])
				b11 := float64(src.Pix[p11+2])
				a11 := float64(src.Pix[p11+3])

				r0 := uint64(r00)*invWx + uint64(r10)*wx
				g0 := uint64(g00)*invWx + uint64(g10)*wx
				b0 := uint64(b00)*invWx + uint64(b10)*wx
				a0 := uint64(a00)*invWx + uint64(a10)*wx

				r1 := uint64(r01)*invWx + uint64(r11)*wx
				g1 := uint64(g01)*invWx + uint64(g11)*wx
				b1 := uint64(b01)*invWx + uint64(b11)*wx
				a1 := uint64(a01)*invWx + uint64(a11)*wx

				offset := dstRow + x*4
				dst.Pix[offset] = uint8((r0*invWy + r1*wy + resizeBilinearRound) >> 32)
				dst.Pix[offset+1] = uint8((g0*invWy + g1*wy + resizeBilinearRound) >> 32)
				dst.Pix[offset+2] = uint8((b0*invWy + b1*wy + resizeBilinearRound) >> 32)
				dst.Pix[offset+3] = uint8((a0*invWy + a1*wy + resizeBilinearRound) >> 32)
			}
		}
	}

	if !shouldParallelResize(newWidth, newHeight) {
		renderRows(0, newHeight)
		return dst
	}

	workers := runtime.GOMAXPROCS(0)
	if workers > newHeight {
		workers = newHeight
	}
	if workers < 1 {
		workers = 1
	}

	chunkSize := (newHeight + workers - 1) / workers
	var wg sync.WaitGroup
	for startY := 0; startY < newHeight; startY += chunkSize {
		endY := startY + chunkSize
		if endY > newHeight {
			endY = newHeight
		}
		wg.Add(1)
		go func(startY, endY int) {
			defer wg.Done()
			renderRows(startY, endY)
		}(startY, endY)
	}
	wg.Wait()

	return dst
}

func resizeGray(src *image.Gray, newWidth, newHeight int) *image.RGBA {
	bounds := src.Bounds()
	srcWidth := bounds.Dx()
	srcHeight := bounds.Dy()
	dst := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

	xSamples := precomputeResizeAxis(srcWidth, newWidth)
	ySamples := precomputeResizeAxis(srcHeight, newHeight)

	renderRows := func(startY, endY int) {
		for y := startY; y < endY; y++ {
			ys := ySamples[y]
			row0 := ys.i0 * src.Stride
			row1 := ys.i1 * src.Stride
			wy := uint64(ys.w)
			invWy := uint64(resizeWeightScale) - wy
			dstRow := y * dst.Stride

			for x := 0; x < newWidth; x++ {
				xs := xSamples[x]
				wx := uint64(xs.w)
				invWx := uint64(resizeWeightScale) - wx

				g00 := uint64(src.Pix[row0+xs.i0])
				g10 := uint64(src.Pix[row0+xs.i1])
				g01 := uint64(src.Pix[row1+xs.i0])
				g11 := uint64(src.Pix[row1+xs.i1])

				g0 := g00*invWx + g10*wx
				g1 := g01*invWx + g11*wx
				gray := uint8((g0*invWy + g1*wy + resizeBilinearRound) >> 32)

				offset := dstRow + x*4
				dst.Pix[offset] = gray
				dst.Pix[offset+1] = gray
				dst.Pix[offset+2] = gray
				dst.Pix[offset+3] = 0xff
			}
		}
	}

	if !shouldParallelResize(newWidth, newHeight) {
		renderRows(0, newHeight)
		return dst
	}

	workers := runtime.GOMAXPROCS(0)
	if workers > newHeight {
		workers = newHeight
	}
	if workers < 1 {
		workers = 1
	}

	chunkSize := (newHeight + workers - 1) / workers
	var wg sync.WaitGroup
	for startY := 0; startY < newHeight; startY += chunkSize {
		endY := startY + chunkSize
		if endY > newHeight {
			endY = newHeight
		}
		wg.Add(1)
		go func(startY, endY int) {
			defer wg.Done()
			renderRows(startY, endY)
		}(startY, endY)
	}
	wg.Wait()

	return dst
}

func resizeGray16(src *image.Gray16, newWidth, newHeight int) *image.RGBA {
	bounds := src.Bounds()
	srcWidth := bounds.Dx()
	srcHeight := bounds.Dy()
	dst := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

	xSamples := precomputeResizeAxis(srcWidth, newWidth)
	ySamples := precomputeResizeAxis(srcHeight, newHeight)

	for y := 0; y < newHeight; y++ {
		ys := ySamples[y]
		row0 := ys.i0 * src.Stride
		row1 := ys.i1 * src.Stride
		wy := uint64(ys.w)
		invWy := uint64(resizeWeightScale) - wy
		dstRow := y * dst.Stride

		for x := 0; x < newWidth; x++ {
			xs := xSamples[x]
			wx := uint64(xs.w)
			invWx := uint64(resizeWeightScale) - wx

			p00 := row0 + xs.i0*2
			p10 := row0 + xs.i1*2
			p01 := row1 + xs.i0*2
			p11 := row1 + xs.i1*2

			g00 := uint64(src.Pix[p00])
			g10 := uint64(src.Pix[p10])
			g01 := uint64(src.Pix[p01])
			g11 := uint64(src.Pix[p11])

			g0 := g00*invWx + g10*wx
			g1 := g01*invWx + g11*wx
			gray := uint8((g0*invWy + g1*wy + resizeBilinearRound) >> 32)

			offset := dstRow + x*4
			dst.Pix[offset] = gray
			dst.Pix[offset+1] = gray
			dst.Pix[offset+2] = gray
			dst.Pix[offset+3] = 0xff
		}
	}

	return dst
}

func resizeYCbCr(src *image.YCbCr, newWidth, newHeight int) *image.RGBA {
	bounds := src.Bounds()
	srcWidth := bounds.Dx()
	srcHeight := bounds.Dy()
	baseX := bounds.Min.X
	baseY := bounds.Min.Y
	dst := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

	xSamples := precomputeResizeAxis(srcWidth, newWidth)
	ySamples := precomputeResizeAxis(srcHeight, newHeight)

	for y := 0; y < newHeight; y++ {
		ys := ySamples[y]
		ay0 := baseY + ys.i0
		ay1 := baseY + ys.i1
		wy := uint64(ys.w)
		invWy := uint64(resizeWeightScale) - wy
		dstRow := y * dst.Stride

		for x := 0; x < newWidth; x++ {
			xs := xSamples[x]
			ax0 := baseX + xs.i0
			ax1 := baseX + xs.i1
			wx := uint64(xs.w)
			invWx := uint64(resizeWeightScale) - wx

			c00 := src.YCbCrAt(ax0, ay0)
			c10 := src.YCbCrAt(ax1, ay0)
			c01 := src.YCbCrAt(ax0, ay1)
			c11 := src.YCbCrAt(ax1, ay1)

			r00, g00, b00 := color.YCbCrToRGB(c00.Y, c00.Cb, c00.Cr)
			r10, g10, b10 := color.YCbCrToRGB(c10.Y, c10.Cb, c10.Cr)
			r01, g01, b01 := color.YCbCrToRGB(c01.Y, c01.Cb, c01.Cr)
			r11, g11, b11 := color.YCbCrToRGB(c11.Y, c11.Cb, c11.Cr)

			r0 := uint64(r00)*invWx + uint64(r10)*wx
			g0 := uint64(g00)*invWx + uint64(g10)*wx
			b0 := uint64(b00)*invWx + uint64(b10)*wx

			r1 := uint64(r01)*invWx + uint64(r11)*wx
			g1 := uint64(g01)*invWx + uint64(g11)*wx
			b1 := uint64(b01)*invWx + uint64(b11)*wx

			offset := dstRow + x*4
			dst.Pix[offset] = uint8((r0*invWy + r1*wy + resizeBilinearRound) >> 32)
			dst.Pix[offset+1] = uint8((g0*invWy + g1*wy + resizeBilinearRound) >> 32)
			dst.Pix[offset+2] = uint8((b0*invWy + b1*wy + resizeBilinearRound) >> 32)
			dst.Pix[offset+3] = 0xff
		}
	}

	return dst
}

// getRGBA 将 color.Color 转换为 uint8 的 RGBA 值
func getRGBA(c color.Color) (uint8, uint8, uint8, uint8) {
	r, g, b, a := c.RGBA()
	return uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)
}

// bilinearInterpolate 执行双线性插值
func bilinearInterpolate(c00, c10, c01, c11 color.Color, dx, dy float64) color.RGBA {
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

func resizeGeneric(src image.Image, newWidth, newHeight int) *image.RGBA {
	bounds := src.Bounds()
	srcWidth := bounds.Dx()
	srcHeight := bounds.Dy()
	baseX := bounds.Min.X
	baseY := bounds.Min.Y
	dst := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	xSamples := precomputeResizeAxis(srcWidth, newWidth)
	ySamples := precomputeResizeAxis(srcHeight, newHeight)

	for y := 0; y < newHeight; y++ {
		ys := ySamples[y]
		iy := ys.i0
		iy1 := ys.i1
		dy := float64(ys.w) / resizeWeightScale

		for x := 0; x < newWidth; x++ {
			xs := xSamples[x]
			ix := xs.i0
			ix1 := xs.i1
			dx := float64(xs.w) / resizeWeightScale

			c00 := src.At(baseX+ix, baseY+iy)
			c10 := src.At(baseX+ix1, baseY+iy)
			c01 := src.At(baseX+ix, baseY+iy1)
			c11 := src.At(baseX+ix1, baseY+iy1)

			dst.SetRGBA(x, y, bilinearInterpolate(c00, c10, c01, c11, dx, dy))
		}
	}

	return dst
}
