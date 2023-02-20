package component

import (
	"time"

	"github.com/yohamta/donburi"
)

type VelocityData struct {
	InitialSpeed float64 //pixels per second
	MaxSpeed     float64
	Speed        float64 //pixels per second
	Acceleration float64 //speed increase per second
}

func NewVelocity(initialSpeed, maxSpeed, acceleration float64) *VelocityData {
	return &VelocityData{
		InitialSpeed: initialSpeed,
		MaxSpeed:     maxSpeed,
		Speed:        initialSpeed,
		Acceleration: acceleration,
	}
}

func GetMove(player *donburi.Entry, d time.Duration) float64 {
	v := Velocity.Get(player)
	if v.Speed < v.MaxSpeed {
		v.Speed = v.Speed + v.Acceleration*d.Seconds()
	}
	return v.Speed * d.Seconds()
}

func SlowSpeed(player *donburi.Entry, d time.Duration) {
	v := Velocity.Get(player)
	if v.Speed > 0 {
		v.Speed = v.Speed - v.Acceleration*2*d.Seconds()
	}
}

var Velocity = donburi.NewComponentType[VelocityData]()
