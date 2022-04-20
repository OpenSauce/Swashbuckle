package game

import (
	"fmt"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	MAX_SPEED = 10.0
)

var (
	gameFont font.Face
)

type Game struct {
	levelData LevelData
	GameData
}

func New() *Game {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	gameFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

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
	startingXPos := g.levelData.p.x - g.ScreenWidth/2
	startingYPos := g.levelData.p.y - g.ScreenHeight/2
	tileX := startingXPos / g.tileSize
	tileY := startingYPos / g.tileSize

	for x := 0; x*g.tileSize <= g.ScreenWidth; x++ {
		for y := 0; y*g.tileSize <= g.ScreenHeight+g.tileSize; y++ {
			tileXI := tileX + x
			tileYI := tileY + y

			if tileXI < 0 || tileYI < 0 || tileXI >= len(g.levelData.layout) || tileYI >= len(g.levelData.layout[0]) {
				continue
			}

			tile := g.levelData.layout[tileXI][tileYI]
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x*g.tileSize-startingXPos%g.tileSize), float64(y*g.tileSize-startingYPos%g.tileSize))
			screen.DrawImage(tile.img, op)
			text.Draw(screen, fmt.Sprintf("Pos: %v %v Start: %v %v Tile: %v %v",
				g.levelData.p.x, g.levelData.p.y, startingXPos, startingYPos, tileX, tileY), gameFont, 20, 20, color.White)
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
