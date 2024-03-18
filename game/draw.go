package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (g *Game) Draw(screen *ebiten.Image) {
	if g.isGameOver {
		ebitenutil.DebugPrint(screen, "Game Over")
		return
	}

	vector.DrawFilledRect(screen, 0, 0, float32(g.windowWidth), float32(g.windowHeight), color.Black, false)

	g.player.Draw(screen)

	for _, b := range g.bullets {
		b.Draw(screen)
	}

	for _, e := range g.enemies {
		e.Draw(screen)
	}
}
