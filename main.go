package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

type WindowConfig struct {
	width, height int
	title         string
}

var config WindowConfig

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.width, config.height
}

func main() {
	config = WindowConfig{
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
