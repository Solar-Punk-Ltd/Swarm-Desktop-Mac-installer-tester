package main

import (
	"fmt"
	"os"
)

func readBeeLog() {
	file, fileErr := os.ReadFile("/Users/zolmac/Library/Logs/Swarm Desktop/bee.11.log")

	if fileErr != nil {
		fmt.Println("Read bee.log ERROR")
	} else {
		fmt.Println(len(file))
	}
}
