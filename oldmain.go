package main

import (
	adapters2 "goGravity/internal/adapters"
	"goGravity/internal/application"
	domain2 "goGravity/internal/domain"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	simulator *application.Simulator
	renderer  *adapters2.EbitenRenderer
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
	return 800, 600
}

func main() {
	bodies := []*domain2.Body{
		{Mass: 5e10, Position: domain2.Vector{400, 300}, Velocity: domain2.Vector{0, 0}},
		{Mass: 1e5, Position: domain2.Vector{500, 300}, Velocity: domain2.Vector{0, 1}},
	}
	simulation := &domain2.Simulation{Bodies: bodies}

	renderer := &adapters2.EbitenRenderer{}
	input := adapters2.NewEbitenInput()

	simulator := &application.Simulator{
		Simulation:   simulation,
		Renderer:     renderer,
		InputHandler: input,
		TimeScale:    1,
	}

	game := &Game{simulator: simulator, renderer: renderer}
	ebiten.RunGame(game)
}
