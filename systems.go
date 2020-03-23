package main

import "github.com/hajimehoshi/ebiten"

const numElements = 600

var world struct {
	gameObjects []gameObject
}

func initializerSystem() {
	world.gameObjects = make([]gameObject, 0, numElements)

	for i := 0; i < numElements; i++ {
		e := gameObject{
			velocity:   getRandomVelocity(),
			shape:      getRandomShape(),
			position:   getRandomPosition(),
			renderable: renderable{true},
		}

		world.gameObjects = append(world.gameObjects, e)
	}
}

// Since ebiten runs in a constant 60 ticks per second, the delta can be
// calculated as a constant (in milliseconds).
const delta = 1 / 60.0 * 1000

func movableSystem() {
	count := len(world.gameObjects)

	for i := 0; i < count; i++ {
		e := &world.gameObjects[i]

		vel := e.velocity
		pos := e.position

		pos.X += vel.vX * delta
		pos.Y += vel.vY * delta

		if pos.X > canvasWidth+shapeHalfSize {
			pos.X = -shapeHalfSize
		}
		if pos.X < -shapeHalfSize {
			pos.X = canvasWidth + shapeHalfSize
		}
		if pos.Y > canvasHeight+shapeHalfSize {
			pos.Y = -shapeHalfSize
		}
		if pos.Y < -shapeHalfSize {
			pos.Y = canvasHeight + shapeHalfSize
		}

		e.velocity = vel
		e.position = pos
	}
}

func rendererSystem(screen *ebiten.Image) {
	for _, e := range world.gameObjects {
		if e.shape.Primitive == box {
			drawBox(screen, e.position)
		} else {
			drawCircle(screen, e.position)
		}
	}
}
