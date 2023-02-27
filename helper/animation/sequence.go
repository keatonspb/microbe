package animation

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Getter interface {
	GetCell(row, coll int) *ebiten.Image
}

type Sequence struct {
	playTime   time.Duration
	maxTime    time.Duration
	fps        time.Duration
	frames     []*ebiten.Image
	repeatable bool
	isFinished bool
}

func NewSequence(fps time.Duration) *Sequence {
	return &Sequence{
		fps: fps,
	}
}

func (s *Sequence) IsFinished() bool {
	return s.isFinished
}

func (s *Sequence) SetFrames(img []*ebiten.Image) {
	s.frames = make([]*ebiten.Image, len(img))
	copy(s.frames, img)
	s.maxTime = time.Duration(len(img)) * s.fps
}

func (s *Sequence) Start(repeatable bool) {
	s.playTime = 0
	s.repeatable = repeatable
	s.isFinished = false
}

func (s *Sequence) GetImage(delay time.Duration) *ebiten.Image {
	s.playTime += delay
	if s.playTime > s.maxTime {
		s.playTime = s.playTime % s.maxTime
		s.isFinished = true
		if !s.repeatable {
			return s.frames[len(s.frames)-1]
		}
	}

	return s.frames[s.playTime/s.fps]
}
