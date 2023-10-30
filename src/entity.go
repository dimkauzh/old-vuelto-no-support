package src

type Entity struct {
    X      int
    Y      int
    Width  int
    Height int
    Image  *Image
    Rect   *Rect
}

func (w *Window) NewEntity(x int, y int, width int, height int) Entity {
    return Entity{x, y, width, height, nil, nil}
}

func (e *Entity) SetImage(image Image) {
    e.Image = &image
}

func (e *Entity) SetRect(rect Rect) {
    e.Rect = &rect
}
