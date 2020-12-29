package main

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

// Pos is a Point.
type Pos struct {
	X float64
	Y float64
}

type pixball struct {
	Pos
	V    pixel.Vec
	Size float64
}

type bar struct {
	Height  float64
	X       float64
	BottomY float64
}

var (
	myBar bar
	enBar bar
	ball  pixball
)

func genBar(startX, Y float64) *imdraw.IMDraw {
	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(1, 1, 1)
	imd.Push(pixel.V(startX, Y))
	imd.Color = pixel.RGB(1, 1, 1)
	imd.Push(pixel.V(startX+10, Y+barHeight))
	imd.Rectangle(0)
	return imd
}

func genBall(ball *pixball) *imdraw.IMDraw {
	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(1, 1, 1)
	imd.Push(pixel.V(*ball.Pos.X, *ball.Pos.Y))
	imd.Color = pixel.RGB(1, 1, 1)
	imd.Push(pixel.V(*ball.Pos.X+ball.Size, *ball.Pos.Y+ball.Size))
	imd.Rectangle(0)
	return imd
}

func moveBall(ball pixball) {
	*ball.Pos.X = *ball.Pos.X + ball.V.X
	*ball.Pos.Y = *ball.Pos.Y + ball.V.Y
}

func reflectBar(b bar, ball pixball) (float64, float64) {
	// x hit
	if (b.X < ball.Pos.X && ball.Pos.X < b.X+ball.Size) || (ball.Pos.X < barball.Pos.X && ball.Pos.X+ball.Size < barball.Pos.X) {
		// y hit
		if barball.Pos.Y < ball.Pos.Y && ball.Pos.Y < barball.Pos.Y+barHeight {
			fmt.Println("X,Y hit!")
			ball.V.X *= -1
			// ball.V.Y *= -1
		}
	}
	return ball.V.X, ball.V.Y
}

func isPointed(X, min, max float64) bool {
	if X <= min+(ball.Size/2) || max <= X+ball.Size {
		return true
	}
	return false
}

func run() {
	// making window and configuration
	cfg := pixelgl.WindowConfig{
		Title:  "pong",
		Bounds: pixel.R(0, 0, 500, 250),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// ball position X,Y / velocity ball.V.X,ball.V.Y
	ball.Pos.X, ball.Pos.Y = win.Bounds().Center().X, win.Bounds().Center().Y
	ball.V = pixel.V(-3.0, 3.0)
	// ball.Pos.X := win.Bounds().Center().X
	// ball.Pos.Y := win.Bounds().Center().Y
	// ball.V.X := -3.0
	// ball.V.Y := 3.0

	// my bar
	barHeight := 10.0
	myBar.Height = barHeight

	for !win.Closed() {
		if win.Pressed(pixelgl.KeyUp) {
			if !(myBar.BottomY+barHeight > win.Bounds().Max.Y) {
				myBar.BottomY += 5
			}
		}
		if win.Pressed(pixelgl.KeyDown) {
			if !(myBar.BottomY < win.Bounds().Min.Y) {
				myBar.BottomY -= 5
			}
		}

		win.Clear(colornames.Black)
		myBar := genBar(win.Bounds().Min.X+30, myBar.BottomY)
		myBar.Draw(win)

		// enBar := genBar(win.Bounds().Max.X-20, win.Bounds().Max.Y/4)
		// enBar.Draw(win)

		if ball.Pos.X+ball.Size+ball.V.X > win.Bounds().Max.X || ball.Pos.X+ball.V.X < win.Bounds().Min.X {
			ball.V.X *= -1
		}
		if ball.Pos.Y+ball.Size+ball.V.Y > win.Bounds().Max.Y || ball.Pos.Y+ball.V.Y < win.Bounds().Min.Y {
			ball.V.Y *= -1
		}
		moveBall(ball)
		ball.V.X, ball.V.Y = reflectBar(win.Bounds().Min.X+30, myBar.BottomY, ball)

		if isPointed(ball.Pos.X, win.Bounds().Min.X, win.Bounds().Max.X) {
			if ball.Pos.X < win.Bounds().Center().X {
				fmt.Println("right user point!")
			} else {
				fmt.Println("left user point!")
			}
			ball.V.X = -3.0
			ball.V.Y = 3.0
			ball.Pos.X = win.Bounds().Center().X
			ball.Pos.Y = win.Bounds().Center().Y
		}

		ball := genBall(&ball.Pos.X, &ball.Pos.Y)
		ball.Draw(win)

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
