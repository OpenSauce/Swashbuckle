package assets

import (
	"bytes"
	_ "embed"
	"image"
	"image/png"
	"log"
)

//go:embed textures/sea.png
var bg []byte

//go:embed char/boat.png
var boat []byte

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
