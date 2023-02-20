package main

import (
	"embed"
	"log"

	"bacteria/scene"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed assets
var gameAssets embed.FS

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

func (g *Game) Layout(_, _ int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Bacteria")
	if err := ebiten.RunGame(&Game{
		scene: scene.NewBattle(screenWidth, screenHeight, &gameAssets),
	}); err != nil {
		log.Fatal(err)
	}
}
