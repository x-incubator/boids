package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/x-incubator/boids/boids"
)

func main() {
	// pre-fill map array with -1
	for i, row := range boids.BoidMap {
		for j := range row {
			boids.BoidMap[i][j] = -1
		}
	}

	ebiten.SetWindowTitle("BOIDS simulation")
	ebiten.SetWindowSize(boids.ScreenWidth*2, boids.ScreenHeight*2)

	for i := 0; i < boids.BoidCount; i++ {
		boids.CreateBoid(i)
	}

	if err := ebiten.RunGame(&boids.Game{}); err != nil {
		log.Fatal(err)
	}
}
