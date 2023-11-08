package main

import (
	"github.com/dimkauzh/vuelto"
)

const SPEED = 5

func main() {
	vuelto.Init()
	win := vuelto.NewWindow("Vuelto Game Test", 800, 600)
	//win.RemoveFPSLimit()

	x, y := 0, 0

	image := win.LoadImage("test/coin.png")

	for win.Loop() {
		win.SetBackgroundColor([3]int{64, 64, 64})

		if win.IsKeyPressed(win.Key["Up"]) {
			y = y - SPEED
		} else if win.IsKeyPressed(win.Key["Down"]) {
			y = y + SPEED
		} else if win.IsKeyPressed(win.Key["Left"]) {
			x = x - SPEED
		} else if win.IsKeyPressed(win.Key["Right"]) {
			x = x + SPEED
		}

		image.Draw(x, y, 50, 50)

	}
}
