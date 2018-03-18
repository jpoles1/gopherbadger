package main

import (
	"errors"
	"fmt"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
)

func loadGoFontFace(points float64) (font.Face, error) {
	f, err := truetype.Parse(goregular.TTF)
	if err != nil {
		return nil, err
	}
	face := truetype.NewFace(f, &truetype.Options{
		Size: points,
		// Hinting: font.HintingFull,
	})
	return face, nil
}

func drawBadge(coveragePct float64, filename string) error {
	//Grey
	colorGrey := "#777"
	colorDarkGrey := "#333"
	//Green: >= 80% overall coverage
	colorGreen := "#00cc1e"
	colorDarkGreen := "#049100"
	//Yellow: 65% <= overall coverage < 80%
	colorYellow := "#e2bd00"
	colorDarkYellow := "#c6a601"
	//Red: < 65% overall coverage
	colorRed := "#db1a08"
	colorDarkRed := "#a31204"
	var accentColor, accentBorderColor string
	if coveragePct >= 80 {
		accentColor = colorGreen
		accentBorderColor = colorDarkGreen
	} else if coveragePct >= 55 {
		accentColor = colorYellow
		accentBorderColor = colorDarkYellow
	} else if coveragePct >= 0 {
		accentColor = colorRed
		accentBorderColor = colorDarkRed
	} else {
		return errors.New("Coverage value must be >= 0%")
	}
	//Create graphics context
	dc := gg.NewContext(600, 120)

	//Draw background rectangle
	dc.DrawRoundedRectangle(6, 6, 600-6*2, 120-6*2, 10)
	dc.SetHexColor(accentColor)
	dc.FillPreserve()
	dc.SetHexColor(accentBorderColor)
	dc.SetLineWidth(6.0)
	dc.Stroke()

	//Draw coverage background rectangle
	dc.DrawRoundedRectangle(10, 10, 410-10*2, 120-10*2, 5)
	dc.SetHexColor(colorDarkGrey)
	dc.FillPreserve()
	dc.SetHexColor(colorGrey)
	dc.SetLineWidth(2.0)
	dc.Stroke()

	//Drawing text
	fontFace, err := loadGoFontFace(72)
	errCheck("Loading default font-face.", err)
	dc.SetFontFace(fontFace)
	dc.SetHexColor("#ffffffff")
	dc.DrawString("Coverage:", 5+30, 120-5*2-30)
	covPctString := fmt.Sprintf("%2.f", coveragePct) + "%"
	dc.DrawString(covPctString, 410+15, 120-5*2-25)
	//Save to file
	err = dc.SavePNG(filename)
	errCheck("Saving image file", err)
	return err
}
