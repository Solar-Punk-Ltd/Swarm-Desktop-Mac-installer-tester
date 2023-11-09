package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func initTester(steps *int, interval *int, LIB *string, APPNAME *string, HOMEDIR *string, BROWSER *string) {
	getCurrentDir := exec.Command("pwd")
	currentDir, getDirErr := getCurrentDir.Output()
	if getDirErr != nil {
		fmt.Println("Checking current dir ERROR:", getDirErr.Error())
	} else {
		_, after, _ := strings.Cut(string(currentDir), "/Users/")
		beforeSlash, _, _ := strings.Cut(after, "/")
		*HOMEDIR = "/Users/" + beforeSlash
	}

	if len(os.Args) >= 3 {
		_, stepsErr := fmt.Sscan(os.Args[1], steps)

		if stepsErr != nil {
			fmt.Println("ERROR iteration number:", stepsErr.Error())
		}

		_, intervalErr := fmt.Sscan(os.Args[2], interval)

		if intervalErr != nil {
			fmt.Println("ERROR time interval:", intervalErr.Error())
		}

		var version string
		_, versionErr := fmt.Sscan(os.Args[3], version)
		if versionErr != nil {
			fmt.Println("ERROR version name:", versionErr.Error())
		} else {
			*LIB = "v" + version
			*APPNAME = "Swarm.Desktop-" + version + "-arm64.dmg"
		}

		_, browserErr := fmt.Sscan(os.Args[4], *BROWSER)
		if browserErr != nil {
			fmt.Println("ERROR browser name", browserErr.Error())
		}

	} else {
		fmt.Println("Missing iteration and/or interval argument")
		fmt.Println("Running test only once")
		fmt.Println("--------------------------")
	}
}
