package window

func (w *Window) DrawRect(x, y, width, height int, color [3]int) {
	for i := x; i < x+width; i++ {
		for j := y; j < y+height; j++ {
			w.SetPixel(i, j, [3]int{color[0], color[1], color[2]})
		}
	}
}
