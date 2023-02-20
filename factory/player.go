package factory

import (
	"image/color"
	"log"

	"bacteria/collision"
	"bacteria/component"
	"bacteria/layer"
	"bacteria/tag"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func NewPlayer(ecs *ecs.ECS, screenWidth, screenHeight float64) *donburi.Entry {
	entry := ecs.World.Entry(ecs.Create(
		layer.Default,
		tag.Player,
		component.Phagocytosis,
		component.Health,
		component.CollideBox,
		component.Velocity,
		component.Sprite,
		component.Weapon,
		component.Inventory,
	))

	img, _, err := ebitenutil.NewImageFromFile("assets/player.png")
	if err != nil {
		log.Fatal(err)
	}

	playerHeight := 32.0
	playerWidth := 32.0

	sp := component.SpriteData{
		Color:  color.RGBA{R: 172, G: 57, B: 148, A: 255},
		Height: playerHeight,
		Width:  playerWidth,
		Image:  img,
	}

	component.Sprite.SetValue(entry, sp)

	hd := component.HealthData{
		Health: 100,
	}
	component.Health.SetValue(entry, hd)

	collisionObj := resolv.NewObject(screenWidth/2, screenHeight/2, playerHeight, playerWidth, "player")
	collision.SetObject(entry, collisionObj)

	component.Velocity.Set(entry, component.NewVelocity(5, 200, 300))
	component.Phagocytosis.Set(entry, component.NewPhagocytosisData(10))
	component.Weapon.Set(entry, component.NewWeaponData(16, 40, 40))

	return entry
}
