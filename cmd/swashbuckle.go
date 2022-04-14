package main

import (
	"github.com/OpenSauce/Swashbuckle/game"
	"github.com/hajimehoshi/ebiten/v2"
)

const gameName = "Swashbuckle"

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle(gameName)
	game := game.New()
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
