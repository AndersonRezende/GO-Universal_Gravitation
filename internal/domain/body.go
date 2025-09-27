package domain

import (
	"image/color"
	"math"
)

type Vector struct {
	X, Y float64
}

func (v Vector) Add(other Vector) Vector {
	return Vector{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

func (v Vector) Sub(other Vector) Vector {
	return Vector{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}

func (v Vector) Scale(scale float64) Vector {
	return Vector{
		X: v.X * scale,
		Y: v.Y * scale,
	}
}

func (v Vector) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type Body struct {
	Mass     float64
	Position Vector
	Velocity Vector
	Color    color.Color
	Name     string
	Size     float64
}

func NewBody(mass float64, position, velocity Vector, color color.Color, name string, size float64) *Body {
	return &Body{
		Mass:     mass,
		Position: position,
		Velocity: velocity,
		Color:    color,
		Name:     name,
		Size:     size,
	}
}

func (b *Body) ApplyForce(acceleration Vector, dt float64) {
	b.Velocity.X += acceleration.X * dt
	b.Velocity.Y += acceleration.Y * dt
	b.Position = PositionAfterTime(b.Position, b.Velocity, dt*1e-13)
}
