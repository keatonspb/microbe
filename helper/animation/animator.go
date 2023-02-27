package animation

import (
	"time"

	"bacteria/helper"

	"github.com/hajimehoshi/ebiten/v2"
)

type Animator struct {
	sprite           *helper.Sprite
	animations       map[int]*Sequence
	currentAnimation int
	oneTimeAnimation *int
}

func NewAnimator(sprite *helper.Sprite) *Animator {
	return &Animator{
		sprite:     sprite,
		animations: make(map[int]*Sequence),
	}
}

func (a *Animator) AddAnimation(id int, s [][]int, fps time.Duration) {
	seq := NewSequence(fps)

	frames := make([]*ebiten.Image, 0, len(s))
	for _, el := range s {
		frames = append(frames, a.sprite.GetCell(el[0], el[1]))
	}

	seq.SetFrames(frames)
	a.animations[id] = seq
}

func (a *Animator) Play(id int, repeateable bool) {
	if a.currentAnimation == id {
		return
	}
	if repeateable {
		a.oneTimeAnimation = &id
	}

	a.animations[id].Start(repeateable)
	a.currentAnimation = id
}

func (a *Animator) GetImage(delta time.Duration) *ebiten.Image {
	toPlay := a.currentAnimation
	if a.oneTimeAnimation != nil {
		toPlay = *a.oneTimeAnimation
		if a.animations[toPlay].IsFinished() {
			a.oneTimeAnimation = nil
		}
	}
	return a.animations[toPlay].GetImage(delta)
}
