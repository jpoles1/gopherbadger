package main

import (
	"fmt"

	"github.com/fogleman/gg"
)

func drawBadge(coveragePct float64, filename string) {
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
	} else {
		accentColor = colorRed
		accentBorderColor = colorDarkRed
	}
	//Create graphics context
	dc := gg.NewContext(600, 120)

	//Draw background rectangle
	dc.DrawRoundedRectangle(0, 0, 600, 120, 10)
	dc.SetHexColor(accentColor)
	dc.FillPreserve()
	dc.SetHexColor(accentBorderColor)
	dc.SetLineWidth(10.0)
	dc.Stroke()

	//Draw coverage background rectangle
	dc.DrawRoundedRectangle(5, 5, 410-5*2, 120-5*2, 5)
	dc.SetHexColor(colorDarkGrey)
	dc.FillPreserve()
	dc.SetHexColor(colorGrey)
	dc.SetLineWidth(2.0)
	dc.Stroke()

	//Drawing text
	err := dc.LoadFontFace("fonts/luxisr.ttf", 84)
	errCheck("Loading font", err)
	dc.SetHexColor("#ffffffff")
	dc.DrawString("Coverage:", 5+10, 120-5*2-25)
	covPctString := fmt.Sprintf("%2.f", coveragePct) + "%"
	fmt.Println("COV GEN:", covPctString)
	dc.DrawString(covPctString, 410+10, 120-5*2-22)
	//Save to file
	dc.SavePNG(filename)
}
