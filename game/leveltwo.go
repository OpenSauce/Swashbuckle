package game

import (
	"github.com/OpenSauce/Swashbuckle/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type LevelData struct {
	layout []MapTile
	bg     *ebiten.Image
	p      Player
}

func CreateLevelTwo() LevelData {
	layout := CreateLevelOneLayout()
	playerImage := ebiten.NewImageFromImage(assets.Boat2())
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

func CreateLevelTwoLayout() []MapTile {
	return []MapTile{}
}
