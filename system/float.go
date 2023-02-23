package system

import (
	"bacteria/collision"
	"bacteria/component"
	"bacteria/helper"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type FloatController struct {
	gameCtx *helper.Context
	space   *donburi.Entry
}

func NewFloatController(ctx *helper.Context, space *donburi.Entry) *FloatController {
	return &FloatController{gameCtx: ctx, space: space}
}

func (w *FloatController) Update(ecs *ecs.ECS) {
	query.NewQuery(filter.Contains(component.Float, component.CollideBox)).Each(ecs.World, func(entry *donburi.Entry) {
		floatData := component.Float.Get(entry)
		collisionData := component.CollideBox.Get(entry)

		dx, dy := floatData.Move(ecs.Time.DeltaTime())

		collisionData.X = collisionData.X + dx
		collisionData.Y = collisionData.Y + dy

		if collisionData.Y > w.gameCtx.ScreenHeight() || collisionData.X > w.gameCtx.ScreenWidth() {
			collision.Remove(w.space, entry)
			ecs.World.Remove(entry.Entity())
			return
		}

		collisionData.Update()
	})
}
