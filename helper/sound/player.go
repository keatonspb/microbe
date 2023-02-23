package sound

import (
	"fmt"
	"io"
	"os"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/sirupsen/logrus"
)

type Player struct {
	sounds       map[int64]*audio.Player
	audioContext *audio.Context
	sampleRate   int
}

func NewPlayer(context *audio.Context, sampleRate int) (*Player, error) {
	return &Player{
		audioContext: context,
		sounds:       make(map[int64]*audio.Player),
		sampleRate:   sampleRate,
	}, nil
}

func (p *Player) AddSoundFromPath(id int64, path string) error {
	r, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open sound file: %w", err)
	}

	return p.AddSound(id, r)
}

func (p *Player) AddSound(id int64, r io.Reader) error {
	audioStream, err := wav.DecodeWithSampleRate(p.sampleRate, r)
	if err != nil {
		return err
	}

	player, err := p.audioContext.NewPlayer(audioStream)
	if err != nil {
		return err
	}

	p.sounds[id] = player

	return nil
}

func (p *Player) PlaySound(id int64) {
	pl, ok := p.sounds[id]
	if !ok {
		logrus.Errorf("sound with id %d not found", id)
	}
	err := pl.Rewind()
	if err != nil {
		logrus.WithError(err).Error("failed to rewind sound")
	}
	pl.Play()
}
