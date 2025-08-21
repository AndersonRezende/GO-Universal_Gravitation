package domain

import "math"

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
}

func (b Body) ApplyForce(f Vector, dt float64) {
	acceleration := f.Scale(1.0 / b.Mass)
	b.Velocity = b.Velocity.Add(acceleration.Scale(dt))
	b.Position = b.Position.Add(b.Velocity.Scale(dt))
}
