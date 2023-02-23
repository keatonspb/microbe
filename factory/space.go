package factory

import (
	"bacteria/component"
	"bacteria/helper"
	"bacteria/tag"

	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/examples/platformer/layers"
)

func CreateSpace(ctx *helper.Context, ecs *ecs.ECS) *donburi.Entry {
	entry := ecs.World.Entry(ecs.Create(
		layers.Default,
		tag.Space,
		component.Space,
	))

	spaceData := resolv.NewSpace(int(ctx.ScreenWidth()), int(ctx.ScreenHeight()), 16, 16)
	component.Space.Set(entry, spaceData)

	return entry
}
