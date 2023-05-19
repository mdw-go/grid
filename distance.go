package grid

import "math"

func CityBlockDistance[T Numeric](p, q Point[T]) int {
	x, y := diff(p, q)
	return int(abs(x) + abs(y))
}
func EuclideanDistance[T Numeric](p, q Point[T]) float64 {
	x, y := diff(p, q)
	return math.Hypot(float64(x), float64(y))
}
func diff[T Numeric](p, q Point[T]) (x, y T) {
	return p.X() - q.X(), p.Y() - q.Y()
}
func abs[T Numeric](v T) T {
	if v < 0 {
		return -v
	}
	return v
}
