package food

import (
	"image/color"
	"math/rand"

	"github.com/RoshanShrestha123/go-snake/helper"
)

type Food struct {
	Position helper.VectorXY
	Size     helper.VectorWH
	Color    color.RGBA
}

func NewFood() *Food {
	return &Food{
		Position: helper.VectorXY{
			X: rand.Intn(400-20+1) + 20,
			Y: rand.Intn(400-20+1) + 20,
		},
		Size: helper.VectorWH{
			W: 20,
			H: 20,
		},
		Color: color.RGBA{R: 124, G: 252, B: 0},
	}
}
