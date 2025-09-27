package domain

import "math"

const G = 6.67430e-11

func CalculateGravitationalForce(a, b *Body, d float64) float64 {
	// |F| = (G * M * m) / d²
	f := (G * a.Mass * b.Mass) / math.Pow(d, 2)
	return math.Abs(f)
}

func CalculateAcceleration(f float64, m float64) float64 {
	// a = F / m
	return f / m
}

func CalculateVelocity(vt float64, a float64, dt float64) float64 {
	//vt+1=vt+a⋅Δt
	return vt + a*dt
}
