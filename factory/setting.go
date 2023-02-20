package factory

import (
	"bacteria/component"
	"bacteria/layer"
	"bacteria/tag"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func NewSetting(ecs *ecs.ECS, screenWidth, screenHeight float64, maxMobs int) *donburi.Entry {
	entry := ecs.World.Entry(ecs.Create(
		layer.Default,
		tag.Setting,
		component.Setting,
	))

	component.Setting.SetValue(entry, component.SettingData{
		MapWidth:  screenWidth,
		MapHeight: screenHeight,
		MaxMobs:   maxMobs,
	})

	return entry
}
