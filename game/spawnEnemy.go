package game

import (
	"game/enemy"
	"math/rand"
)

func (g *Game) SpawnEnemy() {
	x, y := generateRandomCoords(g.windowWidth, g.windowHeight)

	g.enemies = append(g.enemies, enemy.New(enemy.Options{
		X: float32(x),
		Y: float32(y),
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
