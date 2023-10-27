package window

/*
#include "window.h"

#include <SDL2/SDL.h>
*/
import "C"

type Window struct {
	Window   *C.SDL_Window
	Renderer *C.SDL_Renderer
	Title    string
	Width    int
	Height   int
}

func (w *Window) Loop() bool {
	C.SDL_RenderPresent(w.Renderer)
	C.SDL_SetRenderDrawColor(w.Renderer, 0, 0, 0, 255)
	C.SDL_RenderClear(w.Renderer)
	return bool(C.Loop(w.Window))
}

func NewWindow(title string, width, height int) *Window {
	w := C.NewWindow(C.CString(title), C.int(width), C.int(height))
	r := C.NewRenderer(w)
	return &Window{w, r, title, width, height}
}

func TestLoop(window Window) {
	C.TestLoop(window.Window)
}
