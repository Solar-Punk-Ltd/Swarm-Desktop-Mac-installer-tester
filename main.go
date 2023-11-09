package main

import (
	"fmt"
	"sync"
	"time"
)

var steps, interval int = 1, 10

type result struct {
	success int
	failure int
}

func main() {
	var currentTask string //add to myInterval as arg
	resultCounter := result{0, 0}
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
			for round < steps {
				if round > 0 {
					time.Sleep(time.Duration(interval) * time.Second)
				}
				testRound(steps, interval, &currentTask, round, &resultCounter)
				if round == steps-1 {
					return
				}
				round++
			}
		} else if input == "q" {
			fmt.Println("Quit main")
		} else {
			fmt.Println("Try again")
		}
	}

	wg.Wait()
}
