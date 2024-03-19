package game

import "github.com/hajimehoshi/ebiten/v2"

func (g *Game) handleVictory() error {
	if g.isGamepadButtonPressed(ebiten.StandardGamepadButtonRightBottom) {
		return ebiten.Termination
	}

	return nil
}
