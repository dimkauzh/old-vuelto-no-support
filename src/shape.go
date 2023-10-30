package src

type Rect struct {
	Window *Window
	X      int
	Y      int
	Width  int
	Height int
	Color  [3]int
}

func (w *Window) NewRect(x, y, width, height int, color [3]int) Rect {
	return Rect{w, x, y, width, height, color}
}

func (r *Rect) Draw() {
	r.Window.DrawRect(r.X, r.Y, r.Width, r.Height, r.Color)
}
