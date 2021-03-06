package game

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"time"

	"github.com/OpenSauce/Swashbuckle/assets"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	ACCELERATION = 0.2
	DECELERATION = 0.05
)

var (
	gameFont font.Face
	msg      chan struct{}
)

type Powerup struct {
	image       *ebiten.Image
	powerupType int
	speed       float64
	a           float64
	w           int
	h           int
	x           int
	y           int
}

type Game struct {
	levelData LevelData
	GameData
	powerups []Powerup
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

	audioContext := audio.NewContext(44100)
	msg = make(chan struct{})

	d, err := mp3.DecodeWithSampleRate(44100, assets.LoadMusic())
	if err != nil {
		log.Fatal("Error loading music")
	}

	p, err := audioContext.NewPlayer(d)
	if err != nil {
		log.Fatal("Error loading music")
	}

	p.SetVolume(0.2)

	p.Play()

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
		g.levelData.p.speed -= DECELERATION
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.levelData.p.a += g.levelData.p.turnSpeed * math.Pi / 180
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.levelData.p.a -= g.levelData.p.turnSpeed * math.Pi / 180
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		if g.levelData.p.speed < g.levelData.p.MaxSpeed {
			g.levelData.p.speed += ACCELERATION
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.levelData = CreateLevelTwo()
	}

	newXPos := g.levelData.p.x - int(g.levelData.p.speed*math.Sin(g.levelData.p.a))
	newYPos := g.levelData.p.y + int(g.levelData.p.speed*math.Cos(g.levelData.p.a))

	if g.levelData.layout[1][newXPos/g.tileSize][newYPos/g.tileSize].blocking {
		return nil
	}

	g.levelData.p.x = newXPos
	g.levelData.p.y = newYPos

	for i, powerup := range g.levelData.powerup {
		if g.levelData.p.x < powerup.x+powerup.w &&
			g.levelData.p.x+g.levelData.p.w > powerup.x &&
			g.levelData.p.y < powerup.y+powerup.h &&
			g.levelData.p.h+g.levelData.p.y > powerup.y {
			g.levelData.powerup = append(g.levelData.powerup[:i], g.levelData.powerup[i+1:]...)
			switch powerup.powerupType {
			case 0:
				g.levelData.p.turnSpeed = 10.0
			case 1:
				g.levelData.p.MaxSpeed = 20.0
			}

			go func() {
				time.Sleep(5 * time.Second)
				msg <- struct{}{}
			}()
		}
	}

	select {
	case <-msg:
		g.levelData.p.turnSpeed = 2.0
		g.levelData.p.MaxSpeed = 20.0
	default:
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.renderBackground(screen)
	g.renderPlayer(screen)
	g.renderMisc(screen)
}

func (g *Game) renderMisc(screen *ebiten.Image) {
	startingXPos := g.levelData.p.x - g.ScreenWidth/2
	startingYPos := g.levelData.p.y - g.ScreenHeight/2
	for _, powerup := range g.levelData.powerup {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(powerup.x-startingXPos), float64(powerup.y-startingYPos))
		screen.DrawImage(powerup.image, op)
	}
}

func (g *Game) renderBackground(screen *ebiten.Image) {
	startingXPos := g.levelData.p.x - g.ScreenWidth/2
	startingYPos := g.levelData.p.y - g.ScreenHeight/2
	tileX := startingXPos / g.tileSize
	tileY := startingYPos / g.tileSize

	for l := 0; l < len(g.levelData.layout); l++ {
		for x := 0; x*g.tileSize <= g.ScreenWidth; x++ {
			for y := 0; y*g.tileSize <= g.ScreenHeight+g.tileSize; y++ {
				tileXI := tileX + x
				tileYI := tileY + y

				if tileXI < 0 || tileYI < 0 || tileXI >= len(g.levelData.layout[l]) || tileYI >= len(g.levelData.layout[l][tileXI]) {
					continue
				}

				tile := g.levelData.layout[l][tileXI][tileYI]

				if tile.img == nil {
					continue
				}
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x*g.tileSize-startingXPos%g.tileSize), float64(y*g.tileSize-startingYPos%g.tileSize))
				screen.DrawImage(tile.img, op)
				text.Draw(screen, fmt.Sprintf("Pos: %v %v Start: %v %v Tile: %v %v",
					g.levelData.p.x, g.levelData.p.y, startingXPos, startingYPos, tileX, tileY), gameFont, 20, 20, color.Black)
			}
		}
	}

	text.Draw(screen, fmt.Sprintf("%f", ebiten.CurrentFPS()), gameFont, g.ScreenWidth-200, 20, color.Black)

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
