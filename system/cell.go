package system

import (
	"image/color"

	"bacteria/collision"
	"bacteria/component"
	"bacteria/factory"
	"bacteria/helper"
	"bacteria/meta"
	"bacteria/tag"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type CellController struct {
	gameCtx   *helper.Context
	space     *donburi.Entry
	maxCells  int
	inventory *component.InventoryData
}

func NewCellController(ctx *helper.Context,
	space *donburi.Entry,
	maxCells int,
	inventory *component.InventoryData,
) *CellController {
	return &CellController{
		gameCtx:   ctx,
		space:     space,
		maxCells:  maxCells,
		inventory: inventory,
	}
}

func (w *CellController) Update(ecs *ecs.ECS) {
	w.generate(ecs)
	w.checkHealth(ecs)
}

func (w *CellController) checkHealth(ecs *ecs.ECS) {
	tag.Cell.Each(ecs.World, func(entry *donburi.Entry) {
		health := component.Health.Get(entry)
		if health.Health <= 0 {

			consume := component.Consumable.Get(entry)
			w.inventory.Cells += consume.Amount
			collision.Remove(w.space, entry)
			ecs.World.Remove(entry.Entity())
		}
	})
}

func (w *CellController) generate(ecs *ecs.ECS) {
	if query.NewQuery(filter.Contains(component.Consumable)).Count(ecs.World) >= w.maxCells {
		return
	}

	cell := factory.NewCell(ecs, meta.Point{
		X: helper.RandFloat(w.gameCtx.ScreenWidth()),
		Y: 0,
	}, w.gameCtx.Storage)

	collision.AddToSpace(w.space, cell)
}

func (w *CellController) Draw(ecs *ecs.ECS, screen *ebiten.Image) {
	tag.Cell.Each(ecs.World, func(entry *donburi.Entry) {
		spriteData := component.Sprite.Get(entry)
		collisionData := component.CollideBox.Get(entry)
		hp := component.Health.Get(entry)

		sprite := component.Sprite.Get(entry)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(collisionData.X, collisionData.Y)

		screen.DrawImage(sprite.GetCell(0, 0), op)

		healthBarWidth := spriteData.GetCellWidth() * (hp.Health / hp.MaxHp)
		ebitenutil.DrawRect(screen, collisionData.X, collisionData.Y+spriteData.GetCellHeight()+3, healthBarWidth, 3,
			color.RGBA{255, 0, 0, 255})

	})
}
