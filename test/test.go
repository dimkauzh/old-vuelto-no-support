package main

import (
	"fmt"

	"github.com/dimkauzh/vuelto"
	"github.com/dimkauzh/vuelto/src/keys"
	"github.com/dimkauzh/vuelto/src/window"
)

func main() {
	vuelto.Init()
	win := window.NewWindow("test", 800, 600)
	image := win.LoadImage("test/test.png")

	for win.Loop() {
		image.DrawImage(0, 0, 500, 500)
		win.DrawRect(600, 200, 200, 200, [3]int{255, 0, 0})

		if keys.IsKeyPressed(keys.Key["A"]) {
			fmt.Println("A key pressed")
		}

		if keys.IsKeyPressedOnce(keys.Key["A"]) {
			fmt.Println("A key pressed")
		}
	}
}
