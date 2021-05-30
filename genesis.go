package main

import (
	"math/rand"
	"time"
)

func init() { rand.Seed(time.Now().Unix()) }
func genesis() WorldT {
	world := WorldT{}

	for x, i := range world {
		for y := range i {
			world[x][y] = SellT(rand.Intn(initialPopulation) == 1)
		}
	}

	return world
}
