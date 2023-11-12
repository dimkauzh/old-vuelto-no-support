package src

/*
#include "window.h"

#include <SDL2/SDL.h>
*/
import "C"

type Button struct {
	Window *Window
	Width  int
	Height int
	X      int
	Y      int
	Color  [3]int
}

// Creates new button with given parameters
func (w *Window) NewButton(x, y, width, height int, color [3]int) Button {
	return Button{w, x, y, width, height, color}
}

// Draws button on the screen
func (b *Button) Draw() {
	b.Window.DrawRect(b.X, b.Y, b.Width, b.Height, b.Color)
}

// Returns true if button is clicked
func (b *Button) IsClicked() bool {
	return bool(C.IsButtonPressed(C.int(b.X), C.int(b.Y), C.int(b.Width), C.int(b.Height)))
}
