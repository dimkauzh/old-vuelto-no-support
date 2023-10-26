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
import "runtime"

func Init() {
	runtime.LockOSThread()
	C.Init()
}

func Quit() {
}
