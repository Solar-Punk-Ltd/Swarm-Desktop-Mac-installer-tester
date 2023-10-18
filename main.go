package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

		fmt.Println("THIS IS SWARM INSTALLER TESTER")
		fmt.Println("Before run this app, DELETE SWARM DESKTOP app from /Application, and IF YOU HAVE ANY FUNDS in your wallet backup it - https://docs.ethswarm.org/docs/desktop/backup-restore")
		fmt.Println("Type OK, if you are ready!")

		reader := bufio.NewReader(os.Stdin)
		str, str_err := reader.ReadString('\n')
		input := strings.Trim(str, "\n")
		ok := "OK"

		if str_err != nil {
			fmt.Println("ERROR:", str_err.Error())
		} else {
			if input == ok {
				fmt.Println("Let's run")
				myInterval(steps, interval)
			} else {
				fmt.Println("Try again")
			}
		}
	} else {
		fmt.Println("Missing iteration and/or interval argument")
		fmt.Println("Running test only once")
		myInterval(steps, interval)
	}
	/*
	   beeLog, beeLogErr := os.ReadFile("/Users/zolmac/Library/Logs/Swarm Desktop/bee.current.log")

	   	if beeLogErr != nil {
	   		//log.Fatal(beeLogErr)
	   		fmt.Println(beeLogErr)
	   	} else {

	   		checkBeeLog(string(beeLog))
	   	}
	*/
}
