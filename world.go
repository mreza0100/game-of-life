package main

type WorldT [width][height]SellT

func isInRange(position int, Range int) bool {
	return position < Range && position >= 0
}

func (world WorldT) getAndPluse(x, y int) int {
	if !isInRange(x, width) || !isInRange(y, height) {
		return 0
	}

	if world[x][y].alive {
		return 1
	}

	return 0
}
