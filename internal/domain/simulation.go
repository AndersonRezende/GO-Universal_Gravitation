package domain

type Simulation struct {
	Bodies []*Body
}

func (s *Simulation) Step(dt float64) {
	for i, a := range s.Bodies {
		//totalForce := Vector{}
		for j, b := range s.Bodies {
			if i == j {
				continue
			}
			getGravitationalForceBetweenBodies(a, b, dt)
		}
	}
}

func getGravitationalForceBetweenBodies(a, b *Body, dt float64) {
	distance := DistanceBetweenTwoBodies(a, b)
	scalarDistance := ScalarDistance(distance)
	gravitationalForce := CalculateGravitationalForce(a, b, scalarDistance)

	normalizedVectorX := distance.X / scalarDistance
	normalizedVectorY := distance.Y / scalarDistance

	accelerationX := CalculateAcceleration(gravitationalForce, a.Mass) * normalizedVectorX
	accelerationY := CalculateAcceleration(gravitationalForce, a.Mass) * normalizedVectorY

	a.ApplyForce(Vector{X: accelerationX, Y: accelerationY}, dt)
}
