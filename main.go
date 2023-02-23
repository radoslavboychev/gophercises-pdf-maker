package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New(gofpdf.OrientationPortrait, gofpdf.UnitPoint, gofpdf.PageSizeLetter, "")
	w, h := pdf.GetPageSize()
	fmt.Printf("width=%v, height=%v\n", w, h)
	pdf.AddPage()

	// Text stuff
	pdf.MoveTo(0, 0)
	pdf.SetFont("arial", "B", 30)
	_, lineHt := pdf.GetFontSize()
	pdf.SetTextColor(255, 0, 0)
	pdf.Text(0, lineHt, "Hello, world!")
	pdf.MoveTo(0, lineHt*2.0)

	pdf.SetFont("times", "", 18)
	pdf.SetTextColor(100, 100, 100)
	_, lineHt = pdf.GetFontSize()
	pdf.MultiCell(0, lineHt, "Here is some text. If it is too long it will be word wrapped automatically. If there is a new line it will be\nwrapped as well (unlike other ways of writing text in gofpdf).",
		gofpdf.BorderNone,
		gofpdf.AlignRight,
		false)

	//
	pdf.SetFillColor(0, 255, 0)
	pdf.SetDrawColor(0, 0, 255)
	pdf.Rect(10, 100, 100, 100, "FD")
	pdf.SetFillColor(100, 200, 200)
	pdf.Polygon([]gofpdf.PointType{
		{X: 110, Y: 250},
		{X: 160, Y: 300},
		{X: 110, Y: 350},
		{X: 60, Y: 300},
		{X: 70, Y: 230},
	}, "F")

	//
	pdf.ImageOptions("images/puppy.png", 275, 275, 92, 0, false, gofpdf.ImageOptions{
		ReadDpi: true,
	}, 0, "")

	// Grid
	drawGrid(pdf)

	err := pdf.OutputFileAndClose("p1.pdf")
	if err != nil {
		panic(err)
	}

}

func drawGrid(pdf *gofpdf.Fpdf) {
	w, h := pdf.GetPageSize()
	pdf.SetFont("courier", "", 12)
	pdf.SetTextColor(80, 80, 80)
	pdf.SetDrawColor(200, 200, 200)
	for x := 0.0; x < w; x = x + (w / 20.0) {
		pdf.Line(x, 0, x, h)
		_, lineHt := pdf.GetFontSize()
		pdf.Text(x, lineHt, fmt.Sprintf("%d", int(x)))
	}
	for y := 0.0; y < h; y = y + (w / 20.0) {
		pdf.Line(0, y, w, y)
		pdf.Text(0, y, fmt.Sprintf("%d", int(y)))
	}
}
