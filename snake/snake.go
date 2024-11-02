package snake

import (
	"github.com/RoshanShrestha123/go-snake/helper"
	"github.com/hajimehoshi/ebiten/v2"
)

type Snake struct {
	Size      helper.VectorWH
	Segment   []helper.VectorXY
	IsAlive   bool
	Chance    int
	Speed     int
	Direction helper.VectorXY
	Grow      bool
	Score     int
}

func NewSnake() *Snake {
	return &Snake{
		Size:      helper.VectorWH{W: 15, H: 15},
		IsAlive:   true,
		Chance:    5,
		Speed:     2,
		Grow:      false,
		Direction: helper.VectorXY{X: 1, Y: 0},
		Score:     0,
		Segment: []helper.VectorXY{
			{X: 10, Y: 10},
			{X: 10, Y: 11},
			{X: 10, Y: 12},
			{X: 10, Y: 13},
			{X: 10, Y: 14},
			{X: 10, Y: 15},
			{X: 10, Y: 16},
			{X: 10, Y: 17},
			{X: 10, Y: 18},
			{X: 10, Y: 19},
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
		if s.Direction.Y != -1 {
			s.Direction.Y = 1
			s.Direction.X = 0
		}

	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {

		if s.Direction.Y != 1 {
			s.Direction.Y = -1
			s.Direction.X = 0
		}

	}

}

func (s *Snake) CheckCollision(screenWidth, screenHeight int) {
	head := s.Segment[0]

	if head.X+s.Size.W >= screenWidth || head.X <= 0 || head.Y+s.Size.H >= screenHeight || head.Y <= 0 {
		s.IsAlive = false
	}

}

func (s *Snake) UpdateMovement() {
	if !s.IsAlive {
		return
	}
	newSegment := helper.VectorXY{
		X: s.Segment[0].X + (s.Direction.X * s.Speed),
		Y: s.Segment[0].Y + (s.Direction.Y * s.Speed),
	}

	s.Segment = append([]helper.VectorXY{newSegment}, s.Segment...)

	if !s.Grow {
		s.Segment = s.Segment[:len(s.Segment)-1]
	}

	s.Grow = false

}
