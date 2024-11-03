package main

import (
	"fmt"
	"image/color"
	"time"

	"github.com/RoshanShrestha123/go-snake/food"
	"github.com/RoshanShrestha123/go-snake/snake"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

var player = snake.NewSnake()
var foods []food.Food

func (g *Game) Update() error {

	if !player.IsAlive {
		return nil
	}
	width, height := ebiten.WindowSize()
	player.UpdateDirection()
	player.CheckCollision(width, height)
	player.EatFood(&foods)
	player.UpdateMovement()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	if !player.IsAlive {
		screen.Fill(color.RGBA{R: 255, B: 0, G: 0})

	}
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Score: %d", player.Score))

	for _, body := range player.Segment {
		ebitenutil.DrawRect(screen, float64(body.X), float64(body.Y), float64(player.Size.W), float64(player.Size.H), color.White)
	}

	for _, food := range foods {
		ebitenutil.DrawRect(screen, float64(food.Position.X), float64(food.Position.Y), float64(player.Size.W), float64(player.Size.H), food.Color)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 420
}

func Timer() {

	for {
		if !player.IsAlive {
			break
		}
		time.Sleep(5 * time.Second)

		if len(foods) > 1 {
			foods = foods[1:]

		}
		foods = append(foods, *food.NewFood())
		fmt.Println("Generate new food", foods)

	}

}
