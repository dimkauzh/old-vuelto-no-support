package src

type Vector2D struct {
	x int
	y int
}

func (w *Window) NewVector2D(x, y int) Vector2D {
	return Vector2D{x, y}
}
