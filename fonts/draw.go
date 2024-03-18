package fonts

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

func DrawTextInCenter(screen *ebiten.Image, msg string, color color.Color) {
	textBounds, _ := font.BoundString(IBMPlexMono, msg)
	dx := textBounds.Max.X.Ceil() - textBounds.Min.X.Ceil()
	dy := textBounds.Max.Y.Ceil() - textBounds.Min.Y.Ceil()
	text.Draw(
		screen,
		msg,
		IBMPlexMono,
		screen.Bounds().Min.X+(screen.Bounds().Dx()-dx)/2,
		screen.Bounds().Min.Y+(screen.Bounds().Dy()+dy)/2,
		color,
	)
}
	)
}
