package animation

import (
	"testing"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/stretchr/testify/require"
)

func TestUnit_Sequence(t *testing.T) {
	s := NewSequence(time.Second / 2)

	img1 := ebiten.NewImage(10, 10)
	img2 := ebiten.NewImage(10, 10)
	img3 := ebiten.NewImage(10, 10)
	s.SetFrames([]*ebiten.Image{img1, img2, img3})

	s.Start(true)

	tCases := []struct {
		delay time.Duration
		want  *ebiten.Image
	}{
		{0, img1},                       //0 sec
		{time.Millisecond * 300, img1},  //0.3 sec
		{time.Millisecond * 200, img2},  //0.5 sec
		{time.Millisecond * 200, img2},  //0.7 sec
		{time.Millisecond * 400, img3},  //1.1 sec
		{time.Millisecond * 1000, img2}, //2.1 sec == 0.6 sec
	}

	for _, tc := range tCases {
		actual := s.GetImage(tc.delay)
		require.Equal(t, tc.want, actual)
	}

}
