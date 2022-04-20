package assets

import (
	"bytes"
	_ "embed"
	"image"
	"image/png"
	"io"
	"log"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed images/textures/sea.png
var bg []byte

//go:embed images/char/boat.png
var boat []byte

//go:embed images/char/boat-2.png
var boat2 []byte

func Background() image.Image {
	img, err := png.Decode(bytes.NewReader(bg))
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func Boat() image.Image {
	img, err := png.Decode(bytes.NewReader(boat))
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func Boat2() image.Image {
	img, err := png.Decode(bytes.NewReader(boat2))
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func LoadMusic() io.Reader {
	f, err := ebitenutil.OpenFile("assets/music/main.mp3")
	if err != nil {
		log.Fatal("Error loading music")
	}
	return f
}
