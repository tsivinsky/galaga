package game

import "game/bullet"

func (g *Game) Shoot() {
	g.bullets = append(g.bullets, bullet.New(bullet.Options{
		X: g.player.X + float32(g.player.Width/2),
		Y: g.player.Y - 10,
	}))
}
