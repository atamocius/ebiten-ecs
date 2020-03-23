package main

import "math/rand"

const speedMultiplier = 0.1

func getRandomVelocity() velocity {
	return velocity{
		vX: speedMultiplier * (2*rand.Float32() - 1),
		vY: speedMultiplier * (2*rand.Float32() - 1),
	}
}

func getRandomPosition() position {
	return position{
		X: rand.Float32() * canvasWidth,
		Y: rand.Float32() * canvasHeight,
	}
}

func getRandomShape() shape {
	if rand.Float32() >= 0.5 {
		return shape{
			Primitive: circle,
		}
	}

	return shape{
		Primitive: box,
	}
}
