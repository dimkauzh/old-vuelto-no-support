package keys

/*
#include <SDL2/SDL.h>

#include "keys.h"
*/
import "C"

var keyPressed bool

func IsKeyPressed(key C.SDL_Scancode) bool {
	return bool(C.IsKeyPressed(key))
}

func IsKeyPressedOnce(key C.SDL_Scancode) bool {
	if IsKeyPressed(key) && !keyPressed {
		keyPressed = true
		return true
	}
	if !IsKeyPressed(key) {
		keyPressed = false
	}
	return false
}

var Key = map[string]C.SDL_Scancode{
	"A":       C.SDL_SCANCODE_A,
	"B":       C.SDL_SCANCODE_B,
	"C":       C.SDL_SCANCODE_C,
	"D":       C.SDL_SCANCODE_D,
	"E":       C.SDL_SCANCODE_E,
	"F":       C.SDL_SCANCODE_F,
	"G":       C.SDL_SCANCODE_G,
	"H":       C.SDL_SCANCODE_H,
	"I":       C.SDL_SCANCODE_I,
	"J":       C.SDL_SCANCODE_J,
	"K":       C.SDL_SCANCODE_K,
	"L":       C.SDL_SCANCODE_L,
	"M":       C.SDL_SCANCODE_M,
	"N":       C.SDL_SCANCODE_N,
	"O":       C.SDL_SCANCODE_O,
	"P":       C.SDL_SCANCODE_P,
	"Q":       C.SDL_SCANCODE_Q,
	"R":       C.SDL_SCANCODE_R,
	"S":       C.SDL_SCANCODE_S,
	"T":       C.SDL_SCANCODE_T,
	"U":       C.SDL_SCANCODE_U,
	"V":       C.SDL_SCANCODE_V,
	"W":       C.SDL_SCANCODE_W,
	"X":       C.SDL_SCANCODE_X,
	"Y":       C.SDL_SCANCODE_Y,
	"Z":       C.SDL_SCANCODE_Z,
	"Num0":    C.SDL_SCANCODE_0,
	"Num1":    C.SDL_SCANCODE_1,
	"Num2":    C.SDL_SCANCODE_2,
	"Num3":    C.SDL_SCANCODE_3,
	"Num4":    C.SDL_SCANCODE_4,
	"Num5":    C.SDL_SCANCODE_5,
	"Num6":    C.SDL_SCANCODE_6,
	"Num7":    C.SDL_SCANCODE_7,
	"Num8":    C.SDL_SCANCODE_8,
	"Num9":    C.SDL_SCANCODE_9,
	"Space":   C.SDL_SCANCODE_SPACE,
	"Enter":   C.SDL_SCANCODE_RETURN,
	"Escape":  C.SDL_SCANCODE_ESCAPE,
	"Tab":     C.SDL_SCANCODE_TAB,
	"Shift":   C.SDL_SCANCODE_LSHIFT,
	"Control": C.SDL_SCANCODE_LCTRL,
	"Alt":     C.SDL_SCANCODE_LALT,
	"F1":      C.SDL_SCANCODE_F1,
	"F2":      C.SDL_SCANCODE_F2,
	"F3":      C.SDL_SCANCODE_F3,
	"F4":      C.SDL_SCANCODE_F4,
	"F5":      C.SDL_SCANCODE_F5,
	"F6":      C.SDL_SCANCODE_F6,
	"F7":      C.SDL_SCANCODE_F7,
	"F8":      C.SDL_SCANCODE_F8,
	"F9":      C.SDL_SCANCODE_F9,
	"F10":     C.SDL_SCANCODE_F10,
	"F11":     C.SDL_SCANCODE_F11,
	"F12":     C.SDL_SCANCODE_F12,
}
