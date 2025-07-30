package canvasx

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"image"
	"image/color"
)

type ScreenShot struct {
	Rectangle        *canvas.Rectangle
	startPos, endPos fyne.Position
	isStarted        bool
}

func (s *ScreenShot) MinSize() fyne.Size {
	return s.Rectangle.MinSize()
}

func (s *ScreenShot) Size() fyne.Size {
	return s.Rectangle.Size()
}

func (s *ScreenShot) Visible() bool {
	return s.Rectangle.Visible()
}

func (s *ScreenShot) Hide() {
	s.Rectangle.Hide()
}

func (s *ScreenShot) Show() {
	s.Rectangle.Show()
}

func (s *ScreenShot) Move(pos fyne.Position) {
	s.Rectangle.Move(pos)
}

func (s *ScreenShot) Refresh() {
	s.Rectangle.Refresh()
}

func (s *ScreenShot) Resize(size fyne.Size) {
	s.Rectangle.Resize(size)
}

func (s *ScreenShot) Position() fyne.Position {
	return s.Rectangle.Position()
}

func (s *ScreenShot) MouseIn(event *desktop.MouseEvent) {
	fmt.Println("MouseIn")
}

func (s *ScreenShot) MouseMoved(event *desktop.MouseEvent) {
	fmt.Println("MouseMoved")
	if s.isStarted {
		selectRect := canvas.NewRectangle(color.Transparent)
		selectRect.FillColor = color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x00}
		selectRect.StrokeColor = color.Black
		selectRect.StrokeWidth = 2
		selectRectWidth := s.endPos.X - s.startPos.X
		selectRectHeight := s.endPos.Y - s.startPos.Y
		selectRect.Resize(fyne.NewSize(selectRectWidth, selectRectHeight))
		selectRect.Move(s.startPos)
	}
}

func (s *ScreenShot) MouseOut() {
	fmt.Println("MouseOut")

}

func (s *ScreenShot) MouseDown(event *desktop.MouseEvent) {
	fmt.Println("MouseDown")
	if !s.isStarted {
		s.isStarted = true
		s.startPos = event.Position
		s.Rectangle.Hide()
	}
}

func (s *ScreenShot) MouseUp(event *desktop.MouseEvent) {
	fmt.Println("MouseUp")
	if s.isStarted {
		s.isStarted = false
		s.endPos = event.Position
	}
}

func (s *ScreenShot) Capture(callback func(img image.Image)) {

}

func NewScreenShot(color color.Color, size fyne.Size) *ScreenShot {
	s := &ScreenShot{
		Rectangle: canvas.NewRectangle(color),
	}
	s.Rectangle.Resize(size)
	s.Rectangle.Refresh()
	return s
}
