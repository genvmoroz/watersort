package drawable

import (
	"image"
	"image/color"
	"image/draw"

	"watersort/model"
	"watersort/resources"
)

type Flask struct {
	flask *model.Flask

	res resources.Resources
}

var flaskColour = color.White

func NewFlask(flask *model.Flask, res resources.Resources) Flask {
	return Flask{flask: flask, res: res}
}

func (f *Flask) Draw(img *image.RGBA, rect image.Rectangle) {
	canvas := image.NewRGBA(image.Rect(0, 0, rect.Dx(), rect.Dy()))

	edgeWidth := rect.Dx() / 50 // edge width is 2% of total width

	f.drawEdges(canvas, edgeWidth)
	f.drawFlasks(canvas, 0+edgeWidth, 0+edgeWidth, canvas.Bounds().Dx()-edgeWidth*2, canvas.Bounds().Dy()-edgeWidth*2)

	draw.Draw(img, rect, canvas, canvas.Rect.Min, draw.Over)
}

func (f Flask) drawEdges(img *image.RGBA, width int) {
	for x := 0; x < img.Bounds().Dx(); x++ {
		for y := 0; y < width; y++ {
			img.Set(x, y, flaskColour)
			img.Set(x, img.Bounds().Dy()-y-1, flaskColour)
		}
	}
	for y := 0; y <= img.Bounds().Dy(); y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, flaskColour)
			img.Set(img.Bounds().Dx()-x-1, y, flaskColour)
		}
	}
}

func (f Flask) drawFlasks(img *image.RGBA, x, y, width, height int) {
	for xx := x; xx <= width; xx++ {
		for yy := y; yy <= height; yy++ {
			img.Set(xx, yy, color.NRGBA{
				R: 0xff,
				G: 0xff,
				B: 0,
				A: 0xff,
			})
		}
	}
}
