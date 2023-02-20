package factory

import (
	"image/color"

	"bacteria/collision"
	"bacteria/component"
	"bacteria/layer"
	"bacteria/meta"
	"bacteria/tag"

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

	component.Sprite.Set(entry, &component.SpriteData{
		Width:  30,
		Height: 30,
		Color: color.RGBA{
			R: 217,
			G: 230,
			B: 125,
			A: 255,
		},
	})

	collision.SetObject(entry, resolv.NewObject(position.X, position.Y, 30, 30, "cell"))

	return entry
}
