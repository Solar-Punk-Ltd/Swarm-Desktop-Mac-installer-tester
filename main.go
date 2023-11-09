package main

import (
	"fmt"
	"sync"
)

var steps, interval int = 1, 10

func main() {
	var currentTask string //add to myInterval as arg
	var wg sync.WaitGroup
	round := 0

	initTester(&steps, &interval, &LIB, &APPNAME, &HOMEDIR, &BROWSER)

	fmt.Println("THIS IS SWARM INSTALLER TESTER")
	fmt.Println("Before run this app, DELETE SWARM DESKTOP app from /Applications, and IF YOU HAVE ANY FUNDS in your wallet backup them - https://docs.ethswarm.org/docs/desktop/backup-restore")
	fmt.Println("Type OK, if you are ready! - Type q to quit, qi to quit immediate!")

	input, err := myReader()
	if err != nil {
		fmt.Println("ERROR:", err.Error())
	} else {
		if input == "OK" {
			fmt.Println("Let's run")
			wg.Add(1)
			go quit(&round, steps, &wg, &currentTask)
			for i := 0; i < steps; i++ {
				testRound(steps, interval, &currentTask, i)
				if i == steps-1 {
					return
				}
			}
		} else if input == "q" {
			fmt.Println("Quit main")
		} else {
			fmt.Println("Try again")
		}
	}

	wg.Wait()
}
