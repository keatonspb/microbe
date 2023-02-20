package system

import (
	"bacteria/collision"
	"bacteria/component"
	"bacteria/factory"
	"bacteria/meta"
	"bacteria/tag"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sirupsen/logrus"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

func CollideMob(ecs *ecs.ECS) {
	settingEntry, ok := tag.Setting.First(ecs.World)
	if !ok {
		logrus.Error("setting not found")
		return
	}
	settings := component.Setting.Get(settingEntry)
	if settings.Stop {
		return
	}

	tag.Mob.Each(ecs.World, func(e *donburi.Entry) {
		mob := component.Mob.Get(e)
		if mob.Attack {
			return
		}

		object := collision.GetObject(e)

		if coll := object.Check(0, 0, "player"); coll != nil {
			target := coll.Objects[0]

			donburi.NewQuery(
				filter.Contains(component.CollideBox, component.Health)).Each(ecs.World, func(p *donburi.Entry) {
				o := component.CollideBox.Get(p)
				if o == target {
					ph := component.Phagocytosis.Get(p)
					ph.AddAttacker(mob)
					mob.Attack = true
				}
			})
		}
	})
}

func GenerateMob(ecs *ecs.ECS) {
	settingEntry, ok := tag.Setting.First(ecs.World)
	if !ok {
		logrus.Error("setting not found")
		return
	}

	spaceEntry, ok := tag.Space.First(ecs.World)
	if !ok {
		logrus.Error("space not found")
		return
	}

	settings := component.Setting.Get(settingEntry)

	query := donburi.NewQuery(filter.Contains(tag.Mob))

	if query.Count(ecs.World) < settings.MaxMobs {
		en := factory.NewMob(ecs, settings.MapWidth, meta.RandMobType())
		collision.AddToSpace(spaceEntry, en)
	}

}

func DrawMob(ecs *ecs.ECS, screen *ebiten.Image) {
	tag.Mob.Each(ecs.World, func(e *donburi.Entry) {
		object := component.CollideBox.Get(e)
		sprite := component.Sprite.Get(e)
		mob := component.Mob.Get(e)
		if mob.Attack {
			return
		}
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(object.X, object.Y)
		screen.DrawImage(sprite.Image, op)
	})
}
