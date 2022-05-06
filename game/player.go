package game

import "github.com/hajimehoshi/ebiten/v2"

type Player struct {
	turnSpeed float64
	MaxSpeed  float64
	image     *ebiten.Image
	speed     float64
	a         float64
	w         int
	h         int
	x         int
	y         int
}
