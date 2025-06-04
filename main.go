package main

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth, screenHeight = 320, 240
	boidCount                 = 420
	viewRadius                = 13
	adjRate                   = 0.015
)

var (
	cyan    = color.RGBA{106, 214, 227, 255}
	black   = color.RGBA{0, 0, 0, 255}
	boids   [boidCount]*Boid
	boidMap [screenWidth + 1][screenHeight + 1]int
	lock    = sync.Mutex{}
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
	screen.Fill(black)
	for _, boid := range boids {
		vector.DrawFilledCircle(
			screen,
			float32(boid.position.x),
			float32(boid.position.y+1),
			4,
			cyan,
			true,
		)
	}
}

func (g *Game) Layout(_, _ int) (w, h int) {
	return screenWidth, screenHeight
}

func main() {
	// pre-fill map array with -1
	for i, row := range boidMap {
		for j := range row {
			boidMap[i][j] = -1
		}
	}

	ebiten.SetWindowTitle("BOIDS simulation")
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)

	for i := 0; i < boidCount; i++ {
		createBoid(i)
	}

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
