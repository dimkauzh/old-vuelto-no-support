package main

import (
	"github.com/dimkauzh/vuelto"
)

func main() {
	vuelto.Init()
	win := vuelto.NewWindow("Vuelto", 800, 600)

	for win.Loop() {
		win.SetBackgroundColor([3]int{64, 64, 64})
		win.DrawRect(100, 100, 100, 100, [3]int{0, 255, 0})
		win.DrawRect(200, 300, 100, 100, [3]int{0, 255, 0})
	}
}
