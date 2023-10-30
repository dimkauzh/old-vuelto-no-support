package main

import (
	"fmt"

	"github.com/dimkauzh/vuelto"
	"github.com/dimkauzh/vuelto/src/keys"
	"github.com/dimkauzh/vuelto/src/window"
)

func main() {
	vuelto.Init()
	win := window.NewWindow("Vuelto Test", 800, 600)

	image := win.LoadImage("test/image.png")

	for win.Loop() {
		win.SetBackgroundColor([3]int{200, 200, 200})

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
