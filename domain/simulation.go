package domain

const G = 6.67430e-11

type Simulation struct {
	Bodies []Body
}

func (s *Simulation) Step(dt float64) {
	for i, a := range s.Bodies {
		totalForce := Vector{}
		for j, b := range s.Bodies {
			if i == j {
				continue
			}
			direction := b.Position.Sub(a.Position)
			dist := direction.Magnitude()
			if dist == 0 {
				continue
			}
			forceMagnitude := G * a.Mass * b.Mass / (dist * dist)
			force := direction.Scale(forceMagnitude / dist)
			totalForce = totalForce.Add(force)
		}
		a.ApplyForce(totalForce, dt)
	}
}
