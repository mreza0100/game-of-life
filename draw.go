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

func draw(data WorldT) {
	startFromTopTerminal()

	for _, i := range data {
		for _, j := range i {
			p := deadSymbol

			if j.alive {
				p = liveSymbol
			}

			fmt.Print(p)
		}

		fmt.Print("                             |")
		fmt.Println()
	}
}
