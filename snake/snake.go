package snake

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type VectorXY struct {
	X, Y int
}

type VectorWH struct {
	W, H int
}

type Snake struct {
	Size      VectorWH
	Segment   []VectorXY
	IsAlive   bool
	Chance    int
	Speed     int
	Direction VectorXY
	Grow      bool
}

func NewSnake() *Snake {
	return &Snake{
		Size:      VectorWH{W: 10, H: 10},
		IsAlive:   true,
		Chance:    5,
		Speed:     2,
		Grow:      false,
		Direction: VectorXY{X: 1, Y: 0},
		Segment: []VectorXY{
			{X: 10, Y: 10},
			{X: 10, Y: 11},
			{X: 10, Y: 12},
			{X: 10, Y: 13},
			{X: 10, Y: 14},
			{X: 10, Y: 15},
		},
	}
}

func (s *Snake) UpdateDirection() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		if s.Direction.X != 1 {
			s.Direction.X = -1
			s.Direction.Y = 0
		}

	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		if s.Direction.X != -1 {
			s.Direction.X = 1
			s.Direction.Y = 0
		}

	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		s.Direction.Y = 1
		s.Direction.X = 0
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		s.Direction.Y = -1
		s.Direction.X = 0
	}

}

func (s *Snake) UpdateMovement() {
	newSegment := VectorXY{
		X: s.Segment[0].X + (s.Direction.X * s.Speed),
		Y: s.Segment[0].Y + (s.Direction.Y * s.Speed),
	}

	s.Segment = append([]VectorXY{newSegment}, s.Segment...)

	if !s.Grow {
		s.Segment = s.Segment[:len(s.Segment)-1]
		s.Grow = false

	}

	s.Grow = false

}
