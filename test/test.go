package main

import (
	"github.com/dimkauzh/vuelto"
	"github.com/dimkauzh/vuelto/src/window"
)

func main() {
	vuelto.Init()
	win := window.NewWindow("test", 800, 600)
	image := win.LoadImage("test/test.png")

	for win.Loop() {
		image.DrawImage(0, 0, 500, 500)
		win.DrawRect(600, 200, 200, 200, [3]int{255, 0, 0})
	}
}
