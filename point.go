package grid

import "fmt"

type Point[T Numeric] struct {
	x T
	y T
}

func NewPoint[T Numeric](x, y T) Point[T] {
	return Point[T]{x: x, y: y}
}

func (this Point[T]) Y() T { return this.y }
func (this Point[T]) X() T { return this.x }

func (this Point[T]) Move(d Direction[T]) Point[T] {
	return NewPoint[T](this.x+d.dx, this.y+d.dy)
}
func (this Point[T]) Offset(x, y T) Point[T] {
	return NewPoint[T](this.x+x, this.y+y)
}
func (this Point[T]) String() string {
	return fmt.Sprintf("(%v, %v)", this.x, this.y)
}
func (this Point[T]) Neighbors4() (neighbors []Point[T]) {
	for _, offset := range Neighbors4[T]() {
		neighbors = append(neighbors, this.Offset(offset.dx, offset.dy))
	}
	return neighbors
}
func (this Point[T]) Neighbors8() (neighbors []Point[T]) {
	for _, offset := range Neighbors8[T]() {
		neighbors = append(neighbors, this.Offset(offset.dx, offset.dy))
	}
	return neighbors
}
