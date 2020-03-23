# ebiten-ecs

A simple ECS (Entity Component System) pattern implemented in [_Go_](https://go.dev/) and using [_Ebiten_](https://ebiten.org/) as the renderer.

> The code sample was inspired by [_ecsy_](https://ecsy.io/docs)'s documentation example.

![screenshot-ecs](https://user-images.githubusercontent.com/6222358/77348054-b0c85c00-6d30-11ea-800b-f2e34ebf41bd.png)

## Implementation

The idea is to implement ECS as a _pattern_ rather than a set of tools or a library/package.

### Components

Components are implemented as structs.

```go
type velocity struct {
	VX float32
	VY float32
}

type position struct {
	X float32
	Y float32
}

type shape struct {
	Primitive primitive
}

type renderable struct {
	Flag bool
}
```

It is important that these structs are composed of _value types_ (no reference or pointer types) to preserve [_data locality_](http://gameprogrammingpatterns.com/data-locality.html) as much as possible (no [_pointer chasing_](http://gameprogrammingpatterns.com/images/data-locality-pointer-chasing.png)).

### Entities

Entities are implemented as structs that _embed_ component structs. These essentially act as _archetypes_ (a definition of an entity's shape).

```go
type gameObject struct {
	velocity
	shape
	position
	renderable
}
```

Archetypes are created at _design-time_. As a limitation, we cannot create or manipulate archetypes at runtime. This is not an issue for small to medium sized projects, but for anything larger or with complex data interactions, you might have to resort to other ECS implementations or libraries/packages.

> If _name clashing_ is a concern, then just name the component fields instead of embedding them.

In order to keep track of the instances, we use slices.

```go
gameObjects []gameObject
```

Since entities and components are purely value types, there is no need to maintain an object pool.

### Systems

Systems are implemented as functions.

```go
func movableSystem(delta float32) {
	count := len(world.gameObjects)

	for i := 0; i < count; i++ {
		e := &world.gameObjects[i]

		vel := e.velocity
		pos := e.position

		pos.X += vel.VX * delta
		pos.Y += vel.VY * delta

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
```

You can be as simple or as fancy as you want to build up your systems, but they are just essentially functions that will be called within the game loop (or even outside the loop if you only need them to run once).
