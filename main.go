package main

import (
	"fmt"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Line struct {
	x1, y1  int
	length  int
	radians float64
	c       color.Color
}

type game struct {
	l *Line
}

func (g *game) Layout(outWidth, outHeight int) (w, h int) { return screenWidth, screenHeight }
func (g *game) Update() error {
	g.l.radians += math.Pi / 180
	if g.l.radians > math.Pi/2 {
		g.l.radians = 0
	}
	return nil
}
func (g *game) Draw(screen *ebiten.Image) {
	x := float64(g.l.length) * math.Cos(g.l.radians)
	y := float64(g.l.length) * math.Sin(g.l.radians)
	x2, y2 := g.l.x1+int(x), g.l.y1+int(y)
	DrawLine(screen, g.l.x1, g.l.y1, x2, y2, g.l.c)
	ebitenutil.DebugPrint(screen, fmt.Sprint(d))
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	g := game{&Line{screenWidth / 2, screenHeight / 2, 100, 0, color.RGBA{1, 100, 100, 255}}}
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}

var d float64

func DrawLine(img *ebiten.Image, x1, y1, x2, y2 int, c color.Color) {

	img.Set(x1, y1, color.RGBA{1, 255, 1, 255})
	img.Set(x2, y2, color.RGBA{1, 255, 1, 255})

	// abs(Dy) < abs(dx) | / abs(dx) => abs(Dy)/abs(Dx) < 1 == abs(k) < 1
	if math.Abs(float64(y2-y1)) < math.Abs(float64(x2-x1)) {
		if x1 > x2 {
			x1, x2 = x2, x1
			y1, y2 = y2, y1
		}
		Dx, Dy := x2-x1, y2-y1
		dirY := 1
		midY := 0.5
		if Dy < 0 {
			dirY = -1
			midY = -0.5
		}
		f := func(x, y float64) float64 {
			A, B, C := Dy, -Dx, Dx*y1-Dy*x1
			return float64(A)*x + float64(B)*y + float64(C)
		}
		for x, y := x1, y1; x < x2; x++ {
			img.Set(x, y, c)
			xm, ym := float64(x)+1, float64(y)+midY
			d = f(xm, ym) * float64(dirY)
			if d > 0 {
				y += dirY
			}
		}
	} else {
		if y1 > y2 {
			x1, x2 = x2, x1
			y1, y2 = y2, y1
		}
		Dx, Dy := x2-x1, y2-y1
		dirX := 1
		midX := 0.5
		if Dx < 0 {
			dirX = -1
			midX = -0.5
		}
		f := func(x, y float64) float64 {
			A, B, C := Dy, -Dx, Dx*y1-Dy*x1
			return float64(A)*x + float64(B)*y + float64(C)
		}
		for x, y := x1, y1; y < y2; y++ {
			img.Set(x, y, c)
			xm, ym := float64(x)+midX, float64(y)+1
			d = f(ym, xm) * float64(dirX)
			if d > 0 {
				x += dirX
			}
		}

	}
}
