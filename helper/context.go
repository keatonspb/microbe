package helper

import (
	"bacteria/helper/sound"
	"bacteria/helper/storage"
)

type Context struct {
	Storage *storage.Storage
	Audio   *sound.Player

	screenHeight float64
	screenWidth  float64
}

func NewContext() *Context {
	return &Context{}
}

func (c *Context) SetScreenSize(width, height float64) {
	c.screenWidth = width
	c.screenHeight = height
}

func (c *Context) ScreenWidth() float64 {
	return c.screenWidth
}

func (c *Context) ScreenHeight() float64 {
	return c.screenHeight
}
