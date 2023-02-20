package component

import (
	"image/color"

	"bacteria/meta"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type SpriteData struct {
	Color  color.Color
	Height float64
	Width  float64
	Shape  meta.Shape
	Image  *ebiten.Image
}

var Sprite = donburi.NewComponentType[SpriteData]()
