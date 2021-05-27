package main

type SellT struct {
	alive bool
	x     int
	y     int
}

func init() {
	liveLocations = append(liveLocations,
		[]int{1, 0},
		[]int{1, 1},
		[]int{1, -1},
		[]int{0, 1},
		[]int{0, -1},
		[]int{-1, 0},
		[]int{-1, 1},
		[]int{-1, -1},
	)
}

var liveLocations = make([][]int, 0, 8)

func (s SellT) getAlives(world WorldT) int {
	lives := 0

	for _, i := range liveLocations {
		lives = lives + world.getAndPluse(i[0]+s.x, i[1]+s.y)
	}

	return lives
}

func (s SellT) canStayAlive(world WorldT) bool {
	lives := s.getAlives(world)

	if lives == 2 || lives == 3 {
		return true
	}

	return false
}

func (s SellT) canBeAlive(world WorldT) bool {
	return s.getAlives(world) == 3
}
