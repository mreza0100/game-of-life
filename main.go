package main

import (
	"sync"
	"time"
)

func rebuildWorld(prevWorld WorldT) WorldT {
	newWorld := WorldT{}
	var wg sync.WaitGroup

	for x := range prevWorld {
		wg.Add(1)

		go func(x int) {
			for y, isALiveSell := range prevWorld[x] {
				if isALiveSell {
					newWorld[x][y] = SellT(isALiveSell.canStayAlive(x, y, prevWorld))
				} else {
					newWorld[x][y] = SellT(isALiveSell.canBackToLife(x, y, prevWorld))
				}
			}

			wg.Done()
		}(x)
	}
	wg.Wait()

	if detectRepeatedPatterns && isRepeatedPattern(newWorld) {
		newWorld = injectSells(newWorld)
	}

	return newWorld
}

func main() {
	world := genesis(initialPopulation)
	draw(world)

	for {
		world = rebuildWorld(world)
		time.Sleep(delay)
		draw(world)
	}
}
