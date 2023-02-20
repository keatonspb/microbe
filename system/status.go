package system

import (
	"fmt"

	"bacteria/component"
	"bacteria/tag"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func DrawStatus(ecs *ecs.ECS, screen *ebiten.Image) {
	p, ok := tag.Player.First(ecs.World)
	if !ok {
		return
	}

	health := component.Health.Get(p)
	ph := component.Phagocytosis.Get(p)
	inv := component.Inventory.Get(p)

	enemies := 0
	tag.Mob.Each(ecs.World, func(entry *donburi.Entry) {
		enemies++
	})

	ebitenutil.DebugPrint(screen, fmt.Sprintf("Health: %0.0f\nAttackers: %d\nEnemies: %d\nCells collected: %d", health.Health, len(ph.Attackers), enemies, inv.Cells))
}
