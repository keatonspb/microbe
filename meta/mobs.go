package meta

import (
	"image/color"
	"math/rand"
)

const (
	MobTypeA = iota
	MobTypeB
	MobTypeC
)

var availableMobTypes = []int{MobTypeA, MobTypeB, MobTypeC}

type Mob struct {
	Type   int
	Color  color.Color
	Speed  float64
	Damage float64 //per second
}

var Mobs = map[int]Mob{
	MobTypeA: {
		Type: MobTypeA,
		Color: color.RGBA{
			R: 161,
			G: 215,
			B: 217,
			A: 255,
		},
		Speed:  80,
		Damage: 0.5,
	},
	MobTypeB: {
		Type: MobTypeB,
		Color: color.RGBA{
			R: 192,
			G: 129,
			B: 129,
			A: 255,
		},
		Speed:  50,
		Damage: 0.3,
	},
	MobTypeC: {
		Type: MobTypeC,
		Color: color.RGBA{
			R: 164,
			G: 215,
			B: 222,
			A: 255,
		},
		Speed:  30,
		Damage: 0.1,
	},
}

func RandMobType() int {
	return availableMobTypes[rand.Intn(len(availableMobTypes))]
}
