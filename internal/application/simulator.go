package application

import (
	"goGravity/internal/domain"
)

type Simulator struct {
	Simulation   *domain.Simulation
	Renderer     Renderer
	InputHandler InputHandler
	TimeScale    float64
}

func (s *Simulator) Update() {
	dt := s.InputHandler.GetDeltaTime() * s.TimeScale
	s.Simulation.Step(dt)
}

func (s *Simulator) Draw() {
	s.Renderer.DrawBodies(s.Simulation.Bodies)
}
