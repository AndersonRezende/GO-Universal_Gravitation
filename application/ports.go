package application

import "goGravity/domain"

type Renderer interface {
	DrawBodies(bodies []*domain.Body)
}

type InputHandler interface {
	GetDeltaTime() float64
}
