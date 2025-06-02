package main

import (
	"math/rand/v2"
	"time"
)

type Boid struct {
	position Vector2D
	velocity Vector2D
	id       int
}

func (b *Boid) moveOne() {
	b.position = b.position.Add(b.velocity)
	next := b.position.Add(b.velocity)

	if next.x >= screenWidth || next.x < 0 {
		b.velocity = Vector2D{
			x: -b.velocity.x,
			y: b.velocity.y,
		}
	}

	if next.y >= screenHeight || next.y < 0 {
		b.velocity = Vector2D{
			x: b.velocity.x,
			y: -b.velocity.y,
		}
	}
}

func (b *Boid) start() {
	for {
		b.moveOne()
		time.Sleep(5 * time.Millisecond)
	}
}

func createBoid(id int) {
	b := Boid{
		position: Vector2D{
			x: (rand.Float64() * screenWidth),
			y: (rand.Float64() * screenHeight),
		},
		velocity: Vector2D{
			x: (rand.Float64() * 2) - 1.0,
			y: (rand.Float64() * 2) - 1.0,
		},
		id: id,
	}

	boids[id] = &b

	go b.start()
}
