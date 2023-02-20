package system

import (
	"bacteria/collision"
	"bacteria/component"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type FloatController struct {
	screenWidth  float64
	screenHeight float64
	space        *donburi.Entry
}

func NewFloatController(screenWidth, screenHeight float64, space *donburi.Entry) *FloatController {
	return &FloatController{screenWidth: screenWidth, screenHeight: screenHeight, space: space}
}

func (w *FloatController) Update(ecs *ecs.ECS) {
	query.NewQuery(filter.Contains(component.Float, component.CollideBox)).Each(ecs.World, func(entry *donburi.Entry) {
		floatData := component.Float.Get(entry)
		collisionData := component.CollideBox.Get(entry)

		dx, dy := floatData.Move(ecs.Time.DeltaTime())

		collisionData.X = collisionData.X + dx
		collisionData.Y = collisionData.Y + dy

		if collisionData.Y > w.screenHeight || collisionData.X > w.screenWidth {
			collision.Remove(w.space, entry)
			ecs.World.Remove(entry.Entity())
			return
		}

		collisionData.Update()
	})
}
