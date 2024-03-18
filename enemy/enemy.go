package enemy

type Enemy struct {
	X      float32
	Y      float32
	Width  int
	Height int
}

type Options struct {
	X float32
	Y float32
}

func New(options Options) *Enemy {
	return &Enemy{
		X:      options.X,
		Y:      options.Y,
		Width:  40,
		Height: 20,
	}
}
