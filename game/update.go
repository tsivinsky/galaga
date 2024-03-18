package game

import (
	"fmt"
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

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
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

	xAxis := ebiten.GamepadAxisValue(0, 0)
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

	return nil
}

func (g *Game) handleGamepads() {
	if g.gamepadIDs == nil {
		g.gamepadIDs = map[ebiten.GamepadID]struct{}{}
	}

	g.gamepadIDsBuf = inpututil.AppendJustConnectedGamepadIDs(g.gamepadIDsBuf[:0])
	for _, id := range g.gamepadIDsBuf {
		g.gamepadIDs[id] = struct{}{}
	}
	for id := range g.gamepadIDs {
		if inpututil.IsGamepadJustDisconnected(id) {
			delete(g.gamepadIDs, id)
		}
	}

	g.axes = map[ebiten.GamepadID][]string{}
	for id := range g.gamepadIDs {
		maxAxis := ebiten.GamepadAxisCount(id)
		for a := 0; a < maxAxis; a++ {
			v := ebiten.GamepadAxisValue(id, a)
			g.axes[id] = append(g.axes[id], fmt.Sprintf("%d: %f", a, v))
		}
	}
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
