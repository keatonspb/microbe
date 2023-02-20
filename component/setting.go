package component

import "github.com/yohamta/donburi"

type SettingData struct {
	MapWidth  float64
	MapHeight float64
	MaxMobs   int
	Stop      bool
}

var Setting = donburi.NewComponentType[SettingData]()
