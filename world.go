package main

type WorldT [height][width]SellT

func isInRange(position int, Range int) bool {
	return position >= 0 && position < Range
}

func (w WorldT) isALiveSell(x, y int) int {
	if !isInRange(x, height) || !isInRange(y, width) {
		return 0
	}

	if w[x][y] {
		return 1
	}

	return 0
}
func (w WorldT) countLives() int {
	lives := 0

	for _, i := range w {
		for _, j := range i {

			if j {
				lives++
			}

		}
	}

	return lives
}
