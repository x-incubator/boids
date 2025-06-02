package main

import (
	"fmt"
	"image/color"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth, screenHeight = 320, 240
	boidCount                 = 69
)

var (
	white = color.White
	boids [boidCount]*Boid
)

type Game struct{}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		fmt.Println("Bye!")
		os.Exit(0)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, boid := range boids {
		screen.Set(int(boid.position.x+1), int(boid.position.y), white)
		screen.Set(int(boid.position.x-1), int(boid.position.y), white)

		screen.Set(int(boid.position.x), int(boid.position.y+1), white)
		screen.Set(int(boid.position.x), int(boid.position.y-1), white)
	}
}

func (g *Game) Layout(_, _ int) (w, h int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowTitle("BOIDS simulation")
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)

	for i := 0; i < boidCount; i++ {
		createBoid(i)
	}

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
