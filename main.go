package main

import (
	"log"

	"bacteria/scene"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {
	scene *scene.Battle
}

func (g *Game) Update() error {
	g.scene.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.scene.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Bacteria")
	if err := ebiten.RunGame(&Game{
		scene: scene.NewBattle(screenWidth, screenHeight),
	}); err != nil {
		log.Fatal(err)
	}
}
