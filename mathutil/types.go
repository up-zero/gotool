package mathutil

type Point struct {
	X, Y float64
}

type Line struct {
	P1, P2 Point
}

type Rectangle struct {
	Min Point
	Max Point
}
