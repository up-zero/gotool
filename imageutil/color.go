package imageutil

import "image/color"

var (
	ColorWhite   = color.RGBA{R: 255, G: 255, B: 255, A: 255} // 白色
	ColorBlack   = color.RGBA{R: 0, G: 0, B: 0, A: 255}       // 黑色
	ColorRed     = color.RGBA{R: 255, G: 0, B: 0, A: 255}     // 红色
	ColorGreen   = color.RGBA{R: 0, G: 255, B: 0, A: 255}     // 纯绿
	ColorBlue    = color.RGBA{R: 0, G: 0, B: 255, A: 255}     // 纯蓝
	ColorYellow  = color.RGBA{R: 255, G: 255, B: 0, A: 255}   // 黄色
	ColorCyan    = color.RGBA{R: 0, G: 255, B: 255, A: 255}   // 青色
	ColorMagenta = color.RGBA{R: 255, G: 0, B: 255, A: 255}   // 洋红/品红
	ColorOrange  = color.RGBA{R: 255, G: 165, B: 0, A: 255}   // 橙色
	ColorPurple  = color.RGBA{R: 128, G: 0, B: 128, A: 255}   // 紫色
	ColorPink    = color.RGBA{R: 255, G: 192, B: 203, A: 255} // 粉色
	ColorBrown   = color.RGBA{R: 165, G: 42, B: 42, A: 255}   // 棕色

	ColorSuccess = color.RGBA{R: 40, G: 167, B: 69, A: 255}  // 成功绿
	ColorDanger  = color.RGBA{R: 220, G: 53, B: 69, A: 255}  // 危险红
	ColorWarning = color.RGBA{R: 255, G: 193, B: 7, A: 255}  // 警告黄
	ColorInfo    = color.RGBA{R: 23, G: 162, B: 184, A: 255} // 信息蓝

	ColorLightGray = color.RGBA{R: 211, G: 211, B: 211, A: 255} // 浅灰色
	ColorGray      = color.RGBA{R: 128, G: 128, B: 128, A: 255} // 灰色
	ColorDarkGray  = color.RGBA{R: 169, G: 169, B: 169, A: 255} // 深灰色
	ColorSlate     = color.RGBA{R: 112, G: 128, B: 144, A: 255} // 岩板灰
	ColorSkyBlue   = color.RGBA{R: 135, G: 206, B: 235, A: 255} // 天蓝色

	ColorGold   = color.RGBA{R: 255, G: 215, B: 0, A: 255}   // 金色
	ColorSilver = color.RGBA{R: 192, G: 192, B: 192, A: 255} // 银色
	ColorCoral  = color.RGBA{R: 255, G: 127, B: 80, A: 255}  // 珊瑚色
	ColorOlive  = color.RGBA{R: 128, G: 128, B: 0, A: 255}   // 橄榄绿
)
