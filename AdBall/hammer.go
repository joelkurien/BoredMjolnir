package main

import (
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Hammer struct {
	pivotX   float64
	pivotY   float64
	posFixed bool
	drag     bool
	angVel   float64
}

func Mjolnir() *Hammer {
	return &Hammer{
		pivotX:   0,
		pivotY:   0,
		posFixed: false,
	}
}

// Handles MouseClicks
func (g *Game) HammerAnimation(opts *ebiten.DrawImageOptions, image *ebiten.Image) *ebiten.DrawImageOptions {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		var x, y int
		g.hammer.drag = true
		if !g.hammer.posFixed {
			x, y = ebiten.CursorPosition()
			g.hammer.pivotX, g.hammer.pivotY = float64(x), float64(y)
			g.useHammer = true
			g.hammer.posFixed = true
		}
	} else if g.hammer.drag {
		g.hammer.drag = false
	}

	scaleFactor := 0.2
	opts.GeoM.Scale(scaleFactor, scaleFactor)
	opts.GeoM.Rotate(g.angle)
	opts.GeoM.Translate(float64(g.hammer.pivotX), float64(g.hammer.pivotY))
	return opts
}

func (g *Game) HammerSwings() {
	if g.hammer.posFixed && g.hammer.drag {
		cx, cy := ebiten.CursorPosition()
		cx1, cy1 := float64(cx), float64(cy)
		lb := g.hammer.pivotX + float64(hammerImage.Bounds().Max.X)
		rb := g.hammer.pivotY + float64(hammerImage.Bounds().Max.Y)
		if !((cx1 >= g.hammer.pivotX && cx1 <= lb) && (cy1 >= g.hammer.pivotY && cy1 <= rb) && g.angle == 0) {
			dx := cx1 - g.hammer.pivotX
			dy := cy1 - g.hammer.pivotY
			g.angle = math.Atan2(dy, dx)
			g.hammer.angVel = 0
		}
	}

	if !g.hammer.drag {
		if g.angle > 0.1 || g.angle < -0.1 {
			accel := -gravity * math.Sin(g.angle)
			g.hammer.angVel += accel * 0.1
			g.angle += g.hammer.angVel
		} else {
			g.hammer.angVel = 0
			g.angle = 0
		}
	}
	if !g.hammer.drag && g.angle == 0 {
		g.hammer.posFixed = false
		time.Sleep(250 * time.Millisecond)
		g.useHammer = false
	}
}
