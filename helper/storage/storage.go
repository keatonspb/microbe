package storage

import (
	"errors"
	"fmt"
	"io"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	ErrorNotFound = errors.New("asset not found")
)

type AssetKey int

type FS interface {
	Open(name string) (io.Reader, error)
}

type openFn func(path string) io.ReadCloser

type Storage struct {
	imageRegistry map[AssetKey]string
	imageContent  map[AssetKey]*ebiten.Image

	openFn func(path string) io.ReadCloser
}

func NewStorage(openFn openFn) *Storage {
	return &Storage{
		openFn:        openFn,
		imageRegistry: make(map[AssetKey]string),
		imageContent:  make(map[AssetKey]*ebiten.Image),
	}
}

func (s *Storage) AddPath(id AssetKey, path string) {
	s.imageRegistry[id] = path
}

func (s *Storage) RegisterAssets(assets map[AssetKey]string) {
	for id, path := range assets {
		s.AddPath(id, path)
	}
}

func (s *Storage) GetImage(id AssetKey) (*ebiten.Image, error) {
	if r, ok := s.imageContent[id]; ok {
		return r, nil
	}

	if p, ok := s.imageRegistry[id]; ok {
		r := s.openFn(p)

		img, _, err := ebitenutil.NewImageFromReader(r)
		if err != nil {
			return nil, fmt.Errorf("read file error")
		}

		err = r.Close()
		if err != nil {
			return nil, fmt.Errorf("failed to close asset: %w", err)
		}
		s.imageContent[id] = img

		return img, nil
	}

	return nil, ErrorNotFound
}
