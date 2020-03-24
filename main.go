package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func main() {
	initializerSystem()

	err := ebiten.Run(update, canvasWidth, canvasHeight, 1, "Ebiten + ECS")
	if err != nil {
		log.Fatal(err)
	}
}

// Since ebiten runs in a constant 60 ticks per second, the delta can be
// calculated as a constant (in milliseconds).
const delta = 1 / 60.0 * 1000

func update(screen *ebiten.Image) error {
	movableSystem(delta)

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	rendererSystem(screen)

	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf("FPS: %0.2f", ebiten.CurrentFPS()),
	)
	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf("\nTPS: %0.2f", ebiten.CurrentTPS()),
	)

	return nil
}
