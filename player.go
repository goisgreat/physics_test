package main

import (
	physics "github.com/goisgreat/go-physics"
)

func player() physics.Sprite {
	// define a rectangle for the player sprite
	rect := physics.CreateRectangle(0, 0, 10, 10, green)
	// define a player sprite
	sprite := physics.CreateControllableShape(
		// use the rectangle we just defined
		rect,
		// use a KeyboardController returned by KeyboardControllerConfig.Init()
		physics.KeyboardControllerConfig{
			Input:    eventStream.Input,   // for the KeyboardController, use the keyboard input channel from eventStream
			Config:   physics.DIRECT_WASD, // use a directly controlled WASD keyboard controller preset
			Geometry: rect,                // use the rectangle from earlier here too
		}.Init(),
		// use a CollisonBox returned by CollisonBoxConfig.Init()
		physics.CollisionBoxConfig{
			Config: physics.RIGID_BODY,
		}.Init(),
	)
	// return player sprite
	return &sprite
}
