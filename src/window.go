package src

/*
#include "window.h"

#include <SDL2/SDL.h>
*/
import "C"

// The window struct
type Window struct {
	Window *C.SDL_Window
	Title  string
	Width  int
	Height int

	Key map[string]C.SDL_Scancode
}

// A loop function which you use to run a loop. It returns a boolean value
// which you can use to start and break. Use if like this: for your_window.Loop() {}
func (w *Window) Loop() bool {
	return bool(C.Loop())
}

// Do not use! To create a window run vuelto.NewWindow(), this is a internal function
func NewWindow(title string, width, height int) *Window {
	return &Window{C.NewWindow(C.CString(title), C.int(width), C.int(height)), title, width, height, Key}
}

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

// Force quits your application. This is a safer way to quit a application then vuelto.ForceQuit()
func (w *Window) ForceQuit() {
	C.Quit()
}
