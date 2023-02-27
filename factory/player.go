package factory

import (
	"log"
	"time"

	"bacteria/assets"
	"bacteria/collision"
	"bacteria/component"
	"bacteria/helper"
	"bacteria/helper/animation"
	"bacteria/layer"
	"bacteria/tag"

	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

const (
	AnimationIdleRight = iota
	AnimationIdleLeft
	AnimationWalkRight
	AnimationWalkLeft
	AnimationDamageLeft
	AnimationDamageRight
)

func NewPlayer(ctx *helper.Context, ecs *ecs.ECS) *donburi.Entry {
	entry := ecs.World.Entry(ecs.Create(
		layer.Default,
		tag.Player,
		component.Phagocytosis,
		component.Health,
		component.CollideBox,
		component.Velocity,
		component.Direction,
		component.Animate,
		component.Weapon,
		component.Inventory,
	))

	img, err := ctx.Storage.GetImage(assets.ImagePlayer)
	if err != nil {
		log.Fatal(err)
	}

	playerHeight := 32.0
	playerWidth := 32.0

	sprite := helper.NewSprite(img, 32, 32)
	animator := animation.NewAnimator(sprite)
	animator.AddAnimation(AnimationIdleRight, [][]int{{0, 0}, {0, 1}}, 600*time.Millisecond)
	animator.AddAnimation(AnimationIdleLeft, [][]int{{1, 4}, {1, 3}}, 600*time.Millisecond)
	animator.AddAnimation(AnimationWalkRight, [][]int{{0, 2}, {0, 3}}, 600*time.Millisecond)
	animator.AddAnimation(AnimationWalkLeft, [][]int{{1, 1}, {1, 2}}, 600*time.Millisecond)
	animator.AddAnimation(AnimationDamageLeft, [][]int{{1, 0}, {1, 0}}, 600*time.Millisecond)
	animator.AddAnimation(AnimationDamageRight, [][]int{{0, 4}, {0, 4}}, 600*time.Millisecond)

	component.Animate.Set(entry, animator)

	hd := component.HealthData{
		Health: 100,
	}
	component.Health.SetValue(entry, hd)

	collisionObj := resolv.NewObject(ctx.ScreenWidth()/2, ctx.ScreenHeight()/2, playerHeight, playerWidth, "player")
	collision.SetObject(entry, collisionObj)

	component.Velocity.Set(entry, component.NewVelocity(5, 200, 300))
	component.Phagocytosis.Set(entry, component.NewPhagocytosisData(10))
	component.Weapon.Set(entry, component.NewWeaponData(4, 200, 50))

	return entry
}
