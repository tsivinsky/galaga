package bullet

type Bullet struct {
	X float32
	Y float32
}

type Options struct {
	X float32
	Y float32
}

func New(options Options) *Bullet {
	return &Bullet{
		X: options.X,
		Y: options.Y,
	}
}
