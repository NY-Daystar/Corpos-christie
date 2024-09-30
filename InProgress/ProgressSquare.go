package progress

// package main

// import (
// 	"fmt"
// 	"image/color"
// 	"time"

// 	"fyne.io/fyne/v2"
// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/canvas"
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/widget"
// )

// type SquareProgress struct {
// 	widget.BaseWidget
// 	Progress float64 // between 0.0 and 1.0
// 	Width    float32 // width of the square
// 	Height   float32 // height of the square
// }

// type squareProgressRenderer struct {
// 	s      *SquareProgress
// 	top    *canvas.Line
// 	right  *canvas.Line
// 	bottom *canvas.Line
// 	left   *canvas.Line
// }

// func NewSquareProgress() *SquareProgress {
// 	s := &SquareProgress{}
// 	s.ExtendBaseWidget(s)
// 	s.Width = 100  // Default width
// 	s.Height = 100 // Default height
// 	return s
// }

// func (s *SquareProgress) SetSize(width, height float32) {
// 	s.Width = width
// 	s.Height = height
// 	s.Refresh()
// }

// func (s *SquareProgress) CreateRenderer() fyne.WidgetRenderer {
// 	top := canvas.NewLine(color.RGBA{R: 0x00, G: 0x96, B: 0x88, A: 0xff})
// 	right := canvas.NewLine(color.RGBA{R: 0x00, G: 0x96, B: 0x88, A: 0xff})
// 	bottom := canvas.NewLine(color.RGBA{R: 0x00, G: 0x96, B: 0x88, A: 0xff})
// 	left := canvas.NewLine(color.RGBA{R: 0x00, G: 0x96, B: 0x88, A: 0xff})

// 	return &squareProgressRenderer{
// 		s:      s,
// 		top:    top,
// 		right:  right,
// 		bottom: bottom,
// 		left:   left,
// 	}
// }

// func (r *squareProgressRenderer) Layout(size fyne.Size) {
// 	// Calculating the points for each line based on the progress
// 	perimeter := 4 * r.s.Width
// 	progress := r.s.Progress * float64(perimeter)

// 	// Reset positions
// 	r.top.Position1 = fyne.NewPos(0, 0)
// 	r.top.Position2 = fyne.NewPos(0, 0)
// 	r.top.StrokeWidth = 5
// 	r.right.Position1 = fyne.NewPos(r.s.Width, 0)
// 	r.right.Position2 = fyne.NewPos(r.s.Width, 0)
// 	r.right.StrokeWidth = 5
// 	r.bottom.Position1 = fyne.NewPos(r.s.Width, r.s.Width)
// 	r.bottom.Position2 = fyne.NewPos(r.s.Width, r.s.Width)
// 	r.bottom.StrokeWidth = 5
// 	r.left.Position1 = fyne.NewPos(0, r.s.Width)
// 	r.left.Position2 = fyne.NewPos(0, r.s.Width)
// 	r.left.StrokeWidth = 5

// 	if progress <= float64(r.s.Width) {
// 		r.top.Position2 = fyne.NewPos(float32(progress), 0)
// 	} else if progress <= float64(2*r.s.Width) {
// 		r.top.Position2 = fyne.NewPos(r.s.Width, 0)
// 		r.right.Position2 = fyne.NewPos(r.s.Width, float32(progress)-r.s.Width)
// 	} else if progress <= float64(3*r.s.Width) {
// 		r.top.Position2 = fyne.NewPos(r.s.Width, 0)
// 		r.right.Position2 = fyne.NewPos(r.s.Width, r.s.Width)
// 		r.bottom.Position2 = fyne.NewPos(r.s.Width-(float32(progress)-2*r.s.Width), r.s.Width)
// 	} else {
// 		r.top.Position2 = fyne.NewPos(r.s.Width, 0)
// 		r.right.Position2 = fyne.NewPos(r.s.Width, r.s.Width)
// 		r.bottom.Position2 = fyne.NewPos(0, r.s.Width)
// 		r.left.Position2 = fyne.NewPos(0, r.s.Width-(float32(progress)-3*r.s.Width))
// 	}
// }

// func (r *squareProgressRenderer) MinSize() fyne.Size {
// 	return fyne.NewSize(float32(r.s.Width), float32(r.s.Height))
// }

// func (r *squareProgressRenderer) Refresh() {
// 	r.Layout(r.s.Size())
// 	r.top.Refresh()
// 	r.right.Refresh()
// 	r.bottom.Refresh()
// 	r.left.Refresh()
// }

// func (r *squareProgressRenderer) BackgroundColor() color.Color {
// 	return color.Transparent
// }

// func (r *squareProgressRenderer) Objects() []fyne.CanvasObject {
// 	return []fyne.CanvasObject{r.top, r.right, r.bottom, r.left}
// }

// func (r *squareProgressRenderer) Destroy() {}

// func (r *SquareProgress) Animate() {
// 	fmt.Printf("Lancement de l'animation\n")
// 	go func() {
// 		for {
// 			for i := 0; i <= 100; i++ {
// 				time.Sleep(20 * time.Millisecond) // Update interval
// 				r.Progress = float64(i) / 100
// 				r.Refresh()
// 			}
// 			time.Sleep(500 * time.Millisecond) // Optional pause before restarting
// 			fmt.Printf("Animation terminé\n")
// 		}
// 	}()
// }

// // TODO mettre en noir les lignes
// // TODO si le circular fonctionne pas on peut voir pour courber ces lignes

// func main() {
// 	a := app.New()
// 	w := a.NewWindow("Square Progress Bar")

// 	progress := NewSquareProgress() // TODO mettre la taille ici
// 	progress.Progress = 0.0         // Start at 0% progress
// 	progress.Animate()              // Start the animation

// 	w.SetContent(container.NewVBox(progress))

// 	//w.Resize(fyne.NewSize(50, 50))
// 	w.Resize(fyne.NewSize(1100, 540))
// 	w.ShowAndRun()
// }
