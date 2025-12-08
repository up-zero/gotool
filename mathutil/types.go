package mathutil

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Line struct {
	P1 Point `json:"p1"`
	P2 Point `json:"p2"`
}

type Rectangle struct {
	Min Point `json:"min"` // 左上角坐标
	Max Point `json:"max"` // 右下角坐标
}
