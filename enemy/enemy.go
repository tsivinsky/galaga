package enemy

type Enemy struct {
	X      float32
	Y      float32
	Width  int
	Height int
	Speed  float32
}

type Options struct {
	X     float32
	Y     float32
	Speed float32
}

func New(options Options) *Enemy {
	return &Enemy{
		X:      options.X,
		Y:      options.Y,
		Width:  40,
		Height: 20,
		Speed:  options.Speed,
	}
}
