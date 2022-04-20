package game

import (
	"github.com/OpenSauce/Swashbuckle/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type LevelData struct {
	layout [][]MapTile
	p      Player
}

func CreateLevelOne() LevelData {
	layout := CreateLevelOneLayout()
	playerImage := ebiten.NewImageFromImage(assets.Boat())
	w, h := playerImage.Size()
	return LevelData{
		layout: layout,
		p: Player{
			image: playerImage,
			w:     w,
			h:     h,
			x:     1000,
			y:     1000,
		},
	}
}

func CreateLevelOneLayout() [][]MapTile {
	layout := [][]MapTile{}
	bgImage := ebiten.NewImageFromImage(assets.Background())
	for x := 0; x < 2000; x++ {
		layout = append(layout, []MapTile{})
		for y := 0; y < 2000; y++{
			layout[x] = append(layout[x], MapTile{
				img: bgImage,
			})
		}
	}

	return layout
}
