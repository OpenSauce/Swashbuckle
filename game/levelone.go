package game

import (
	"github.com/OpenSauce/Swashbuckle/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type LevelData struct {
	layout  [][][]MapTile
	powerup []Powerup
	p       Player
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
	turnPowerup := ebiten.NewImageFromImage(assets.TurnPowerup())

	w, h := turnPowerup.Size()
	var powerups []Powerup
	for i := 0; i < 3; i++ {
		powerups = append(powerups, Powerup{
			image: turnPowerup,
			x:     1200 * i,
			y:     1200 * i,
			w:     w,
			h:     h,
		})
	}

	w, h = playerImage.Size()
	return LevelData{
		layout:  layout,
		powerup: powerups,
		p: Player{
			turnSpeed: 2.0,
			image:     playerImage,
			w:         w,
			h:         h,
			x:         1000,
			y:         1000,
		},
	}
}

func CreateLevelOneLayout() [][][]MapTile {
	layout := GenerateGrid(2000, 2000)
	return layout
}

func GenerateGrid(maxX, maxY int) [][][]MapTile {
	layout := [][][]MapTile{}
	bgImage := ebiten.NewImageFromImage(assets.Background())
	topLeft := ebiten.NewImageFromImage(assets.TopLeftWall())
	topRight := ebiten.NewImageFromImage(assets.TopRightWall())
	bottomLeft := ebiten.NewImageFromImage(assets.BottomLeftWall())
	bottomRight := ebiten.NewImageFromImage(assets.BottomRightWall())
	XWall := ebiten.NewImageFromImage(assets.XWall())
	YWall := ebiten.NewImageFromImage(assets.YWall())

	layout = append(layout, [][]MapTile{})

	for x := 0; x < maxX; x++ {
		layout[0] = append(layout[0], []MapTile{})
		for y := 0; y < maxY; y++ {
			layout[0][x] = append(layout[0][x], MapTile{
				img: bgImage,
			})
		}
	}

	layout = append(layout, [][]MapTile{})

	for x := 0; x < maxX; x++ {
		layout[1] = append(layout[1], []MapTile{})
		for y := 0; y < maxY; y++ {
			if x == 0 && y == 0 {
				layout[1][x] = append(layout[1][x], MapTile{
					blocking: true,
					img:      topLeft,
				})
				continue
			}

			if x == 0 && y == maxX-1 {
				layout[1][x] = append(layout[1][x], MapTile{
					blocking: true,
					img:      topRight,
				})
				continue
			}

			if x == maxX-1 && y == 0 {
				layout[1][x] = append(layout[1][x], MapTile{
					blocking: true,
					img:      bottomLeft,
				})
				continue
			}

			if x == maxX-1 && y == maxX-1 {
				layout[1][x] = append(layout[1][x], MapTile{
					blocking: true,
					img:      bottomRight,
				})
				continue
			}

			if x == 0 || x == maxX-1 {
				layout[1][x] = append(layout[1][x], MapTile{
					blocking: true,
					img:      YWall,
				})
				continue
			}

			if y == 0 || y == maxX-1 {
				layout[1][x] = append(layout[1][x], MapTile{
					blocking: true,
					img:      XWall,
				})
				continue
			}

			layout[1][x] = append(layout[1][x], MapTile{})

		}

	}
	return layout
}
