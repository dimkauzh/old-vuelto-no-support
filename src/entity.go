package src

type Entity struct {
	X      int
	Y      int
	Width  int
	Height int
	Image  *Image
	Rect   *Rect
}

// Creates a new entity that you can later connect a rect or image to.
func (w *Window) NewEntity(x int, y int, width int, height int) Entity {
	return Entity{x, y, width, height, nil, nil}
}

// Connects a image to the entity. You can later draw it like this: your_entity.Image.DrawImage()
func (e *Entity) SetImage(image Image) {
	e.Image = &image
}

// Connects a rect to the entity. You can later draw it like this: your_entity.Rect.Draw()
func (e *Entity) SetRect(rect Rect) {
	e.Rect = &rect
}
