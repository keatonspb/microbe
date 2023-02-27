package helper

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sirupsen/logrus"
)

type Sprite struct {
	cellHeight float64
	cellWidth  float64
	img        *ebiten.Image
	matrix     [][]*ebiten.Image
}

func NewSprite(img *ebiten.Image, cellWidth, cellHeight float64) *Sprite {
	s := &Sprite{
		cellHeight: cellHeight,
		cellWidth:  cellWidth,
		img:        img,
	}
	s.init()

	return s
}

func (s *Sprite) init() {
	width, height := s.img.Size()
	logrus.Infof("width: %d, height: %d", width, height)
	colls := width / int(s.cellWidth)
	rows := height / int(s.cellHeight)

	s.matrix = make([][]*ebiten.Image, rows)

	x1, y1 := 0, 0
	cW := int(s.cellWidth)
	cH := int(s.cellHeight)

	for i := 0; i < rows; i++ {
		s.matrix[i] = make([]*ebiten.Image, colls)
		for j := 0; j < colls; j++ {
			subImg, ok := s.img.SubImage(image.Rect(x1, y1, x1+cW, y1+cH)).(*ebiten.Image)
			if !ok {
				logrus.Errorf("failed to convert image to ebiten.Image")
				continue
			}
			s.matrix[i][j] = subImg
			x1 += cW
		}
		y1 += cH
		x1 = 0
	}
}

func (s *Sprite) GetImage() *ebiten.Image {
	return s.img
}

func (s *Sprite) GetCell(row, coll int) *ebiten.Image {
	return s.matrix[row][coll]
}

func (s *Sprite) GetCellHeight() float64 {
	return s.cellHeight
}

func (s *Sprite) GetCellWidth() float64 {
	return s.cellWidth
}
