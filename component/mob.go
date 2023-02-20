package component

import (
	"time"

	"bacteria/helper"
	"bacteria/meta"

	"github.com/yohamta/donburi"
)

type MobData struct {
	Type   int
	Attack bool
}

func (m *MobData) Move(d time.Duration) (dx, dy float64) {
	t := meta.Mobs[m.Type]
	dy = t.Speed * d.Seconds()
	dx = t.Speed * d.Seconds()
	if helper.RandBool() {
		dx = -dx
	}
	return dx, dy
}

func (m *MobData) GetDamage(d time.Duration) float64 {
	return meta.Mobs[m.Type].Damage * d.Seconds()
}

var Mob = donburi.NewComponentType[MobData]()
