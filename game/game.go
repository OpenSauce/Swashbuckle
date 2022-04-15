package game

import (
	"math"

	"github.com/OpenSauce/Swashbuckle/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	bgImage,
	boatImage *ebiten.Image

	tileSize     int
	orientation  float64
	charX        int
	charY        int
	screenWidth  int
	screenHeight int
	viewport
}

type viewport struct {
	speed float64
	x     int
	y     int
}

func New() *Game {
	bgImage := ebiten.NewImageFromImage(assets.Background())
	boatImage := ebiten.NewImageFromImage(assets.Boat())

	return &Game{
		bgImage:   bgImage,
		boatImage: boatImage,

		tileSize:     64,
		charX:        66,
		charY:        113,
		screenWidth:  640,
		screenHeight: 480,
	}
}

func (g *Game) Update() error {
	if g.speed > 0 {
		g.speed -= 0.05
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.orientation += 2.0 * math.Pi / 180
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.orientation -= 2.0 * math.Pi / 180
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		if g.speed < 10 {
			g.speed += 0.2
		}
	}

	g.x -= int(float64(g.speed) * math.Sin(g.orientation))
	g.y += int(float64(g.speed) * math.Cos(g.orientation))

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.RenderBackground(screen)
	g.RenderPlayer(screen)
}

func (g *Game) RenderBackground(screen *ebiten.Image) {
	for x := 0; x*g.tileSize <= g.screenWidth*20; x++ {
		for y := 0; y*g.tileSize <= g.screenHeight*20; y++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64((x*g.tileSize)-g.x), float64((y*g.tileSize)-g.y))
			screen.DrawImage(g.bgImage, op)
		}
	}
}

func (g *Game) RenderPlayer(screen *ebiten.Image) {
	w, h := g.boatImage.Size()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	op.GeoM.Rotate(g.orientation)
	op.GeoM.Translate(
		float64(g.screenWidth)/2.0,
		float64(g.screenHeight)/2.0,
	)

	screen.DrawImage(g.boatImage, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}
