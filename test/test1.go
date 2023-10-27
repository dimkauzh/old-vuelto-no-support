package main

import (
	"github.com/dimkauzh/vuelto"
	"github.com/dimkauzh/vuelto/src/window"
)

func main() {
	vuelto.Init()
	win := window.NewWindow("test", 800, 600)

	for win.Loop() {
		win.DrawRect(100, 100, 100, 100, [3]int{0, 255, 0})
		win.DrawRect(200, 300, 100, 100, [3]int{0, 255, 0})

	}
}
