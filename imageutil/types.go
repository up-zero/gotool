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

// Blob 存储单个连通区域的特征
type Blob struct {
	ID       int
	Points   []image.Point   // 原始像素点集合
	Area     int             // 面积
	Bounds   image.Rectangle // 正外接矩形 (AABB)
	Centroid image.Point     // 质心
}

// BlobResult 处理后的所有结果
type BlobResult struct {
	Blobs  []Blob
	Width  int
	Height int
}
