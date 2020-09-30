package main

import physics "github.com/goisgreat/go-physics"

func cube() physics.Sprite {
	// define a rectangle
	rect := physics.CreateRectangle(10, 10, 20, 20, green)
	// create a StaticShape
	sprite := physics.CreateStaticShape(
		// use rect defined earlier
		rect,
		// use a CollisionBox returned from CollisionBoxConfig.Init()
		physics.CollisionBoxConfig{
			Config: physics.RIGID_BODY,
		}.Init(),
	)
	// return sprite
	return &sprite
}
