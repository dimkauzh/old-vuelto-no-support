package main

import (
	"fmt"

	"github.com/dimkauzh/vuelto"
)

func main() {
	vuelto.Init()
	win := vuelto.NewWindow("Vuelto Test", 800, 600)

	img := win.LoadImage("test/image.png")

	button := win.NewButton(100, 100, 100, 100, [3]int{255, 0, 0})

	//audio := win.OpenAudioFile("test/audio.wav")
	//defer audio.Close()

	win.SetFPS(60)
	win.SetIcon("test/image.bmp")

	//audio.Start()

	for win.Loop() {
		win.SetBackgroundColor([3]int{69, 69, 69})

		img.Draw(0, 0, 100, 100)

		button.Draw()

		if button.IsClicked() {
			fmt.Println("Button clicked")
		}

		if win.IsKeyPressed(win.Key["A"]) {
			fmt.Println("A key pressed")

		}

		if win.IsKeyPressedOnce(win.Key["B"]) {
			fmt.Println("B key pressed")
		}
	}
}
