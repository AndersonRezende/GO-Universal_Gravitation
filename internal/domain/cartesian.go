package domain

import "math"

func ScalarDistance(vector Vector) float64 {
	return math.Sqrt(math.Pow(vector.X, 2) + math.Pow(vector.Y, 2))
}

func DistanceBetweenTwoBodies(a, b *Body) Vector {
	return Vector{
		X: b.Position.X - a.Position.X,
		Y: b.Position.Y - a.Position.Y,
	}
}

func PositionAfterTime(position, velocity Vector, dt float64) Vector {
	return Vector{
		X: position.X + velocity.X*dt,
		Y: position.Y + velocity.Y*dt,
	}
}
