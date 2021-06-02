package main

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
	"time"

	fun_stuff "github.com/mreza0100/golog/fun_stuff"
)

func clearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func startFromTopTerminal() {
	fmt.Fprint(os.Stdout, "\033[H")
}

func clearCurrentLine() {
	fmt.Printf("%c[2K", 27)
}

// ---

var (
	rendredTimes int
	startedTime  time.Time
	once         sync.Once

	framePerSecond int
)

func frameTicker() {
	lastRendre := rendredTimes

	for {
		time.Sleep(time.Second)
		framePerSecond = rendredTimes - lastRendre
		lastRendre = rendredTimes
	}
}

func getRuntime() string {
	since := time.Since(startedTime)
	return fmt.Sprintf("%vH : %vM : %vS", int(since.Hours()), int(since.Minutes())%60, int(since.Seconds())%60)
}

func init() { clearTerminal(); go frameTicker(); startedTime = time.Now() }

func draw(world WorldT) {
	startFromTopTerminal()
	rendredTimes++

	{
		fmt.Printf(fun_stuff.ColorGreen)

		fmt.Printf("%c[2K%v : %v\n", 27, "rendredTimes    ", rendredTimes)
		fmt.Printf("%c[2K%v : %v\n", 27, "population      ", world.countLives())
		fmt.Printf("%c[2K%v : %v\n", 27, "runtime         ", getRuntime())
		fmt.Printf("%c[2K%v : %vfps\n", 27, "frame per second", framePerSecond)
	}

	for _, i := range world {
		for _, j := range i {
			if j {
				fmt.Print(liveSellColor, liveSellSymbol)
			} else {
				fmt.Print(deadSellColor, deadSellSymbol)
			}
		}
		fmt.Print("\n")
	}

	fmt.Printf("%s[%dm", "\x1b", 0)
}
