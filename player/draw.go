package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (p Player) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(p.X), float32(p.Y), float32(p.Width), float32(p.Height), p.color, false)
}
