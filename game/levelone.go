package game

import (
	"github.com/OpenSauce/Swashbuckle/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type LevelData struct {
	layout [][]MapTile
	p      Player
}

type MapTile struct {
	x        int
	y        int
	blocking bool
	img      *ebiten.Image
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
	layout := GenerateGrid(2000, 2000)
	return layout
}

func GenerateGrid(maxX, maxY int) [][]MapTile {
	layout := [][]MapTile{}
	bgImage := ebiten.NewImageFromImage(assets.Background())
	topLeft := ebiten.NewImageFromImage(assets.TopLeftWall())
	topRight := ebiten.NewImageFromImage(assets.TopRightWall())
	bottomLeft := ebiten.NewImageFromImage(assets.BottomLeftWall())
	bottomRight := ebiten.NewImageFromImage(assets.BottomRightWall())
	XWall := ebiten.NewImageFromImage(assets.XWall())
	YWall := ebiten.NewImageFromImage(assets.YWall())
	for x := 0; x < maxX; x++ {
		layout = append(layout, []MapTile{})
		for y := 0; y < maxY; y++ {
			if x == 0 && y == 0 {
				layout[x] = append(layout[x], MapTile{
					blocking: true,
					img:      topLeft,
				})
				continue
			}

			if x == 0 && y == maxX-1 {
				layout[x] = append(layout[x], MapTile{
					blocking: true,
					img:      topRight,
				})
				continue
			}

			if x == maxX-1 && y == 0 {
				layout[x] = append(layout[x], MapTile{
					blocking: true,
					img:      bottomLeft,
				})
				continue
			}

			if x == maxX-1 && y == maxX-1 {
				layout[x] = append(layout[x], MapTile{
					blocking: true,
					img:      bottomRight,
				})
				continue
			}

			if x == 0 || x == maxX-1 {
				layout[x] = append(layout[x], MapTile{
					blocking: true,
					img:      YWall,
				})
				continue
			}

			if y == 0 || y == maxX-1 {
				layout[x] = append(layout[x], MapTile{
					blocking: true,
					img:      XWall,
				})
				continue
			}

			layout[x] = append(layout[x], MapTile{
				img: bgImage,
			})

		}

	}
	return layout
}
