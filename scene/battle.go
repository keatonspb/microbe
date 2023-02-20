package scene

import (
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
}

func NewBattle(screenWidth, screenHeight float64) *Battle {
	s := &Battle{
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
	}

	s.init()

	return s
}

func (s *Battle) init() {
	s.ecs = ecs.NewECS(donburi.NewWorld())
	space := factory.CreateSpace(s.ecs, s.screenWidth, s.screenHeight)

	s.weaponController = system.NewWeaponController(s.screenWidth, s.screenHeight, space)
	s.floatingController = system.NewFloatController(s.screenWidth, s.screenHeight, space)

	pl := factory.NewPlayer(s.ecs, s.screenWidth, s.screenHeight)
	collision.AddToSpace(space, pl)
	inv := component.Inventory.Get(pl)

	cellController := system.NewCellController(s.screenWidth, s.screenHeight, space, 4, inv)

	s.ecs.AddSystem(system.UpdatePlayerDamage)
	s.ecs.AddSystem(system.GenerateMob)
	s.ecs.AddSystem(system.CollideMob)
	s.ecs.AddSystem(s.weaponController.Update)
	s.ecs.AddSystem(s.floatingController.Update)
	s.ecs.AddSystem(cellController.Update)
	s.ecs.AddSystem(system.UpdatePlayer)

	s.ecs.AddRenderer(layer.Default, system.DrawMob)
	s.ecs.AddRenderer(layer.Default, system.DrawStatus)
	s.ecs.AddRenderer(layer.Default, s.weaponController.Draw)
	s.ecs.AddRenderer(layer.Default, cellController.Draw)
	s.ecs.AddRenderer(layer.Default, system.DrawPlayer)

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
