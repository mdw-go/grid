package grid

type Direction[T Numeric] struct{ dx, dy T }

func NewDirection[T Numeric](dx, dy T) Direction[T] {
	return Direction[T]{dx: dx, dy: dy}
}

func (this Direction[T]) Dx() T { return this.dx }
func (this Direction[T]) Dy() T { return this.dy }

func (this Direction[T]) TurnRight() Direction[T] { return Clockwise(this) }
func (this Direction[T]) TurnLeft() Direction[T]  { return CounterClockwise(this) }

func Neighbors4[T Numeric]() []Direction[T] {
	return []Direction[T]{
		NewDirection[T](1, 0),
		NewDirection[T](-1, 0),
		NewDirection[T](0, 1),
		NewDirection[T](0, -1),
	}
}
func Neighbors8[T Numeric]() []Direction[T] {
	return append(Neighbors4[T](),
		NewDirection[T](1, 1),
		NewDirection[T](-1, 1),
		NewDirection[T](1, -1),
		NewDirection[T](-1, -1),
	)
}
func Clockwise[T Numeric](from Direction[T]) Direction[T] {
	switch from {
	case NewDirection[T](0, -1):
		return NewDirection[T](-1, 0)
	case NewDirection[T](-1, 0):
		return NewDirection[T](0, 1)
	case NewDirection[T](0, 1):
		return NewDirection[T](1, 0)
	case NewDirection[T](1, 0):
		return NewDirection[T](0, -1)
	default:
		return from
	}
}
func CounterClockwise[T Numeric](from Direction[T]) Direction[T] {
	switch from {
	case NewDirection[T](0, -1):
		return NewDirection[T](1, 0)
	case NewDirection[T](1, 0):
		return NewDirection[T](0, 1)
	case NewDirection[T](0, 1):
		return NewDirection[T](-1, 0)
	case NewDirection[T](-1, 0):
		return NewDirection[T](0, -1)
	default:
		return from
	}
}
