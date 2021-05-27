package main

import (
	"math/rand"
	"time"
)

func init() { rand.Seed(time.Now().Unix()) }
func genesis() WorldT {
	world := WorldT{}

	for idx, i := range world {
		for idx2 := range i {
			world[idx][idx2] = rand.Intn(initialPopulation) == 1
		}
	}

	draw(world)
	return world
}
