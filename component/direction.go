package component

import "github.com/yohamta/donburi"

type DirectionData struct {
	up   bool
	left bool
}

func (d *DirectionData) SetLeft() {
	d.left = true
}

func (d *DirectionData) SetRight() {
	d.left = false
}

func (d *DirectionData) SetUp() {
	d.up = true
}

func (d *DirectionData) SetDown() {
	d.up = false
}

func (d *DirectionData) IsLeft() bool {
	return d.left
}

func (d *DirectionData) IsRight() bool {
	return !d.left
}

func (d *DirectionData) IsUp() bool {
	return d.up
}

func (d *DirectionData) IsDown() bool {
	return !d.up
}

var Direction = donburi.NewComponentType[DirectionData]()
