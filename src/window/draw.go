package window

/*
#include <SDL2/SDL.h>

#include "draw.h"
*/
import "C"

func (w *Window) SetPixel(x, y int, color [3]int) {
	r := w.Renderer
	C.SetPixel(r, C.int(x), C.int(y), C.int(color[0]), C.int(color[1]), C.int(color[2]))
}

func (w *Window) DrawRect(x, y, width, height int, color [3]int) {
	C.DrawRect(w.Renderer, C.int(x), C.int(y), C.int(width), C.int(height), C.int(color[0]), C.int(color[1]), C.int(color[2]))
}
