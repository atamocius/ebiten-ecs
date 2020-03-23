package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/vector"
)

const shapeSize = 20
const shapeHalfSize = shapeSize / 2

const tau = float32(2 * math.Pi)
const circleStep = 20.0 * math.Pi / 180.0

var boxOp *vector.DrawPathOptions = &vector.DrawPathOptions{
	StrokeColor: color.RGBA{0xb7, 0x48, 0x43, 0xff},
	LineWidth:   2,
}

var circleOp *vector.DrawPathOptions = &vector.DrawPathOptions{
	StrokeColor: color.RGBA{0x0b, 0x84, 0x5b, 0xff},
	LineWidth:   2,
}

func drawBox(screen *ebiten.Image, position position) {
	var path vector.Path

	left := position.X - shapeHalfSize
	top := position.Y - shapeHalfSize
	right := position.X + shapeHalfSize
	bottom := position.Y + shapeHalfSize

	path.MoveTo(left, top)
	path.LineTo(right, top)
	path.LineTo(right, bottom)
	path.LineTo(left, bottom)
	path.LineTo(left, top)

	path.Draw(screen, boxOp)
}

func drawCircle(screen *ebiten.Image, position position) {
	var path vector.Path

	radius := float32(shapeHalfSize)
	theta := 0.0

	path.MoveTo(
		position.X+radius*float32(math.Cos(theta)),
		position.Y+radius*float32(math.Sin(theta)),
	)
	theta += circleStep

	for float32(theta) <= tau {
		x := position.X + radius*float32(math.Cos(theta))
		y := position.Y + radius*float32(math.Sin(theta))
		path.LineTo(x, y)
		theta += circleStep
	}

	path.Draw(screen, circleOp)
}
