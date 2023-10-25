package main

import (
	"github.com/dimkauzh/vuelto"
	"github.com/dimkauzh/vuelto/src/window"
)

func main() {
	vuelto.Init()
	w := window.NewWindow("test", 800, 600)
	for window.Loop(w) {
	}
	vuelto.Quit()
}
