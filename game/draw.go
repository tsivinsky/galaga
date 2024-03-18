package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/tsivinsky/galaga/fonts"
)

func (g *Game) Draw(screen *ebiten.Image) {
	if g.isGameOver {
		fonts.DrawTextInCenter(screen, "Game Over", color.White)
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

	if g.isPaused {
		fonts.DrawTextInCenter(screen, "Game Paused", color.White)
	}
}
