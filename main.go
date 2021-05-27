package main

import (
	"time"
)

// configs
const (
	height     = 10
	width      = 150
	delay      = 0
	population = 2

	liveSymbol = "$"
	deadSymbol = " "
)

func rebuildWorld(prevWorld WorldT) WorldT {
	newWorld := WorldT{}

	for x, i := range prevWorld {
		for y, j := range i {
			newWorld[x][y] = SellT{
				x: x,
				y: y,
			}

			if j.alive {
				newWorld[x][y].alive = j.canStayAlive(prevWorld)
			} else {
				newWorld[x][y].alive = j.canBeAlive(prevWorld)
			}

		}
	}

	return newWorld
}

func main() {
	world := genesis()

	for {
		time.Sleep(delay)

		world = rebuildWorld(world)
		draw(world)
	}
}
