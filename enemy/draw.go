package enemy

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (e *Enemy) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(e.X), float32(e.Y), float32(e.Width), float32(e.Height), color.White, false)
}
