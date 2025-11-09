package imageutil

import "image"

type SizeReply struct {
	Width  int // 图片宽
	Height int // 图片高
}

// StructuringElement 结构元素
type StructuringElement struct {
	// Kernel 形状, true 表示该像素在核的范围内
	Kernel [][]bool
	// Anchor 是核的中心点（锚点）
	Anchor image.Point
}
