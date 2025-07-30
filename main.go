package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/runsong2018/tools-screen-shot/canvasx"
	"image/color"
)

func main() {
	a := app.New()
	window := a.NewWindow("Screen Shot")

	window.SetFullScreen(true)
	selectionRect := canvasx.NewScreenShot(color.Transparent, window.Canvas().Size())
	content := container.NewStack(
		widget.NewLabel("Drag the mouse to select the screenshot area"),
		selectionRect,
	)

	window.SetContent(content)

	window.ShowAndRun()
}
