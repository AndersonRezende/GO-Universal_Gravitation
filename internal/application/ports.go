package application

import (
	"goGravity/internal/domain"
)

type Renderer interface {
	DrawBodies(bodies []*domain.Body)
}

type InputHandler interface {
	GetDeltaTime() float64
}
