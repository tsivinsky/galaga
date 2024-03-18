package game

import (
	"fmt"
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

	if g.isVictory {
		fonts.DrawTextInCenter(screen, "Victory", color.White)
		return
	}

	g.drawBackground(screen)
	g.drawEnemyCount(screen)
	g.player.Draw(screen)
	g.drawEnemies(screen)
	g.drawBullets(screen)

	if g.isPaused {
		fonts.DrawTextInCenter(screen, "Game Paused", color.White)
	}
}

func (g *Game) drawBackground(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, 0, 0, float32(g.windowWidth), float32(g.windowHeight), color.Black, false)
}

func (g *Game) drawBullets(screen *ebiten.Image) {
	for _, b := range g.bullets {
		b.Draw(screen)
	}
}

func (g *Game) drawEnemies(screen *ebiten.Image) {
	for _, e := range g.enemies {
		e.Draw(screen)
	}
}

func (g *Game) drawEnemyCount(screen *ebiten.Image) {
	msg := fmt.Sprintf("Defeated: %d", g.defeatedEnemiesCount)
	fonts.DrawTextInCoords(screen, msg, 10, 10, color.White)
}
