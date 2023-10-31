package main

import (
	"fmt"

	"github.com/dimkauzh/vuelto"
)

func main() {
	vuelto.Init()
	win := vuelto.NewWindow("Vuelto Test", 800, 600)

	image := win.LoadImage("test/image.png")

	for win.Loop() {
		win.SetBackgroundColor([3]int{200, 200, 200})

		image.DrawImage(0, 0, 500, 500)
		win.DrawRect(600, 200, 200, 200, [3]int{255, 0, 0})

		if win.IsKeyPressed(win.Key["A"]) {
			fmt.Println("A key pressed")
		}

		if win.IsKeyPressedOnce(win.Key["B"]) {
			fmt.Println("B key pressed")
		}
	}
}
