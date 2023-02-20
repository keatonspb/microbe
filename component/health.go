package component

import (
	"github.com/yohamta/donburi"
)

type HealthData struct {
	Health float64
	MaxHp  float64
}

var Health = donburi.NewComponentType[HealthData]()
