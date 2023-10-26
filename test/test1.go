package main

import (
	"github.com/dimkauzh/vuelto"
	w "github.com/dimkauzh/vuelto/src/window"
)

func main() {
	vuelto.Init()
	win := w.NewWindow("test", 800, 600)

	for win.Loop() {
	}
}
