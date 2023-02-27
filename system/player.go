package system

import (
	"bacteria/collision"
	"bacteria/component"
	"bacteria/factory"
	"bacteria/helper"
	"bacteria/tag"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sirupsen/logrus"
	"github.com/yohamta/donburi/ecs"
)

type PlayerController struct {
	gameCtx *helper.Context
}

func NewPlayerController(ctx *helper.Context) *PlayerController {
	return &PlayerController{gameCtx: ctx}
}

func (c *PlayerController) Update(ecs *ecs.ECS) {
	c.updatePlayer(ecs)
	c.updatePlayerDamage(ecs)
}

// nolint: gocyclo
func (c *PlayerController) updatePlayer(ecs *ecs.ECS) {
	player, ok := tag.Player.First(ecs.World)
	if !ok {
		logrus.Error("setting not found")
		return
	}

	animator := component.Animate.Get(player)
	direction := component.Direction.Get(player)
	object := collision.GetObject(player)
	if coll := object.Check(0, 0, "mob"); coll != nil {
		if direction.IsLeft() {
			animator.Play(factory.AnimationDamageLeft, false)
		} else {
			animator.Play(factory.AnimationDamageRight, false)
		}

	}

	isMoving := false
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		if object.Y > 0 {
			object.Y -= component.GetMove(player, ecs.Time.DeltaTime())
		}
		object.Update()
		isMoving = true
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		if object.Y+object.H < c.gameCtx.ScreenHeight() {
			object.Y += component.GetMove(player, ecs.Time.DeltaTime())
		}
		object.Update()
		isMoving = true
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		if object.X > 0 {
			object.X -= component.GetMove(player, ecs.Time.DeltaTime())
		}
		object.Update()
		isMoving = true
		direction.SetLeft()
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		if object.X+object.W < c.gameCtx.ScreenWidth() {
			object.X += component.GetMove(player, ecs.Time.DeltaTime())
		}
		object.Update()
		isMoving = true
		direction.SetRight()
	}

	if !isMoving {
		component.SlowSpeed(player, ecs.Time.DeltaTime())
		if direction.IsLeft() {
			animator.Play(factory.AnimationIdleLeft, true)
		} else {
			animator.Play(factory.AnimationIdleRight, true)
		}
	} else {
		if direction.IsLeft() {
			animator.Play(factory.AnimationWalkLeft, true)
		} else {
			animator.Play(factory.AnimationWalkRight, true)
		}
	}
}

func (c *PlayerController) updatePlayerDamage(ecs *ecs.ECS) {
	playerEntry, ok := tag.Player.First(ecs.World)
	if !ok {
		logrus.Error("setting not found")
		return
	}

	health := component.Health.Get(playerEntry)
	if health.Health <= 0 {
		ecs.Pause()
		return
	}

	phagacytos := component.Phagocytosis.Get(playerEntry)

	phagacytos.Proceed(health, ecs.Time.DeltaTime())

}

func (c *PlayerController) Draw(ecs *ecs.ECS, screen *ebiten.Image) {
	c.drawPlayer(ecs, screen)
}

func (c *PlayerController) drawPlayer(ecs *ecs.ECS, screen *ebiten.Image) {
	playerEntry, ok := tag.Player.First(ecs.World)
	if !ok {
		logrus.Error("player not found")
		return
	}

	object := collision.GetObject(playerEntry)
	ani := component.Animate.Get(playerEntry)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(object.X, object.Y)
	img := ani.GetImage(ecs.Time.DeltaTime())
	if img == nil {
		logrus.Error("image is nil")
		return
	}
	screen.DrawImage(img, op)
}
