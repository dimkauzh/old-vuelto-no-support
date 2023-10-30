package src

/*
#include "window.h"

#include <SDL2/SDL.h>
*/
import "C"

type Window struct {
	Window *C.SDL_Window
	Title  string
	Width  int
	Height int

	Key map[string]C.SDL_Scancode
}

func (w *Window) Loop() bool {
	return bool(C.Loop())
}

func NewWindow(title string, width, height int) *Window {
	c_w := C.NewWindow(C.CString(title), C.int(width), C.int(height))
	w := &Window{c_w, title, width, height, nil}
	w.Key = Key
	return w
}

func (w *Window) SetPixel(x, y int, color [3]int) {
	C.SetPixel(C.int(x), C.int(y), C.int(color[0]), C.int(color[1]), C.int(color[2]))
}

func (w *Window) ForceQuit() {
	C.Quit()
}
