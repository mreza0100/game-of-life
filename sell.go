package main

type SellT bool

func init() {
	aroundLocations = append(aroundLocations,
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

var aroundLocations = make([][]int, 0, 8)

func (s SellT) getAliveNeighbors(x, y int, world WorldT) int {
	lives := 0

	for _, i := range aroundLocations {
		lives = lives + world.isALiveSell(i[0]+x, i[1]+y)
	}

	return lives
}

func (s SellT) canStayAlive(x, y int, world WorldT) bool {
	lives := s.getAliveNeighbors(x, y, world)

	return lives == 2 || lives == 3
}

func (s SellT) canBackToLife(x, y int, world WorldT) bool {
	return s.getAliveNeighbors(x, y, world) == 3
}
