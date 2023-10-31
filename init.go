// A small CGo Game Engine that uses a custom Software Renderer.
package vuelto

/*
#cgo LDFLAGS: -lSDL2

#include <SDL2/SDL.h>

void Init() {
    if (SDL_Init(SDL_INIT_VIDEO) < 0) {
        printf("error initializing SDL: %s\n", SDL_GetError());
        SDL_Quit();
    }
}
*/
import "C"
import (
	"runtime"

	"github.com/dimkauzh/vuelto/src"
)

// All structs
type Window = src.Window
type Vector2D = src.Vector2D
type Rect = src.Rect
type Event = src.Event
type Image = src.Image
type Entity = src.Entity

// The init function. Run it before all other vuelto functions.
func Init() {
	runtime.LockOSThread()
	C.Init()
}

// Creates a new window. This returns a Window struct which you can use to draw things on the screen.
func NewWindow(title string, width, height int) *src.Window {
	return src.NewWindow(title, width, height)
}

// Deprecated: This can cause a memory leak. Please use: your_window.ForceQuit()
func ForceQuit() {
	C.SDL_Quit()
}
