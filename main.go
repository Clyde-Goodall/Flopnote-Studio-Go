package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"log"
)

type GameConfig struct {
	width, height int
	title         string
}

var config GameConfig

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(
		screen,
		0,
		0,
		float32(config.width),
		float32(config.height),
		color.Color(
			color.RGBA{
				R: 255,
				G: 255,
				B: 255,
				A: 1,
			},
		),
		true,
	)
	vector.DrawFilledCircle(
		screen,
		0,
		0,
		float32(300),
		color.Color(
			color.RGBA{
				R: 1,
				G: 1,
				B: 1,
				A: 255,
			},
		),
		true,
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.width, config.height
}

func main() {
	config = GameConfig{
		width:  1000,
		height: 600,
		title:  "Shrimp Mode",
	}
	// setup
	ebiten.SetWindowSize(config.width, config.height)
	ebiten.SetWindowTitle(config.title)

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}

}
