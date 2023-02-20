package scene

import (
	"embed"
	"image/color"

	"bacteria/collision"
	"bacteria/component"
	"bacteria/factory"
	"bacteria/layer"
	"bacteria/system"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

// Battle main battle scene
type Battle struct {
	ecs                       *ecs.ECS
	screenWidth, screenHeight float64
	weaponController          *system.WeaponController
	floatingController        *system.FloatController
	fs                        *embed.FS
}

func NewBattle(screenWidth, screenHeight float64, fs *embed.FS) *Battle {
	s := &Battle{
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		fs:           fs,
	}

	s.init()

	return s
}

func (s *Battle) init() {
	s.ecs = ecs.NewECS(donburi.NewWorld())
	space := factory.CreateSpace(s.ecs, s.screenWidth, s.screenHeight)
	pl := factory.NewPlayer(s.ecs, s.screenWidth, s.screenHeight, s.fs)
	collision.AddToSpace(space, pl)
	inv := component.Inventory.Get(pl)

	playerController := system.NewPlayerController(s.screenWidth, s.screenHeight)
	s.ecs.AddSystem(playerController.Update)
	s.ecs.AddRenderer(layer.Default, playerController.Draw)

	s.weaponController = system.NewWeaponController(s.screenWidth, s.screenHeight, space)
	s.ecs.AddSystem(s.weaponController.Update)
	s.ecs.AddRenderer(layer.Default, s.weaponController.Draw)

	s.floatingController = system.NewFloatController(s.screenWidth, s.screenHeight, space)
	s.ecs.AddSystem(s.floatingController.Update)

	mobController := system.NewMobController(s.fs)
	s.ecs.AddSystem(mobController.Update)
	s.ecs.AddRenderer(layer.Default, mobController.Draw)

	cellController := system.NewCellController(s.screenWidth, s.screenHeight, space, 4, inv, s.fs)
	s.ecs.AddSystem(cellController.Update)
	s.ecs.AddRenderer(layer.Default, cellController.Draw)

	s.ecs.AddRenderer(layer.Default, system.DrawStatus)

	factory.NewSetting(s.ecs, s.screenWidth, s.screenHeight, 50)
}

func (s *Battle) Update() {
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		s.init()
		return
	}
	s.ecs.Update()
}

func (s *Battle) Draw(screen *ebiten.Image) {
	screen.Clear()
	screen.Fill(color.RGBA{35, 2, 41, 255})
	s.ecs.Draw(screen)
}
