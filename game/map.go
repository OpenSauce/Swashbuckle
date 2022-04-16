package game

import "github.com/hajimehoshi/ebiten/v2"

type MapTile struct {
	x        int
	y        int
	blocking bool
	img      *ebiten.Image
}
