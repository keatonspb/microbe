package system

import (
	"fmt"
	"math"
	"time"

	"bacteria/assets/audio"
	"bacteria/collision"
	"bacteria/component"
	"bacteria/factory"
	"bacteria/helper"
	"bacteria/meta"
	"bacteria/tag"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/sirupsen/logrus"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

type WeaponController struct {
	gameCtx *helper.Context
	space   *donburi.Entry
}

func NewWeaponController(ctx *helper.Context, space *donburi.Entry) (*WeaponController, error) {
	c := &WeaponController{gameCtx: ctx, space: space}

	err := c.gameCtx.Audio.AddSoundFromPath(audio.Shoot, "assets/audio/shoot.wav")
	if err != nil {
		return nil, fmt.Errorf("failed to create sound player: %w", err)
	}

	err = c.gameCtx.Audio.AddSoundFromPath(audio.Hit, "assets/audio/hit.wav")
	if err != nil {
		return nil, fmt.Errorf("failed to create sound player: %w", err)
	}

	return c, nil
}

func (w *WeaponController) Update(ecs *ecs.ECS) {
	w.updateShoot(ecs)
	w.updateBulletMove(ecs)
	w.collide(ecs)
}

func (w *WeaponController) collide(ecs *ecs.ECS) {
	tag.Bullet.Each(ecs.World, func(e *donburi.Entry) {
		target := component.CollideBox.Get(e)

		if coll := target.Check(0, 0, "mob"); coll != nil {
			collision.Remove(w.space, e)
			ecs.World.Remove(e.Entity())
			object := coll.Objects[0]
			w.gameCtx.Audio.PlaySound(audio.Hit)
			tag.Mob.Each(ecs.World, func(e *donburi.Entry) {
				mobObj := component.CollideBox.Get(e)
				if mobObj == object {
					collision.Remove(w.space, e)
					ecs.World.Remove(e.Entity())
				}
			})
		}
		if coll := target.Check(0, 0, "cell"); coll != nil {
			collision.Remove(w.space, e)
			ecs.World.Remove(e.Entity())

			object := coll.Objects[0]
			w.gameCtx.Audio.PlaySound(audio.Hit)

			tag.Cell.Each(ecs.World, func(e *donburi.Entry) {
				obj := component.CollideBox.Get(e)
				if obj == object {
					health := component.Health.Get(e)
					health.Health -= 1
				}
			})
		}
	})
}

func (w *WeaponController) updateShoot(ecs *ecs.ECS) {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		player, ok := tag.Player.First(ecs.World)
		if !ok {
			logrus.Info("player not found")
			return
		}

		weapon := component.Weapon.Get(player)
		obj := component.CollideBox.Get(player)

		w.shoot(ecs, weapon, meta.NewPoint(obj.X+obj.W/2-1, obj.Y+obj.H/2-1))
	}
}

func (w *WeaponController) shoot(ecs *ecs.ECS, weapon *component.WeaponData, start meta.Point) {
	if weapon.NextShot.After(time.Now()) {
		logrus.Infof("reloading")
		return
	}

	vectors := getVector(weapon.Bullets, weapon.Speed)

	for _, v := range vectors {
		bullet := factory.NewBullet(ecs, v, start)
		collision.AddToSpace(w.space, bullet)
	}

	w.gameCtx.Audio.PlaySound(audio.Shoot)

	weapon.NextShot = time.Now().Add(time.Second)
}

func (w *WeaponController) updateBulletMove(ecs *ecs.ECS) {
	settingEntry, ok := tag.Setting.First(ecs.World)
	if !ok {
		logrus.Error("setting not found")
		return
	}
	settings := component.Setting.Get(settingEntry)

	spaceEntry, ok := tag.Space.First(ecs.World)
	if !ok {
		logrus.Error("space not found")
		return
	}

	tag.Bullet.Each(ecs.World, func(e *donburi.Entry) {
		move := component.Move.Get(e)
		obj := component.CollideBox.Get(e)

		if move.LiveTill.Before(time.Now()) {
			collision.Remove(spaceEntry, e)
			ecs.World.Remove(e.Entity())
			return
		}

		obj.X += move.Direction.X * ecs.Time.DeltaTime().Seconds()
		obj.Y += move.Direction.Y * ecs.Time.DeltaTime().Seconds()

		if obj.Y > settings.MapHeight || obj.X > settings.MapWidth {
			collision.Remove(spaceEntry, e)
			ecs.World.Remove(e.Entity())
			return
		}

		obj.Update()
	})
}

func (w *WeaponController) Draw(ecs *ecs.ECS, screen *ebiten.Image) {
	tag.Bullet.Each(ecs.World, func(e *donburi.Entry) {
		sprite := component.Shape.Get(e)
		object := component.CollideBox.Get(e)
		ebitenutil.DrawRect(screen, object.X, object.Y, sprite.Height, sprite.Width, sprite.Color)
	})
}

func getVector(num int, speed float64) []meta.Vector {
	anglStep := (2 * math.Pi) / float64(num)
	res := make([]meta.Vector, 0, num)
	angl := 0.0
	for num != 0 {
		vector := meta.NewVector(math.Cos(angl)*speed, math.Sin(angl)*speed)
		res = append(res, vector)
		angl += anglStep
		num--
	}
	return res
}
