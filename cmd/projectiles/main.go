package main

import (
	"fmt"
	"ray-tracer/feature"
)

type projectile struct {
	position feature.Tuple // it has to be a point
	velocity feature.Tuple // it has to be a vector
}

type environment struct {
	gravity feature.Tuple // it has to be a vector
	wind    feature.Tuple // it has to be a vector
}

// tick return a new projectile after one unit of time (tick).
func tick(env environment, proj projectile) (projectile, error) {
	var p projectile

	pos, err := proj.position.Add(proj.velocity)
	if err != nil {
		return p, err
	}

	vel, err := proj.velocity.Add(env.gravity)
	if err != nil {
		return p, err
	}

	vel, err = vel.Add(env.wind)
	if err != nil {
		return p, err
	}

	p.position = pos
	p.velocity = vel

	return p, nil
}

func main() {
	// projectile starts one unit above the origin.​
	// velocity is normalized to 1 unit/tick.​
	vel, err := feature.NewVector(1, 1, 0).Normalize()
	if err != nil {
		panic(fmt.Errorf("failed to initialize the velocity: %v", err))
	}

	proj := projectile{
		position: feature.NewPoint(0, 1, 0),
		velocity: vel,
	}

	// gravity -0.1 unit/tick, and wind is -0.01 unit/tick.​
	env := environment{
		gravity: feature.NewVector(0, -0.1, 0),
		wind:    feature.NewVector(-0.01, 0, 0),
	}

	for {
		if proj.position.Y <= 0 {
			break
		}

		proj, err = tick(env, proj)
		if err != nil {
			panic(fmt.Errorf("failed to tick: %v", err))
		}

		fmt.Printf("proj new position: %v\n", proj.position)
	}

	fmt.Println("\nmission accomplished!")
}
