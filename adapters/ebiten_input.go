package adapters

import "time"

type EbitenInput struct {
	last time.Time
}

func NewEbitenInput() *EbitenInput {
	return &EbitenInput{time.Now()}
}

func (i *EbitenInput) GetDeltaTime() float64 {
	now := time.Now()
	dt := now.Sub(i.last).Seconds()
	i.last = now
	return dt
}
