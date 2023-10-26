package window

// #cgo LDFLAGS: -lSDL2
// #include <window.h>
// #include <SDL2/SDL.h>
import "C"

type Window struct {
	window   *C.SDL_Window
	renderer *C.SDL_Renderer
	title    string
	width    int
	height   int
}

func (w *Window) Loop() bool {
	return bool(C.Loop(w.window))
}

func NewWindow(title string, width, height int) Window {
	w := C.NewWindow(C.CString(title), C.int(width), C.int(height))
	r := C.NewRenderer(w)
	return Window{w, r, title, width, height}
}

func TestLoop(window Window) {
	C.TestLoop(window.window)
}
