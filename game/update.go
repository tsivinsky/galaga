package game

import (
	"fmt"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const shootInterval = 100 * time.Millisecond
const enemySpawnInterval = 1 * time.Second

func (g *Game) Update() error {
	g.handleGamepads()

	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		return ebiten.Termination
	}

	if g.isGameOver {
		// handle game over screen maybe
		return nil
	}

	if g.isVictory {
		return nil
	}

	isOptionsPressed := false
	for id := range g.pressedGamepadButtons {
		for button := range g.pressedGamepadButtons[id] {
			if button == ebiten.GamepadButton(9) {
				isOptionsPressed = true
			}
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) || isOptionsPressed {
		g.isPaused = !g.isPaused
		return nil
	}

	if g.isPaused {
		return nil
	}

	now := time.Now()
	if now.Sub(g.lastShootingTime) >= shootInterval {
		g.Shoot()
		g.lastShootingTime = now
	}

	if now.Sub(g.lastEnemySpawn) >= enemySpawnInterval {
		g.SpawnEnemy()
		g.lastEnemySpawn = now
	}

	xAxis := 0.0
	if len(g.gamepadIDs) > 0 {
		xAxis = ebiten.GamepadAxisValue(0, 0)
	}
	g.player.Update(g.windowWidth, g.windowHeight, xAxis)

	// TODO: need to delete bullets that are outside the screen
	// don't wanna be like EA or Ubisoft lol

	for _, b := range g.bullets {
		b.Update()
	}

	for _, e := range g.enemies {
		e.Update()
		if e.Wins(g.windowHeight) {
			g.isGameOver = true
		}
	}

	if index := g.shotEnemy(); index != -1 {
		g.defeatedEnemiesCount++
		g.enemies = append(g.enemies[:index], g.enemies[index+1:]...)
		for _, e := range g.enemies {
			e.Speed = g.calculateEnemySpeed()
		}
	}

	const niceNumber = 69
	if g.defeatedEnemiesCount == niceNumber {
		g.isVictory = true
		return nil
	}

	return nil
}

func (g *Game) handleGamepads() {
	if g.gamepadIDs == nil {
		g.gamepadIDs = map[ebiten.GamepadID]struct{}{}
	}

	g.gamepadIDsBuf = inpututil.AppendJustConnectedGamepadIDs(g.gamepadIDsBuf[:0])
	for _, id := range g.gamepadIDsBuf {
		name := ebiten.GamepadName(id)
		if strings.Contains(strings.ToLower(name), "touchpad") {
			continue
		}
		g.gamepadIDs[id] = struct{}{}
	}
	for id := range g.gamepadIDs {
		if inpututil.IsGamepadJustDisconnected(id) {
			delete(g.gamepadIDs, id)
		}
	}

	g.axes = map[ebiten.GamepadID][]string{}
	g.pressedGamepadButtons = map[ebiten.GamepadID]map[ebiten.GamepadButton]bool{}
	for id := range g.gamepadIDs {
		maxAxis := ebiten.GamepadAxisCount(id)
		for a := 0; a < maxAxis; a++ {
			v := ebiten.GamepadAxisValue(id, a)
			g.axes[id] = append(g.axes[id], fmt.Sprintf("%d: %f", a, v))
		}

		if ebiten.IsStandardGamepadLayoutAvailable(id) {
			for b := ebiten.StandardGamepadButton(0); b <= ebiten.StandardGamepadButtonMax; b++ {
				if inpututil.IsStandardGamepadButtonJustPressed(id, b) {
					g.addPressedGamepadButton(id, ebiten.GamepadButton(b))
				}
				if inpututil.IsStandardGamepadButtonJustReleased(id, b) {
					g.removePressedGamepadButton(id, ebiten.GamepadButton(b))
				}
			}
		}
	}
}

func (g *Game) addPressedGamepadButton(gamepadId ebiten.GamepadID, button ebiten.GamepadButton) {
	if g.pressedGamepadButtons[gamepadId] == nil {
		g.pressedGamepadButtons[gamepadId] = map[ebiten.GamepadButton]bool{}
	}
	g.pressedGamepadButtons[gamepadId][button] = true
}

func (g *Game) removePressedGamepadButton(gamepadId ebiten.GamepadID, button ebiten.GamepadButton) {
	delete(g.pressedGamepadButtons[gamepadId], button)
}

// returns enemy index if hit, -1 otherwise
func (g *Game) shotEnemy() int {
	for i, e := range g.enemies {
		if e.CollidesWithBullet(g.bullets) {
			return i
		}
	}

	return -1
}
