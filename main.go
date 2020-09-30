package main

import (
	tools "gfx/ebiten-tools"

	physics "github.com/goisgreat/go-physics"
	"github.com/hajimehoshi/ebiten"
)

func main() {
	// instantiate event stream
	doEventStream()

	// spawn sprites
	physics.Spawn(player)
	physics.Spawn(cube)

	// instantiate window object
	window := tools.Window{
		Width:  250,
		Height: 250,
		Scale:  2.5,
		Title:  "My game",
	}
	// game loop
	window.Loop(func(image *ebiten.Image) error {
		// clear screen by filling it with white
		image.Fill(white)

		// update & draw sprites
		physics.Spritelist.Update()
		physics.Spritelist.Draw(image)

		// return nil to indicate no error happened
		return nil
	})
}
