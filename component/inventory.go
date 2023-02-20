package component

import "github.com/yohamta/donburi"

type InventoryData struct {
	Cells int64
}

var Inventory = donburi.NewComponentType[InventoryData]()
