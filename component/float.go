package component

import (
	"time"

	"bacteria/helper"

	"github.com/yohamta/donburi"
)

type FloatData struct {
	Speed float64
}

func (m *FloatData) Move(d time.Duration) (dx, dy float64) {
	dy = m.Speed * d.Seconds()
	dx = m.Speed * d.Seconds()
	if helper.RandBool() {
		dx = -dx
	}
	return dx, dy
}

var Float = donburi.NewComponentType[FloatData]()
