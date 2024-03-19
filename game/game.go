package game

import (
	"time"

	"github.com/tsivinsky/galaga/bullet"
	"github.com/tsivinsky/galaga/enemy"
	"github.com/tsivinsky/galaga/player"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	player                player.Player
	windowWidth           int
	windowHeight          int
	bullets               []*bullet.Bullet
	gamepadIDsBuf         []ebiten.GamepadID
	gamepadIDs            map[ebiten.GamepadID]struct{}
	activeGamepad         ebiten.GamepadID
	axes                  map[ebiten.GamepadID][]string
	pressedGamepadButtons map[ebiten.GamepadID]map[ebiten.GamepadButton]bool
	lastShootingTime      time.Time
	enemies               []*enemy.Enemy
	lastEnemySpawn        time.Time
	isGameOver            bool
	isPaused              bool
	isVictory             bool
	defeatedEnemiesCount  int
}

type GameOptions struct {
	Player       player.Player
	WindowWidth  int
	WindowHeight int
}

func New(options GameOptions) Game {
	return Game{
		player:                options.Player,
		windowWidth:           options.WindowWidth,
		windowHeight:          options.WindowHeight,
		bullets:               []*bullet.Bullet{},
		enemies:               []*enemy.Enemy{},
		gamepadIDsBuf:         make([]ebiten.GamepadID, 10),
		gamepadIDs:            make(map[ebiten.GamepadID]struct{}),
		axes:                  make(map[ebiten.GamepadID][]string),
		pressedGamepadButtons: make(map[ebiten.GamepadID]map[ebiten.GamepadButton]bool),
	}
}
