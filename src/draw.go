package src

// Draws a rectangle on the screen with a specific color.
func (w *Window) DrawRect(x, y, width, height int, color [3]int) {
	for i := x; i < x+width; i++ {
		for j := y; j < y+height; j++ {
			w.SetPixel(i, j, [3]int{color[0], color[1], color[2]})
		}
	}
}

// Sets the background color of the screen.
func (w *Window) SetBackgroundColor(color [3]int) {
	w.DrawRect(0, 0, w.Width, w.Height, color)
}
