package game

import (
	"github.com/OpenSauce/Swashbuckle/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

func CreateLevelOne() LevelData {
	layout := CreateLevelOneLayout()
	playerImage := ebiten.NewImageFromImage(assets.Boat())
	bgImage := ebiten.NewImageFromImage(assets.Background())
	w, h := playerImage.Size()
	return LevelData{
		layout: layout,
		bg:     bgImage,
		p: Player{
			image: playerImage,
			w:     w,
			h:     h,
			x:     0,
			y:     0,
		},
	}
}

func CreateLevelOneLayout() []MapTile {
	return []MapTile{}
}
