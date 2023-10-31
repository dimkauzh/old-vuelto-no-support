package src

type Rect struct {
	Window *Window
	X      int
	Y      int
	Width  int
	Height int
	Color  [3]int
}

// Creates a new rectangle that you can later draw with the function Draw(), like this: your_rect.Draw()
func (w *Window) NewRect(x, y, width, height int, color [3]int) Rect {
	return Rect{w, x, y, width, height, color}
}

// The draw function that you run to draw your rectangle that you created.
func (r *Rect) Draw() {
	r.Window.DrawRect(r.X, r.Y, r.Width, r.Height, r.Color)
}
