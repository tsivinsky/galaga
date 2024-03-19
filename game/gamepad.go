package game

import (
	"fmt"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) updateGamepads() {
	g.gamepadIDsBuf = inpututil.AppendJustConnectedGamepadIDs(g.gamepadIDsBuf[:0])
	for _, id := range g.gamepadIDsBuf {
		if shouldConnectGamepad(id) {
			g.gamepadIDs[id] = struct{}{}
		}
	}

	for id := range g.gamepadIDs {
		if inpututil.IsGamepadJustDisconnected(id) {
			delete(g.gamepadIDs, id)
		}
	}

	for id := range g.gamepadIDs {
		maxAxis := ebiten.GamepadAxisCount(id)
		for a := 0; a < maxAxis; a++ {
			v := ebiten.GamepadAxisValue(id, a)
			g.axes[id] = append(g.axes[id], fmt.Sprintf("%f", v))
		}

		if ebiten.IsStandardGamepadLayoutAvailable(id) {
			for button := ebiten.StandardGamepadButton(0); button <= ebiten.StandardGamepadButtonMax; button++ {
				if inpututil.IsStandardGamepadButtonJustPressed(id, button) {
					g.setActiveGamepad(id)
					g.addPressedGamepadButton(id, ebiten.GamepadButton(button))
				}

				if inpututil.IsStandardGamepadButtonJustReleased(id, button) {
					g.setActiveGamepad(id)
					g.removePressedGamepadButton(id, ebiten.GamepadButton(button))
				}
			}
		}
	}
}

func (g *Game) isGamepadButtonPressed(button ebiten.StandardGamepadButton) bool {
	return inpututil.IsStandardGamepadButtonJustPressed(g.activeGamepad, button)
}

func (g *Game) setActiveGamepad(id ebiten.GamepadID) {
	g.activeGamepad = id
}

func (g *Game) addPressedGamepadButton(gamepadID ebiten.GamepadID, button ebiten.GamepadButton) {
	if g.pressedGamepadButtons[gamepadID] == nil {
		g.pressedGamepadButtons[gamepadID] = make(map[ebiten.GamepadButton]bool)
	}
	g.pressedGamepadButtons[gamepadID][button] = true
}

func (g *Game) removePressedGamepadButton(gamepadID ebiten.GamepadID, button ebiten.GamepadButton) {
	delete(g.pressedGamepadButtons[gamepadID], button)
}

func shouldConnectGamepad(id ebiten.GamepadID) bool {
	name := strings.ToLower(ebiten.GamepadName(id))
	return !strings.Contains(name, "touchpad")
}
