package src

/*
#include "window.h"
*/
import "C"

// Sets a pixel on the screen with a color you profided.
func (w *Window) SetPixel(x, y int, color [3]int) {
	C.SetPixel(C.int(x), C.int(y), C.int(color[0]), C.int(color[1]), C.int(color[2]))
}

// Draws a rectangle on the screen with a specific color.
func (w *Window) DrawRect(x, y, width, height int, color [3]int) {
	C.DrawRect(C.int(x), C.int(y), C.int(width), C.int(height), C.int(color[0]), C.int(color[1]), C.int(color[2]))
}

// Sets the background color of the screen.
func (w *Window) SetBackgroundColor(color [3]int) {
	C.SetBackgroundColor(C.int(color[0]), C.int(color[1]), C.int(color[2]))
}
