package boids

import (
	"fmt"
	"image/color"
	"os"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth, ScreenHeight = 320, 240
	BoidCount                 = 420
	ViewRadius                = 9
	AdjRate                   = 0.025
)

var (
	Cyan    = color.RGBA{106, 214, 227, 255}
	Black   = color.RGBA{0, 0, 0, 255}
	Boids   [BoidCount]*Boid
	BoidMap [ScreenWidth + 1][ScreenHeight + 1]int
	RWLock  = sync.RWMutex{}
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
	screen.Fill(Black)
	for _, boid := range Boids {
		screen.Set(int(boid.Position.X-1), int(boid.Position.Y), Cyan)
		screen.Set(int(boid.Position.X-2), int(boid.Position.Y), Cyan)
		screen.Set(int(boid.Position.X), int(boid.Position.Y), Cyan)
		screen.Set(int(boid.Position.X), int(boid.Position.Y-1), Cyan)
		screen.Set(int(boid.Position.X), int(boid.Position.Y-2), Cyan)
	}
}

func (g *Game) Layout(_, _ int) (w, h int) {
	return ScreenWidth, ScreenHeight
}
