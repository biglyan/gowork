package main

import (
	"golang.org/x/tour/pic"
)

/*type Image struct {
	w int
	h int
}

func (image Image) ColorModel() color.Model {

}

func (image Image) Bounds() Rectangle {
	return image.Rect()
}

func (image Image) At(x, y int) color.Color {

}*/

func main() {
	m := Image{}
	pic.ShowImage(m)
}
