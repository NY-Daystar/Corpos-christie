package progress

// package main

// import (
// 	"image/color"
// 	"time"

// 	"fyne.io/fyne/v2"
// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/canvas"
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/widget"
// )

// type CircularProgress struct {
// 	widget.BaseWidget
// 	Progress float64 // between 0.0 and 1.0
// }

// func NewCircularProgress() *CircularProgress {
// 	c := &CircularProgress{}
// 	c.ExtendBaseWidget(c)
// 	return c
// }

// func (c *CircularProgress) CreateRenderer() fyne.WidgetRenderer {
// 	bg := canvas.NewCircle(color.Gray{Y: 0x99})
// 	fg := canvas.NewCircle(color.RGBA{R: 0x00, G: 0x96, B: 0x88, A: 0xff})
// 	fg.StrokeColor = color.RGBA{R: 0x00, G: 0x96, B: 0x88, A: 0xff}
// 	fg.StrokeWidth = 10

// 	return &circularProgressRenderer{c: c, bg: bg, fg: fg}
// }

// type circularProgressRenderer struct {
// 	c  *CircularProgress
// 	bg *canvas.Circle
// 	fg *canvas.Circle
// }

// func (r *circularProgressRenderer) Layout(size fyne.Size) {
// 	r.bg.Resize(size)
// 	r.bg.Move(fyne.NewPos(0, 0))

// 	r.fg.Resize(size)
// 	r.fg.Move(fyne.NewPos(0, 0))
// }

// func (r *circularProgressRenderer) MinSize() fyne.Size {
// 	return fyne.NewSize(100, 100)
// }

// func (r *circularProgressRenderer) Refresh() {
// 	r.fg.FillColor = color.Transparent // Clear the fill color to make the arc visible
// 	r.fg.Refresh()

// 	r.fg.StrokeWidth = 10
// 	r.fg.StrokeColor = color.RGBA{R: 0x00, G: 0x96, B: 0x88, A: 0xff}
// 	r.fg.Refresh()

// 	// Create a new arc for progress
// 	arc := canvas.NewCircle(color.Black)
// 	// angle := r.c.Progress * 360.0
// 	//arc := canvas.NewArc(color.RGBA{R: 0x00, G: 0x96, B: 0x88, A: 0xff})
// 	// arc.StartAngle = 0
// 	// arc.SweepAngle = angle
// 	arc.Resize(r.fg.Size())
// 	arc.Move(r.fg.Position())

// 	arc.StrokeWidth = 10
// 	arc.StrokeColor = r.fg.StrokeColor

// 	r.c.Hide() // Hide the original circle
// 	r.fg = arc
// 	r.c.Show() // Show the updated arc
// 	r.fg.Refresh()
// }

// func (r *circularProgressRenderer) BackgroundColor() color.Color {
// 	return color.Transparent
// }

// func (r *circularProgressRenderer) Objects() []fyne.CanvasObject {
// 	return []fyne.CanvasObject{r.bg, r.fg}
// }

// func (r *circularProgressRenderer) Destroy() {}

// TODO a faire fonctionner
// func main() {
// 	a := app.New()
// 	w := a.NewWindow("Circular Progress Bar")

// 	progress := NewCircularProgress()
// 	progress.Progress = 0.75 // 75% progress

// 	go func() {
// 		for i := 0; i <= 100; i++ {
// 			time.Sleep(50 * time.Millisecond)
// 			progress.Progress = float64(i) / 100
// 			progress.Refresh()
// 		}
// 	}()

// 	w.SetContent(container.NewVBox(progress))

// 	w.Resize(fyne.NewSize(200, 200))
// 	w.ShowAndRun()
// }
