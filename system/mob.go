package system

import (
	"fmt"

	"bacteria/assets/audio"
	"bacteria/collision"
	"bacteria/component"
	"bacteria/factory"
	"bacteria/helper"
	"bacteria/meta"
	"bacteria/tag"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sirupsen/logrus"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

type MobController struct {
	gameCtx *helper.Context
}

func NewMobController(ctx *helper.Context) (*MobController, error) {
	err := ctx.Audio.AddSoundFromPath(audio.NeitrofillAttack, "assets/audio/neitrofil_attack.wav")
	if err != nil {
		return nil, fmt.Errorf("failed to create sound player: %w", err)
	}
	return &MobController{
		gameCtx: ctx,
	}, nil
}

func (c *MobController) Update(ecs *ecs.ECS) {
	c.generateMob(ecs)
	c.collideMob(ecs)
}

func (c *MobController) collideMob(ecs *ecs.ECS) {
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
				c.gameCtx.Audio.PlaySound(audio.NeitrofillAttack)
			})
		}
	})
}

func (c *MobController) generateMob(ecs *ecs.ECS) {
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
		en := factory.NewMob(ecs, settings.MapWidth, meta.RandMobType(), c.gameCtx.Storage)
		collision.AddToSpace(spaceEntry, en)
	}

}

func (c *MobController) Draw(ecs *ecs.ECS, screen *ebiten.Image) {
	c.drawMob(ecs, screen)
}

func (c *MobController) drawMob(ecs *ecs.ECS, screen *ebiten.Image) {
	tag.Mob.Each(ecs.World, func(e *donburi.Entry) {
		object := component.CollideBox.Get(e)
		sprite := component.Sprite.Get(e)
		mob := component.Mob.Get(e)
		if mob.Attack {
			return
		}
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(object.X, object.Y)
		screen.DrawImage(sprite.GetImage(), op)
	})
}
