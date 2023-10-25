package vuelto

/*
#cgo LDFLAGS: -lSDL2

#include <SDL2/SDL.h>
*/
import "C"

func Init() {
	C.SDL_Init(C.SDL_INIT_EVERYTHING)
}

func Quit() {
}
