package game

import (
	"math"

	"github.com/OpenSauce/Swashbuckle/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	MAX_SPEED = 10.0
)

type Game struct {
	bgImage,
	boatImage *ebiten.Image

	p Player
	GameData
}

func New() *Game {
	bgImage := ebiten.NewImageFromImage(assets.Background())
	boatImage := ebiten.NewImageFromImage(assets.Boat())

	w, h := boatImage.Size()

	return &Game{
		bgImage:   bgImage,
		boatImage: boatImage,

		p: Player{
			w: w,
			h: h,
			x: 0,
			y: 0,
		},

		GameData: GameData{
			tileSize:     64,
			ScreenWidth:  1920,
			ScreenHeight: 1080,
		},
	}
}

func (g *Game) Update() error {
	if g.p.speed > 0 {
		g.p.speed -= 0.05
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.p.a += 2.0 * math.Pi / 180
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.p.a -= 2.0 * math.Pi / 180
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		if g.p.speed < MAX_SPEED {
			g.p.speed += 0.2
		}
	}

	g.p.x -= int(g.p.speed * math.Sin(g.p.a))
	g.p.y += int(g.p.speed * math.Cos(g.p.a))

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
			op.GeoM.Translate(float64((x*g.tileSize)-g.p.x), float64((y*g.tileSize)-g.p.y))
			screen.DrawImage(g.bgImage, op)
		}
	}
}

// Render
func (g *Game) renderPlayer(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(g.p.w)/2, -float64(g.p.h)/2)
	op.GeoM.Rotate(g.p.a)
	op.GeoM.Translate(
		float64(g.ScreenWidth)/2.0,
		float64(g.ScreenHeight)/2.0,
	)

	screen.DrawImage(g.boatImage, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.ScreenWidth, g.ScreenHeight
}
