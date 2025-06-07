package boids

import (
	"math"
	"math/rand/v2"
	"time"
)

type Boid struct {
	Position Vector2D
	Velocity Vector2D
	ID       int
}

func (b *Boid) calculateAcceleration() Vector2D {
	upper, lower := b.Position.AddV(ViewRadius), b.Position.SubtractV(ViewRadius)
	avgPosition, avgVelocity, separation := Vector2D{
		X: 0,
		Y: 0,
	}, Vector2D{
		X: 0,
		Y: 0,
	}, Vector2D{
		X: 0,
		Y: 0,
	}

	count := 0.0

	RWLock.RLock()
	for i := math.Max(lower.X, 0); i <= math.Min(upper.X, ScreenWidth); i++ {
		for j := math.Max(lower.Y, 0); j <= math.Min(upper.Y, ScreenHeight); j++ {
			if otherBoidId := BoidMap[int(i)][int(j)]; otherBoidId != -1 && otherBoidId != b.ID {
				if dist := Boids[otherBoidId].Position.Distance(b.Position); dist < ViewRadius {
					count++
					avgVelocity = avgVelocity.Add(Boids[otherBoidId].Velocity)
					avgPosition = avgPosition.Add(Boids[otherBoidId].Position)
					separation = separation.Add(
						b.Position.Subtract(Boids[otherBoidId].Position).DivisionV(dist),
					)
				}
			}
		}
	}
	RWLock.RUnlock()

	accel := Vector2D{
		X: b.borderBounce(b.Position.X, ScreenWidth),
		Y: b.borderBounce(b.Position.Y, ScreenHeight),
	}

	if count > 0 {
		avgPosition, avgVelocity = avgPosition.DivisionV(count), avgVelocity.DivisionV(count)
		accelAlignment := avgVelocity.Subtract(b.Velocity).MultiplyV(AdjRate)
		accelCohesion := avgPosition.Subtract(b.Position).MultiplyV(AdjRate)
		accelSeparation := separation.MultiplyV(AdjRate)
		accel = accel.Add(accelAlignment).Add(accelCohesion).Add(accelSeparation)
	}

	return accel
}

func (b *Boid) borderBounce(pos, maxBorderPos float64) float64 {
	if pos < ViewRadius {
		return 1 / pos
	} else if pos > maxBorderPos-ViewRadius {
		return 1 / (pos - maxBorderPos)
	}
	return 0
}

func (b *Boid) moveOne() {
	acceleration := b.calculateAcceleration()

	RWLock.Lock()
	b.Velocity = b.Velocity.Add(acceleration).Limit(-1, 1)
	BoidMap[int(b.Position.X)][int(b.Position.Y)] = -1
	b.Position = b.Position.Add(b.Velocity)
	BoidMap[int(b.Position.X)][int(b.Position.Y)] = b.ID
	RWLock.Unlock()
}

func (b *Boid) start() {
	for {
		b.moveOne()
		time.Sleep(5 * time.Millisecond)
	}
}

func CreateBoid(ID int) {
	b := Boid{
		Position: Vector2D{
			X: (rand.Float64() * ScreenWidth),
			Y: (rand.Float64() * ScreenHeight),
		},
		Velocity: Vector2D{
			X: (rand.Float64() * 2) - 1.0,
			Y: (rand.Float64() * 2) - 1.0,
		},
		ID: ID,
	}

	Boids[ID] = &b
	BoidMap[int(b.Position.X)][int(b.Position.Y)] = b.ID

	go b.start()
}
