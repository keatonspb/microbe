package component

import "github.com/yohamta/donburi"

type ConsumableData struct {
	Amount int64
}

var Consumable = donburi.NewComponentType[ConsumableData]()
