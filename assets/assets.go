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

//go:embed images/textures/walls/topleftwall.png
var topLeftWall []byte

//go:embed images/textures/walls/toprightwall.png
var topRightWall []byte

//go:embed images/textures/walls/bottomrightwall.png
var bottomRightWall []byte

//go:embed images/textures/walls/bottomleftwall.png
var bottomLeftWall []byte

//go:embed images/textures/walls/xwall.png
var xWall []byte

//go:embed images/textures/walls/ywall.png
var yWall []byte

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

func TopLeftWall() image.Image {
	img, err := png.Decode(bytes.NewReader(topLeftWall))
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func TopRightWall() image.Image {
	img, err := png.Decode(bytes.NewReader(topRightWall))
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func BottomLeftWall() image.Image {
	img, err := png.Decode(bytes.NewReader(bottomLeftWall))
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func BottomRightWall() image.Image {
	img, err := png.Decode(bytes.NewReader(bottomRightWall))
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func XWall() image.Image {
	img, err := png.Decode(bytes.NewReader(xWall))
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func YWall() image.Image {
	img, err := png.Decode(bytes.NewReader(yWall))
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
