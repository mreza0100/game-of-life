package main

import (
	fun_stuff "github.com/mreza0100/golog/fun_stuff"
)

// configs
const (
	// world sizes
	height = 30
	width  = 140

	// deley between frames
	delay = 0

	// must be 1 > initialPopulation
	// closer to 1 = more initialPopulation
	initialPopulation = 20

	// size of "detectRepeatedPatterns" for detecting
	// repeated patterns more number = slower render
	memoryPatternSize = 10
	// prevent live block in world
	detectRepeatedPatterns = true

	// live/dead symbol and colors
	liveSellSymbol = "+"
	liveSellColor  = fun_stuff.ColorRed
	deadSellSymbol = " "
	deadSellColor  = ""
)
