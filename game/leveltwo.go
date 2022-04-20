package game

import (
	"github.com/OpenSauce/Swashbuckle/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

func CreateLevelTwo() LevelData {
	layout := CreateLevelOneLayout()
	playerImage := ebiten.NewImageFromImage(assets.Boat2())
	w, h := playerImage.Size()
	return LevelData{
		layout: layout,
		p: Player{
			image: playerImage,
			w:     w,
			h:     h,
			x:     0,
			y:     0,
		},
	}
}

func CreateLevelTwoLayout() []MapTile {
	return []MapTile{}
}
