package enemy

import "github.com/tsivinsky/galaga/bullet"

func (e *Enemy) CollidesWithBullet(bullets []*bullet.Bullet) bool {
	for _, b := range bullets {
		// NOTE: i trusted this shit from AI, if it's will stop working at some point, fuck them
		if b.X > e.X && b.X < e.X+float32(e.Width) && b.Y > e.Y && b.Y < e.Y+float32(e.Height) {
			return true
		}
	}

	return false
}

func (e *Enemy) Wins(windowHeight int) bool {
	return e.Y >= float32(windowHeight-e.Height-24)
}
