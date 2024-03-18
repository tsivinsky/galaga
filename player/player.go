package player

import "image/color"

type Player struct {
	X      float32
	Y      float32
	Width  int
	Height int
	speed  int
	color  color.Color
}

type Options struct {
	X      float32
	Y      float32
	Width  int
	Height int
	Speed  int
	Color  color.Color
}

func New(options Options) Player {
	return Player{
		X:      options.X,
		Y:      options.Y,
		Width:  options.Width,
		Height: options.Height,
		speed:  options.Speed,
		color:  options.Color,
	}
}
