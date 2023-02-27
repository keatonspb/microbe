package main

import (
	"embed"
	"io"
	"log"

	"bacteria/helper"
	"bacteria/helper/sound"
	"bacteria/helper/storage"
	"bacteria/scene"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

//go:embed assets
var gameAssets embed.FS

const (
	screenWidth  = 1280
	screenHeight = 720
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

	fs := storage.NewStorage(func(path string) io.ReadCloser {
		f, err := gameAssets.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		return f
	})
	audioPlayer, err := sound.NewPlayer(audio.NewContext(44100), 44100)
	if err != nil {
		log.Fatal(err)
	}

	gameContext := helper.NewContext()

	gameContext.Audio = audioPlayer
	gameContext.Storage = fs
	gameContext.SetScreenSize(screenWidth, screenHeight)

	battleScene, err := scene.NewBattle(gameContext)
	if err != nil {
		log.Fatal(err)
	}

	if err := ebiten.RunGame(&Game{
		scene: battleScene,
	}); err != nil {
		log.Fatal(err)
	}
}
