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
		g.orientation++
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.orientation--
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

	w, h := g.boatImage.Size()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	op.GeoM.Rotate(float64(g.orientation%360) * 2 * math.Pi / 360)
	op.GeoM.Translate(
		float64(g.screenWidth)/2.0,
		float64(g.screenHeight)/2.0,
	)

	screen.DrawImage(g.boatImage, op)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}
