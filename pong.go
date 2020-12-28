package main

import (
	"fmt"
	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var (
	barHeight float64 = 100.0
	ballSize  float64 = 10.0
	myBarPos  float64
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

func genBall(posX, posY *float64) *imdraw.IMDraw {
	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(1, 1, 1)
	imd.Push(pixel.V(*posX, *posY))
	imd.Color = pixel.RGB(1, 1, 1)
	imd.Push(pixel.V(*posX+ballSize, *posY+ballSize))
	imd.Rectangle(0)
	return imd
}

func moveBall(posX, posY *float64, vX, vY float64) {
	*posX = *posX + vX
	*posY = *posY + vY
}

func reflectBar(barPosX, barPosY, posX, posY, vX, vY float64) (float64, float64) {
	// x hit
	if (barPosX < posX && posX < barPosX+ballSize) || (posX < barPosX && posX+ballSize < barPosX) {
		// y hit
		if barPosY < posY && posY < barPosY+barHeight {
			fmt.Println("X,Y hit!")
			vX *= -1
			// vY *= -1
		}
	}
	return vX, vY
}

func isPointed(X, min, max float64) bool {
	if X <= min+(ballSize/2) || max <= X+ballSize {
		return true
	}
	return false
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixels!",
		Bounds: pixel.R(0, 0, 500, 250),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	posX := win.Bounds().Center().X
	posY := win.Bounds().Center().Y
	vX := -3.0
	vY := 3.0

	for !win.Closed() {
		if win.Pressed(pixelgl.KeyUp) {
			if !(myBarPos+barHeight > win.Bounds().Max.Y) {
				myBarPos += 5
			}
		}
		if win.Pressed(pixelgl.KeyDown) {
			if !(myBarPos < win.Bounds().Min.Y) {
				myBarPos -= 5
			}
		}

		win.Clear(colornames.Black)
		myBar := genBar(win.Bounds().Min.X+30, myBarPos)
		myBar.Draw(win)

		// enBar := genBar(win.Bounds().Max.X-20, win.Bounds().Max.Y/4)
		// enBar.Draw(win)

		if posX+ballSize+vX > win.Bounds().Max.X || posX+vX < win.Bounds().Min.X {
			vX *= -1
		}
		if posY+ballSize+vY > win.Bounds().Max.Y || posY+vY < win.Bounds().Min.Y {
			vY *= -1
		}
		moveBall(&posX, &posY, vX, vY)
		vX, vY = reflectBar(win.Bounds().Min.X+30, myBarPos, posX, posY, vX, vY)

		if isPointed(posX, win.Bounds().Min.X, win.Bounds().Max.X) {
			if posX < win.Bounds().Center().X {
				fmt.Println("right user point!")
			} else {
				fmt.Println("left user point!")
			}
			vX = -3.0
			vY = 3.0
			posX = win.Bounds().Center().X
			posY = win.Bounds().Center().Y
		}

		ball := genBall(&posX, &posY)
		ball.Draw(win)

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
