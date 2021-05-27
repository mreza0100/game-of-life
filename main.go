package main

import (
	"sync"
	"time"
)

// configs
const (
	height = 30
	width  = 50
	delay  = time.Second * 1
	// must be (1 > initialPopulation )
	// closer to 1 = more initialPopulation
	initialPopulation = 2

	liveSymbol = "$"
	deadSymbol = " "
)

func rebuildWorld(prevWorld WorldT) WorldT {
	newWorld := WorldT{}

	var wg = sync.WaitGroup{}
	for x := range prevWorld {
		wg.Add(1)

		go func(x int) {

			for y, isALiveSell := range prevWorld[x] {
				if isALiveSell {
					newWorld[x][y] = SellT(isALiveSell.canStayAlive(x, y, prevWorld))
				} else {
					newWorld[x][y] = SellT(isALiveSell.canBeAlive(x, y, prevWorld))
				}
			}

			wg.Done()
		}(x)
	}
	wg.Wait()

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
