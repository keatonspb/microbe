package test

import (
	"io"
	"os"
	"testing"

	"bacteria/helper/storage"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_Storage(t *testing.T) {
	s := storage.NewStorage(func(path string) io.ReadCloser {
		f, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		return f
	})

	s.AddPath(1, "test.png")

	r, err := s.GetImage(1)
	require.NoError(t, err)
	w, h := r.Size()
	assert.Equal(t, 2, w)
	assert.Equal(t, 2, h)

	//from cache
	r, err = s.GetImage(1)
	require.NoError(t, err)
	w, h = r.Size()
	assert.Equal(t, 2, w)
	assert.Equal(t, 2, h)
}
