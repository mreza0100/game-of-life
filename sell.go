package main

type SellT bool

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

func (s SellT) getAlives(x, y int, world WorldT) int {
	lives := 0

	for _, i := range liveLocations {
		lives = lives + world.isALiveSell(i[0]+x, i[1]+y)
	}

	return lives
}

func (s SellT) canStayAlive(x, y int, world WorldT) bool {
	lives := s.getAlives(x, y, world)

	return lives == 2 || lives == 3
}

func (s SellT) canBeAlive(x, y int, world WorldT) bool {
	return s.getAlives(x, y, world) == 3
}
