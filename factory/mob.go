package factory

import (
	"math/rand"

	"bacteria/collision"
	"bacteria/component"
	"bacteria/layer"
	"bacteria/meta"
	"bacteria/tag"

	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func NewMob(ecs *ecs.ECS, screenWidth float64, mobType int) *donburi.Entry {
	entry := ecs.World.Entry(ecs.Create(
		layer.Default,
		tag.Mob,
		component.Mob,
		component.CollideBox,
		component.Sprite,
		component.Float,
	))
	mobDesc := meta.Mobs[mobType]
	component.Float.Set(entry, &component.FloatData{
		Speed: mobDesc.Speed,
	})

	md := component.MobData{
		Type:   mobType,
		Attack: false,
	}
	component.Mob.SetValue(entry, md)

	collisionObj := resolv.NewObject(float64(rand.Int63n(int64(screenWidth))), 0, 5, 5, "mob")
	collision.SetObject(entry, collisionObj)

	return entry
}
