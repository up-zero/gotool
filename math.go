package gotool

type Number interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 | uintptr |
		float32 | float64
}

func MathAbs[T Number](v T) T {
	if v < 0 {
		return -v
	}
	return v
}
