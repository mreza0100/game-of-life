package main

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
	"time"
)

func clearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func startFromTopTerminal() {
	fmt.Fprint(os.Stdout, "\033[H")
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
	return fmt.Sprintf("%vH : %vM : %vS   ", int(since.Hours()), int(since.Minutes())%60, int(since.Seconds())%60)
}

func init() { clearTerminal(); go frameTicker(); startedTime = time.Now() }

func draw(world WorldT) {
	startFromTopTerminal()

	rendredTimes++

	{
		fmt.Println("rendredTimes     : ", rendredTimes)
		fmt.Println("population       : ", world.countLives())
		fmt.Println("runtime          : ", getRuntime())
		fmt.Println("frame per second : ", framePerSecond)
		fmt.Print("\n")
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
