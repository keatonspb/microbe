package scene

import (
	"fmt"
	"image/color"

	"bacteria/assets"
	"bacteria/collision"
	"bacteria/component"
	"bacteria/factory"
	"bacteria/helper"
	"bacteria/layer"
	"bacteria/system"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

// Battle main battle scene
type Battle struct {
	ecs                *ecs.ECS
	weaponController   *system.WeaponController
	floatingController *system.FloatController
	gameContext        *helper.Context
}

func NewBattle(ctx *helper.Context) (*Battle, error) {
	s := &Battle{
		gameContext: ctx,
	}

	err := s.init()
	if err != nil {
		return nil, fmt.Errorf("failed to init battle scene: %w", err)
	}

	return s, nil
}

func (s *Battle) init() error {
	s.ecs = ecs.NewECS(donburi.NewWorld())
	s.gameContext.Storage.RegisterAssets(assets.ImagePaths)

	space := factory.CreateSpace(s.gameContext, s.ecs)

	pl := factory.NewPlayer(s.gameContext, s.ecs)
	collision.AddToSpace(space, pl)
	inv := component.Inventory.Get(pl)

	playerController := system.NewPlayerController(s.gameContext)
	s.ecs.AddSystem(playerController.Update)
	s.ecs.AddRenderer(layer.Default, playerController.Draw)

	wc, err := system.NewWeaponController(s.gameContext, space)
	if err != nil {
		return fmt.Errorf("failed to create weapon controller: %w", err)
	}
	s.weaponController = wc

	s.ecs.AddSystem(s.weaponController.Update)
	s.ecs.AddRenderer(layer.Default, s.weaponController.Draw)

	s.floatingController = system.NewFloatController(s.gameContext, space)
	s.ecs.AddSystem(s.floatingController.Update)

	mobController, err := system.NewMobController(s.gameContext)
	if err != nil {
		return fmt.Errorf("failed to create mob controller: %w", err)
	}

	s.ecs.AddSystem(mobController.Update)
	s.ecs.AddRenderer(layer.Default, mobController.Draw)

	cellController := system.NewCellController(s.gameContext, space, 4, inv)
	s.ecs.AddSystem(cellController.Update)
	s.ecs.AddRenderer(layer.Default, cellController.Draw)

	s.ecs.AddRenderer(layer.Default, system.DrawStatus)

	factory.NewSetting(s.ecs, s.gameContext.ScreenWidth(), s.gameContext.ScreenHeight(), 50)
	return nil
}

func (s *Battle) Update() {
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		err := s.init()
		if err != nil {
			panic(err)
		}
		return
	}
	s.ecs.Update()
}

func (s *Battle) Draw(screen *ebiten.Image) {
	screen.Clear()
	screen.Fill(color.RGBA{35, 2, 41, 255})
	s.ecs.Draw(screen)
}
