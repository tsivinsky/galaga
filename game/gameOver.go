package game

import "github.com/hajimehoshi/ebiten/v2"

func (g *Game) handleGameOver() error {
	if g.isGamepadButtonPressed(ebiten.StandardGamepadButtonRightBottom) {
		return ebiten.Termination
	}

	return nil
}
