package ebitenwrapper

import (
	"image/color"
	"log"

	ebiten "github.com/hajimehoshi/ebiten/v2"
)

type mediator interface {
	Update()
	Draw(screen *ebiten.Image)
	Exit()
}

type Game struct {
	Width    int
	Height   int
	Mediator mediator
	ebiten.Game
	WindowTitle     string
	WindowScale     int
	BackgroundColor color.Color
}

func (g *Game) Update() error {
	g.Mediator.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.BackgroundColor)
	g.Mediator.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth / g.WindowScale, outsideHeight / g.WindowScale
}

func (g *Game) Run() {
	ebiten.SetWindowSize(g.Width, g.Height)
	ebiten.SetWindowTitle(g.WindowTitle)
	ebiten.SetWindowResizable(true)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Exit() {
	g.Exit()
}
