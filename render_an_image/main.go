package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	_ "github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"bytes"
	_ "embed"
	"image"
)

//go:embed resources/images/gopher.png
var png []byte
var img *ebiten.Image

func init() {
	var err error
	//img, _, err = ebitenutil.NewImageFromFile("gopher.png")
	imgDecoded, _, err := image.Decode(bytes.NewReader(png))
	if err != nil {
		log.Fatal(err)
	}

	img = ebiten.NewImageFromImage(imgDecoded)
}

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(img, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Render an image")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
