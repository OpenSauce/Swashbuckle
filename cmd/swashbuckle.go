package main

import (
	"github.com/OpenSauce/Swashbuckle/game"
	"github.com/hajimehoshi/ebiten/v2"
)

const gameName = "Swashbuckle"

func main() {
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle(gameName)
	game := game.New()
	game.ScreenWidth, game.ScreenHeight = ebiten.ScreenSizeInFullscreen()
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
