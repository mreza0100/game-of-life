package main

import (
	"fmt"
)

func clearTerminal() {
	fmt.Print("\033[H\033[2J")
}
func startFromTopTerminal() {
	fmt.Print("\033[H")
}

var rendredTimes int

func getCurrentPopulation(world WorldT) int {
	sells := 0

	for _, i := range world {
		for _, j := range i {
			if j.alive {
				sells++
			}
		}
	}

	return sells
}

func draw(world WorldT) {
	startFromTopTerminal()

	rendredTimes++
	fmt.Println("rendredTimes: ", rendredTimes)
	fmt.Println("population: ", getCurrentPopulation(world))
	fmt.Println()

	for _, i := range world {
		for _, j := range i {
			symbol := deadSymbol
			if j.alive {
				symbol = liveSymbol
			}

			fmt.Print(symbol)
		}

		fmt.Print("                             |")
		fmt.Println()
	}
}
