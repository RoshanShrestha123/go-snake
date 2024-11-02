package main

import (
	"fmt"
	"image/color"
	"time"

	"github.com/RoshanShrestha123/go-snake/snake"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

var player = snake.NewSnake()

func (g *Game) Update() error {
	player.UpdateDirection()
	player.UpdateMovement()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello world")

	for _, body := range player.Segment {
		ebitenutil.DrawRect(screen, float64(body.X), float64(body.Y), float64(player.Size.W), float64(player.Size.H), color.White)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 420
}

func Timer() {

	for {
		fmt.Println("simulating level up")
		time.Sleep(3 * time.Second)
		player.Grow = true

	}

}
