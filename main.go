package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(640, 420)
	ebiten.SetWindowTitle("Snake game")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal("can't run game", err)
	}
}
