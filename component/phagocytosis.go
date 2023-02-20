package component

import (
	"time"

	"github.com/yohamta/donburi"
)

type PhagocytosisData struct {
	Attackers     []*MobData
	NeedAttackers int
}

func NewPhagocytosisData(needAttackers int) *PhagocytosisData {
	return &PhagocytosisData{
		NeedAttackers: needAttackers,
	}
}

func (a *PhagocytosisData) Proceed(e *HealthData, d time.Duration) {
	if len(a.Attackers) < a.NeedAttackers {
		return
	}

	for _, attacker := range a.Attackers {
		e.Health -= attacker.GetDamage(d)
	}
}

func (a *PhagocytosisData) AddAttacker(e *MobData) {
	a.Attackers = append(a.Attackers, e)
}

var Phagocytosis = donburi.NewComponentType[PhagocytosisData]()
