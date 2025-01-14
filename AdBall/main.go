package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	gravity float64 = 0.5
)

var (
	hammerImage *ebiten.Image
	ballImage   *ebiten.Image
)

func returnImage(filepath string) *ebiten.Image {
	var err error
	var img *ebiten.Image
	img, _, err = ebitenutil.NewImageFromFile(filepath)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return img
}

func init() {
	hammerImage = returnImage("C:/Users/susan/OneDrive/Pictures/AdBallHammer.png")
	ballImage = returnImage("C:/Users/susan/OneDrive/Pictures/AdBallBall.png")
}

type Game struct {
	angle     float64
	useHammer bool
	hammer    *Hammer
}

// Handles error situation
func (g *Game) Update() error {
	g.HammerSwings()
	return nil
}

// This displays content
func (g *Game) Draw(screen *ebiten.Image) {
	hop := &ebiten.DrawImageOptions{}
	hop = g.HammerAnimation(hop, hammerImage)
	if g.useHammer {
		screen.DrawImage(hammerImage, hop)
	}
	//screen.DrawImage(ballImage, bop)
}

// sets the canvas size of the game
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 720, 720
}

// Runs the entire system.
func main() {
	ebiten.SetWindowSize(720, 720)
	ebiten.SetWindowTitle("AdBall")

	game := &Game{
		angle:     0,
		useHammer: false,
	}
	game.hammer = Mjolnir()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
