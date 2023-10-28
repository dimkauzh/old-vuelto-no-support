package entity

import (
    "github.com/dimkauzh/vuelto/src/window"
)

type Entity struct {
    X      int
    Y      int
    Width  int
    Height int
    Image  *window.Image
    Rect   *window.Rect
}

func NewEntity(x int, y int, width int, height int) Entity {
    return Entity{x, y, width, height, nil, nil}
}

func (e *Entity) SetImage(image window.Image) {
    e.Image = &image
}

func (e *Entity) SetRect(rect window.Rect) {
    e.Rect = &rect
}
