package game

import (
	"github.com/OpenSauce/Swashbuckle/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	bgImage,
	boatImage *ebiten.Image

	tileSize     int
	orientation  int
	charX        int
	charY        int
	screenWidth  int
	screenHeight int
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
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.orientation = 3
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.orientation = 1
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for x := 0; x*g.tileSize <= g.screenWidth; x++ {
		for y := 0; y*g.tileSize <= g.screenHeight; y++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x*g.tileSize), float64(y*g.tileSize))
			screen.DrawImage(g.bgImage, op)
		}
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Rotate(float64(g.orientation) * 90)
	op.GeoM.Translate(
		float64(g.screenWidth)/2.0-float64(g.charX)/2.0,
		float64(g.screenHeight)/2.0-float64(g.charY)/2.0,
	)

	screen.DrawImage(g.boatImage, op)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}
