package component

import (
	"time"

	"bacteria/meta"

	"github.com/yohamta/donburi"
)

type MoveData struct {
	Direction meta.Vector
	LiveTill  time.Time
}

var Move = donburi.NewComponentType[MoveData]()
