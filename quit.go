package main

import (
	"fmt"
	"sync"
)

func quit(round *int, steps int, wg *sync.WaitGroup, currentTask *string) {

	in, err := myReader()
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		if in == "qi" {
			fmt.Println("Quit immediate")
			if *currentTask != "finish" {
				exitProg(*currentTask, "pkill", *currentTask)
			}
			*round = steps
			wg.Done()
		}
		if in == "q" {
			fmt.Println("Quit at the of the round")
			*round = steps
			wg.Done()
		}
	}
}
