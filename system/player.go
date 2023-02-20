package system

import (
	"bacteria/collision"
	"bacteria/component"
	"bacteria/tag"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sirupsen/logrus"
	"github.com/yohamta/donburi/ecs"
)

func UpdatePlayer(ecs *ecs.ECS) {
	settingEntry, ok := tag.Setting.First(ecs.World)
	if !ok {
		logrus.Error("setting not found")
		return
	}
	settings := component.Setting.Get(settingEntry)

	if settings.Stop {
		return
	}

	player, ok := tag.Player.First(ecs.World)
	if !ok {
		logrus.Error("setting not found")
		return
	}

	isMoving := false
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		object := collision.GetObject(player)
		if object.Y > 0 {
			object.Y -= component.GetMove(player, ecs.Time.DeltaTime())
		}
		object.Update()
		isMoving = true
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		object := collision.GetObject(player)
		if object.Y+object.H < settings.MapHeight {
			object.Y += component.GetMove(player, ecs.Time.DeltaTime())
		}
		object.Update()
		isMoving = true
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		object := collision.GetObject(player)
		if object.X > 0 {
			object.X -= component.GetMove(player, ecs.Time.DeltaTime())
		}
		object.Update()
		isMoving = true
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		object := collision.GetObject(player)
		if object.X+object.W < settings.MapWidth {
			object.X += component.GetMove(player, ecs.Time.DeltaTime())
		}
		object.Update()
		isMoving = true
	}

	if !isMoving {
		component.SlowSpeed(player, ecs.Time.DeltaTime())
	}
}

func UpdatePlayerDamage(ecs *ecs.ECS) {
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

func DrawPlayer(ecs *ecs.ECS, screen *ebiten.Image) {
	playerEntry, ok := tag.Player.First(ecs.World)
	if !ok {
		logrus.Error("setting not found")
		return
	}

	object := collision.GetObject(playerEntry)
	sprite := component.Sprite.Get(playerEntry)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(object.X, object.Y)
	screen.DrawImage(sprite.Image, op)
}
