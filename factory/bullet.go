package factory

import (
	"image/color"
	"time"

	"bacteria/collision"
	"bacteria/component"
	"bacteria/layer"
	"bacteria/meta"
	"bacteria/tag"

	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func NewBullet(ecs *ecs.ECS, vector meta.Vector, position meta.Point) *donburi.Entry {
	entry := ecs.World.Entry(ecs.Create(
		layer.Default,
		tag.Bullet,
		component.CollideBox,
		component.Sprite,
		component.Move,
		component.Velocity,
	))

	collisionObj := resolv.NewObject(position.X, position.Y, 2, 2, "bullet")
	collision.SetObject(entry, collisionObj)

	component.Sprite.Set(entry, &component.SpriteData{
		Color: color.RGBA{
			R: 255,
			G: 255,
			B: 255,
			A: 255,
		},
		Shape:  meta.NewShape(position.X, position.Y, 4, 4),
		Height: 2,
		Width:  2,
	})

	component.Move.Set(entry, &component.MoveData{
		Direction: vector,
		LiveTill:  time.Now().Add(time.Second),
	})

	component.Velocity.Set(entry, &component.VelocityData{
		InitialSpeed: 10,
		MaxSpeed:     100,
		Acceleration: 10,
	})

	return entry
}
