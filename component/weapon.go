package component

import (
	"time"

	"github.com/yohamta/donburi"
)

type WeaponData struct {
	Bullets  int
	Speed    float64
	Radius   float64
	NextShot time.Time
}

func NewWeaponData(bullets int, speed float64, radius float64) *WeaponData {
	return &WeaponData{
		Bullets:  bullets,
		Speed:    speed,
		Radius:   radius,
		NextShot: time.Now(),
	}
}

var Weapon = donburi.NewComponentType[WeaponData]()
