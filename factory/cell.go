package factory

import (
	"log"

	"bacteria/collision"
	"bacteria/component"
	"bacteria/layer"
	"bacteria/meta"
	"bacteria/tag"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func NewCell(ecs *ecs.ECS, position meta.Point) *donburi.Entry {
	entry := ecs.World.Entry(ecs.Create(
		layer.Default,
		tag.Cell,
		component.CollideBox,
		component.Sprite,
		component.Health,
		component.Float,
		component.Consumable,
	))

	component.Consumable.Set(entry, &component.ConsumableData{
		Amount: 10,
	})

	component.Health.Set(entry, &component.HealthData{
		Health: 10,
		MaxHp:  10,
	})

	component.Float.Set(entry, &component.FloatData{
		Speed: 10,
	})

	img, _, err := ebitenutil.NewImageFromFile("assets/cell.png")
	if err != nil {
		log.Fatal(err)
	}

	component.Sprite.Set(entry, &component.SpriteData{
		Width:  32,
		Height: 32,
		Image:  img,
	})

	collision.SetObject(entry, resolv.NewObject(position.X, position.Y, 30, 30, "cell"))

	return entry
}
