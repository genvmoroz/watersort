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
	f.drawFlasks(canvas, edgeWidth, edgeWidth, canvas.Bounds().Dx()-edgeWidth*2, canvas.Bounds().Dy()-edgeWidth*2+1)

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

func (f Flask) drawFlasks(img *image.RGBA, x, _, width, height int) {
	colours := f.flask.All()

	border := height / len(colours)

	for i := 0; i < len(colours); i++ {
		f.drawFlask(img, colours[len(colours)-1-i], x, i*border, width, border)
	}
}

func (f Flask) drawFlask(img *image.RGBA, colour model.Colour, x, y, width, height int) {
	if colour == model.Non {
		return
	}

	c := f.pickColor(colour)

	for xx := 0; xx < width; xx++ {
		for yy := 0; yy < height; yy++ {
			img.Set(xx+x, yy+y, c)
		}
	}
}

func (f Flask) pickColor(colour model.Colour) color.Color {
	switch colour {
	case model.Pink:
		return color.RGBA{R: 232, G: 94, B: 122, A: 0xff}
	case model.Purple:
		return color.RGBA{R: 113, G: 38, B: 147, A: 0xff}
	case model.Red:
		return color.RGBA{R: 196, G: 43, B: 34, A: 0xff}
	case model.Gray:
		return color.RGBA{R: 99, G: 100, B: 102, A: 0xff}
	case model.Yellow:
		return color.RGBA{R: 243, G: 217, B: 48, A: 0xff}
	case model.Brown:
		return color.RGBA{R: 134, G: 69, B: 0, A: 0xff}
	case model.Orange:
		return color.RGBA{R: 246, G: 134, B: 42, A: 0xff}
	case model.Blue:
		return color.RGBA{R: 60, G: 34, B: 184, A: 0xff}
	case model.LightBlue:
		return color.RGBA{R: 49, G: 156, B: 224, A: 0xff}
	case model.DarkGreen:
		return color.RGBA{R: 0, G: 103, B: 44, A: 0xff}
	case model.Green:
		return color.RGBA{R: 112, G: 152, B: 0, A: 0xff}
	case model.LightGreen:
		return color.RGBA{R: 0, G: 219, B: 114, A: 0xff}
	default:
		return color.RGBA{R: 255, G: 255, B: 255, A: 0xff}
	}
}
