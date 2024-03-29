package player

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type moveDirection = string

const (
	left  moveDirection = "left"
	right moveDirection = "right"
)

func (p *Player) Update(windowWidth, windowHeight int, xAxis float64) {
	if xAxis < -0.3 {
		dx := float32(p.speed) * float32(xAxis) * -1
		p.move(left, dx)
	} else if ebiten.IsKeyPressed(ebiten.KeyH) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.move(left, float32(p.speed))
	}

	if xAxis > 0.3 {
		dx := float32(p.speed) * float32(xAxis)
		p.move(right, dx)
	} else if ebiten.IsKeyPressed(ebiten.KeyL) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.move(right, float32(p.speed))
	}

	p.checkCollision(windowWidth, windowHeight)
}

func (p *Player) move(direction moveDirection, dx float32) {
	switch direction {
	case left:
		p.X -= dx
	case right:
		p.X += dx
	}
}

func (p *Player) checkCollision(maxWidth, maxHeight int) {
	if p.X < 0 {
		p.X = 0
	}

	if p.X > float32(maxWidth-p.Width) {
		p.X = float32(maxWidth - p.Width)
	}
}
