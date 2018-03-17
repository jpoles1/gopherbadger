package main

import (
	"image"
	"image/color"

	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
)

func createCanvas(width int, height int) (*image.RGBA, *draw2dimg.GraphicContext) {
	dest := image.NewRGBA(image.Rect(0, 0, width, height))
	gc := draw2dimg.NewGraphicContext(dest)
	return dest, gc
}
func createRoundedRectangle(gc *draw2dimg.GraphicContext, fillColor color.RGBA, borderColor color.RGBA, x, y, width, height, xradius, yradius float64) {
	x = x + xradius
	y = y + yradius
	width = width - xradius*2
	height = height - yradius*2
	gc.SetStrokeColor(borderColor)
	gc.SetLineWidth(10)
	draw2dkit.RoundedRectangle(gc, x, y, x+width, y+height, xradius, yradius)
	gc.SetFillColor(fillColor)
	gc.FillStroke()
}
func drawBadge(coveragePct float64) {
	//Grey
	colorGrey := color.RGBA{0x33, 0x33, 0x33, 0xff}
	//Green: >= 80% overall coverage
	colorGreen := color.RGBA{0x05, 0xa0, 0x00, 0xff}
	colorDarkGreen := color.RGBA{0x04, 0x91, 0x00, 0xff}
	//Yellow: 65% <= overall coverage < 80%
	//colorYellow := color.RGBA{0xe2, 0xbd, 0x00, 0xff}
	//Red: < 65% overall coverage
	//colorRed := color.RGBA{0xdb, 0x1a, 0x08, 0xff}
	// Initialize the graphic context on an RGBA image
	dest, gc := createCanvas(600, 120)

	createRoundedRectangle(gc, colorGreen, colorDarkGreen, 0, 0, 600, 120, 5, 5)
	createRoundedRectangle(gc, colorGrey, colorGrey, 15, 15, 400, 110, 0, 0)
	draw2d.SetFontFolder("./fonts")
	gc.SetStrokeColor(color.White)
	gc.StrokeStringAt("Coverage", 20, 20) //(width float64)
	// Save to file
	draw2dimg.SaveToPngFile("badge.png", dest)
}
