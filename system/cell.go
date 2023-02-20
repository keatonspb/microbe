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
	screenWidth  float64
	screenHeight float64
	space        *donburi.Entry
	maxCells     int
	inventory    *component.InventoryData
}

func NewCellController(screenWidth, screenHeight float64, space *donburi.Entry, maxCells int, inventory *component.InventoryData) *CellController {
	return &CellController{screenWidth: screenWidth, screenHeight: screenHeight, space: space, maxCells: maxCells, inventory: inventory}
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
		X: helper.RandFloat(w.screenWidth),
		Y: 0,
	})

	collision.AddToSpace(w.space, cell)
}

func (w *CellController) Draw(ecs *ecs.ECS, screen *ebiten.Image) {
	tag.Cell.Each(ecs.World, func(entry *donburi.Entry) {
		spriteData := component.Sprite.Get(entry)
		collisionData := component.CollideBox.Get(entry)
		hp := component.Health.Get(entry)

		ebitenutil.DrawRect(screen, collisionData.X, collisionData.Y, spriteData.Height, spriteData.Width, spriteData.Color)

		helthBarWidth := spriteData.Width * (hp.Health / hp.MaxHp)
		ebitenutil.DrawRect(screen, collisionData.X, collisionData.Y+spriteData.Height+3, helthBarWidth, 5,
			color.RGBA{255, 0, 0, 255})

	})
}
