package main

/*

#cgo CFLAGS: -I../lib/sdl2/include
#cgo LDFLAGS: -L../lib/sdl2/lib -lSDL2

#include "draw.c"
*/
import "C"

func main() {
	C.window_test()
}
