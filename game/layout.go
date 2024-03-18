package game

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	g.windowWidth = outsideWidth
	g.windowHeight = outsideHeight

	g.player.Y = float32(outsideHeight - 100)

	return outsideWidth, outsideHeight
}
