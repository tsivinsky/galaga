package fonts

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	_ "embed"
)

var (
	//go:embed IBMPlexMono-Regular.ttf
	IBMPlexMono_ttf []byte
)

var (
	IBMPlexMono      font.Face
	IBMPlexMonoLarge font.Face
)

func init() {
	tt, err := opentype.Parse(IBMPlexMono_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72

	IBMPlexMono, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	IBMPlexMono = text.FaceWithLineHeight(IBMPlexMono, 32)

	IBMPlexMonoLarge, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    96,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	IBMPlexMonoLarge = text.FaceWithLineHeight(IBMPlexMonoLarge, 128)
}
