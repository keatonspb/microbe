package component

import (
	"bacteria/helper/animation"

	"github.com/yohamta/donburi"
)

var Animate = donburi.NewComponentType[animation.Animator]()
