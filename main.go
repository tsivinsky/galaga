package main

import (
	"image/color"
	"log"

	"github.com/tsivinsky/galaga/game"
	"github.com/tsivinsky/galaga/player"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowTitle("Hello World!")
	ebiten.SetFullscreen(true)
	ebiten.SetCursorMode(ebiten.CursorModeHidden)

	p := player.New(player.Options{
		X:      0,
		Y:      0,
		Width:  80,
		Height: 24,
		Speed:  24,
		Color:  color.White,
	})

	g := game.New(game.GameOptions{
		Player:       p,
		WindowWidth:  640,
		WindowHeight: 480,
	})

	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
