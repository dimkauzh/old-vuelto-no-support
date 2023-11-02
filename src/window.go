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

// Force quits your application. This is a safer way to quit a application then vuelto.ForceQuit()
func (w *Window) ForceQuit() {
	C.Quit()
}
