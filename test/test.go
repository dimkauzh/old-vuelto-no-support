package main

import (
	"fmt"

	"github.com/dimkauzh/vuelto"
)

func main() {
	vuelto.Init()
	win := vuelto.NewWindow("Vuelto Test", 800, 600)

	img := win.LoadImage("test/image.png")

	win.SetFPS(60)
	win.SetIcon("test/image.bmp")

	for win.Loop() {
		win.SetBackgroundColor([3]int{255, 12, 76})

		img.Draw(0, 0, 100, 100)

		if win.IsKeyPressed(win.Key["A"]) {
			fmt.Println("A key pressed")
		}

		if win.IsKeyPressedOnce(win.Key["B"]) {
			fmt.Println("B key pressed")
		}
	}
}
