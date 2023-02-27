package component

import (
	"image/color"

	"bacteria/meta"

	"github.com/yohamta/donburi"
)

type ShapeData struct {
	Color  color.Color
	Form   meta.Shape
	Height float64
	Width  float64
}

var Shape = donburi.NewComponentType[ShapeData]()
