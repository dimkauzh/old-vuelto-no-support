package src

type Vector2D struct {
	X int
	Y int
}

// Create a new vector that stores two values: x and y.
func (w *Window) NewVector2D(x, y int) Vector2D {
	return Vector2D{x, y}
}
