package application

import "goGravity/domain"

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
