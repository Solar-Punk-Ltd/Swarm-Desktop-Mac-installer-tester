package main

import (
	"fmt"
	"os"
)

var steps, interval int = 1, 10

func main() {

	if len(os.Args) >= 3 {
		_, steps_err := fmt.Sscan(os.Args[1], &steps)

		if steps_err != nil {
			fmt.Println("ERROR iteration number:", steps_err.Error())
		}

		_, interval_err := fmt.Sscan(os.Args[2], &interval)

		if interval_err != nil {
			fmt.Println("ERROR time interval:", interval_err.Error())
		}

		input, werr := warning()

		if werr != nil {
			fmt.Println("ERROR:", werr.Error())
		} else {
			if input == "OK" {
				fmt.Println("Let's run")
				myInterval(steps, interval)
			} else {
				fmt.Println("Try again")
			}
		}

	} else {
		input, werr := warning()

		if werr != nil {
			fmt.Println("ERROR:", werr.Error())
		} else {
			if input == "OK" {
				fmt.Println("Missing iteration and/or interval argument")
				fmt.Println("Running test only once")
				myInterval(steps, interval)
			} else {
				fmt.Println("Try again")
			}
		}
	}

}
