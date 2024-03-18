package game

import (
	"math/rand"

	"github.com/tsivinsky/galaga/enemy"
)

func (g *Game) calculateEnemySpeed() float32 {
	if g.defeatedEnemiesCount <= 5 {
		return 1
	}

	return 1.5
}

func (g *Game) SpawnEnemy() {
	x, y := generateRandomCoords(g.windowWidth, g.windowHeight)

	enemySpeed := g.calculateEnemySpeed()
	g.enemies = append(g.enemies, enemy.New(enemy.Options{
		X:     float32(x),
		Y:     float32(y),
		Speed: enemySpeed,
	}))
}

func generateRandomNumberFromRange(from, to int) int {
	return rand.Intn(to-from) + from
}

func generateRandomCoords(maxWidth, maxHeight int) (int, int) {
	x := generateRandomNumberFromRange(40, maxWidth-40)
	y := rand.Intn(30)

	return x, y
}
