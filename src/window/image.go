package window

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

type Image struct {
	Window   *Window
	FilePath string
	Image    image.Image
}

func (w *Window) LoadImage(filename string) Image {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Could not open: ", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal("Could not decode: ", err)
	}

	return Image{w, filename, img}
}

func (srcImage *Image) DrawImage(x, y, width, height int) {
	src := srcImage.Image
	srcBounds := src.Bounds()
	srcMinX, srcMinY := srcBounds.Min.X, srcBounds.Min.Y
	srcMaxX, srcMaxY := srcBounds.Max.X, srcBounds.Max.Y

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			srcX := srcMinX + i*srcMaxX/width
			srcY := srcMinY + j*srcMaxY/height

			if srcX >= srcMaxX {
				srcX = srcMaxX - 1
			}
			if srcY >= srcMaxY {
				srcY = srcMaxY - 1
			}

			srcColor := src.At(srcX, srcY)
			r, g, b, _ := srcColor.RGBA()
			srcImage.Window.SetPixel(x+i, y+j, [3]int{int(r >> 8), int(g >> 8), int(b >> 8)})
		}
	}
}