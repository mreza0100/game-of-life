package main

import (
	"runtime"
	"time"

	fun_stuff "github.com/mreza0100/golog/fun_stuff"
)

func init() {
	runtime.GOMAXPROCS(10000)
}

// configs
const (
	// world sizes
	height = 70
	width  = 380

	// deley between frames
	delay = time.Second / 30

	// must be 1 > initialPopulation
	// closer to 1 = more initialPopulation
	initialPopulation = 2

	// size of "detectRepeatedPatterns" for detecting
	// repeated patterns more number = slower render
	memoryPatternSize = 30
	// prevent live block in world
	detectRepeatedPatterns = true

	// live/dead symbol and colors
	liveSellSymbol = "@"
	liveSellColor  = fun_stuff.ColorRed
	deadSellSymbol = " "
	deadSellColor  = ""
)
