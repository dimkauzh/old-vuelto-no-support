package window

// #include <window.h>
import "C"

func NewWindow(title string, width, height int) *C.SDL_Window {
	return C.NewWindow(C.CString(title), C.int(width), C.int(height))
}

func Loop(window *C.SDL_Window) bool {
	return bool(C.Loop(window))
}
