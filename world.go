package main

type WorldT [height][width]SellT

func isInRange(position int, Range int) bool {
	return position < Range && position >= 0
}

func (world WorldT) getAndPluse(x, y int) int {
	if !isInRange(x, height) || !isInRange(y, width) {
		return 0
	}

	if world[x][y].alive {
		return 1
	}

	return 0
}
