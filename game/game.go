package game

import (
	"game/bullet"
	"game/enemy"
	"game/player"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	player           player.Player
	windowWidth      int
	windowHeight     int
	bullets          []*bullet.Bullet
	gamepadIDsBuf    []ebiten.GamepadID
	gamepadIDs       map[ebiten.GamepadID]struct{}
	axes             map[ebiten.GamepadID][]string
	lastShootingTime time.Time
	enemies          []*enemy.Enemy
	lastEnemySpawn   time.Time
	isGameOver       bool
}

type GameOptions struct {
	Player       player.Player
	WindowWidth  int
	WindowHeight int
}

func New(options GameOptions) Game {
	return Game{
		player:       options.Player,
		windowWidth:  options.WindowWidth,
		windowHeight: options.WindowHeight,
		bullets:      []*bullet.Bullet{},
		enemies:      []*enemy.Enemy{},
	}
}
