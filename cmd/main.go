package main

import (
	"goGravity/internal/adapters"
	"goGravity/internal/application"
	"goGravity/internal/domain"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	simulator *application.Simulator
	renderer  *adapters.EbitenRenderer
}

func (g *Game) Update() error {
	g.simulator.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.renderer.SetScreen(screen)
	g.simulator.Draw()
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 1920, 1080
}

func main() {
	bodies := []*domain.Body{
		{Mass: 1.98892e30, Position: domain.Vector{950, 540}, Velocity: domain.Vector{0, 0}, Color: color.RGBA{A: 255, R: 255, G: 255, B: 0}, Name: "Sun", Size: 100},
		{Mass: 5.972e24, Position: domain.Vector{950, 15}, Velocity: domain.Vector{9e14, 0}, Color: color.RGBA{A: 255, B: 255}, Name: "Earth", Size: 10},
	}
	simulation := &domain.Simulation{Bodies: bodies}

	renderer := &adapters.EbitenRenderer{}
	input := adapters.NewEbitenInput()

	simulator := &application.Simulator{
		Simulation:   simulation,
		Renderer:     renderer,
		InputHandler: input,
		TimeScale:    1,
	}

	game := &Game{simulator: simulator, renderer: renderer}
	ebiten.SetWindowSize(1600, 900)
	ebiten.RunGame(game)
}
