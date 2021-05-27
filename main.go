package main

import (
	"time"
)

// configs
const (
	height     = 130
	width      = 680
	delay      = 0
	population = 2

	liveSymbol = "$"
	deadSymbol = " "
)

func rebuildWorld(prevWorld WorldT) WorldT {
	newWorld := WorldT{}

	for x, i := range prevWorld {
		for y, isALiveSell := range i {
			if isALiveSell {
				newWorld[x][y] = SellT(isALiveSell.canStayAlive(x, y, prevWorld))
			} else {
				newWorld[x][y] = SellT(isALiveSell.canBeAlive(x, y, prevWorld))
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
