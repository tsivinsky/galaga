package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const shootInterval = 100 * time.Millisecond
const enemySpawnInterval = 1 * time.Second

func (g *Game) Update() error {
	g.updateGamepads()

	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		return ebiten.Termination
	}

	if g.isGameOver {
		// handle game over screen maybe
		return nil
	}

	if g.isVictory {
		return g.handleVictory()
	}

	isOptionsPressed := g.isGamepadButtonPressed(ebiten.StandardGamepadButtonCenterRight)
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
		xAxis = ebiten.GamepadAxisValue(g.activeGamepad, 0)
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

// returns enemy index if hit, -1 otherwise
func (g *Game) shotEnemy() int {
	for i, e := range g.enemies {
		if e.CollidesWithBullet(g.bullets) {
			return i
		}
	}

	return -1
}
