package factory

import (
	"bacteria/component"
	"bacteria/tag"

	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/examples/platformer/layers"
)

func CreateSpace(ecs *ecs.ECS, screenWidth, screenHeight float64) *donburi.Entry {
	entry := ecs.World.Entry(ecs.Create(
		layers.Default,
		tag.Space,
		component.Space,
	))

	spaceData := resolv.NewSpace(int(screenWidth), int(screenHeight), 16, 16)
	component.Space.Set(entry, spaceData)

	return entry
}
