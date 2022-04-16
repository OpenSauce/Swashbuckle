package game

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	MAX_SPEED = 10.0
)

type Game struct {
	levelData LevelData
	GameData
}

func New() *Game {
	return &Game{
		levelData: CreateLevelOne(),

		GameData: GameData{
			tileSize:     64,
			ScreenWidth:  1920,
			ScreenHeight: 1080,
		},
	}
}

func (g *Game) Update() error {
	if g.levelData.p.speed > 0 {
		g.levelData.p.speed -= 0.05
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.levelData.p.a += 2.0 * math.Pi / 180
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.levelData.p.a -= 2.0 * math.Pi / 180
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		if g.levelData.p.speed < MAX_SPEED {
			g.levelData.p.speed += 0.2
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.levelData = CreateLevelTwo()
	}

	g.levelData.p.x -= int(g.levelData.p.speed * math.Sin(g.levelData.p.a))
	g.levelData.p.y += int(g.levelData.p.speed * math.Cos(g.levelData.p.a))

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.renderBackground(screen)
	g.renderPlayer(screen)
}

func (g *Game) renderBackground(screen *ebiten.Image) {
	for x := 0; x*g.tileSize <= g.ScreenWidth*5; x++ {
		for y := 0; y*g.tileSize <= g.ScreenHeight*5; y++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64((x*g.tileSize)-g.levelData.p.x), float64((y*g.tileSize)-g.levelData.p.y))
			screen.DrawImage(g.levelData.bg, op)
		}
	}
}

// Render
func (g *Game) renderPlayer(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(g.levelData.p.w)/2, -float64(g.levelData.p.h)/2)
	op.GeoM.Rotate(g.levelData.p.a)
	op.GeoM.Translate(
		float64(g.ScreenWidth)/2.0,
		float64(g.ScreenHeight)/2.0,
	)

	screen.DrawImage(g.levelData.p.image, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.ScreenWidth, g.ScreenHeight
}
