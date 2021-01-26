/*
* A tour of Go exercise: Images
 */

package main

import (
	"image"
	"image/color"
	"math"

	"golang.org/x/tour/pic"
)

type Image struct {
	w, h int
}

func (im Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (im Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, im.w, im.h)
}

func (im Image) At(x, y int) color.Color {
	var r, g, b uint8
	dx := float64(x) / float64(im.w) * 255
	dy := float64(y) / float64(im.h) * 255
	xint, _ := math.Modf(dx)
	x256 := uint8(xint)
	yint, _ := math.Modf(dy)
	y256 := uint8(yint)

	if x <= im.w/2 {
		r += x256
		b += 255 - x256
	} else {
		r += 255 - x256
		b += x256
	}

	if y <= im.h/2 {
		g += 255 - y256
		r += y256
	} else {
		r += 255 - y256
		g += y256
	}

	return color.RGBA{r, g, b, 255}
}

func main() {
	m := Image{200, 200}
	pic.ShowImage(m)
}
