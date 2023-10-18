package main

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

func exitprog(progName string) {

	listProg := exec.Command("ps", "-ax")
	listOut, listErr := shellCommandHandler(listProg, progName) //listProg.Output()

	if listErr != nil {
		fmt.Println("Exit ERROR:", listErr)
	} else {
		output := listOut
		//fmt.Println(output)

		var lines []string
		var pids []string

		for len(output) > 0 {
			//fmt.Println("--------for-----------")
			before, after, _ := strings.Cut(output, "\n")
			//fmt.Println("Before:", before)
			//fmt.Println("After:", after)
			lines = append(lines, before)
			output = after
		}

		for i := 0; i < len(lines); i++ {
			findRes := strings.Index(lines[i], progName)
			if findRes > -1 {
				//				_, pid_details, _ := strings.Cut(lines[i], " ")
				//pid_num, _, _ := strings.Cut(lines[i], " ")
				regex := regexp.MustCompile(`\d+`)
				pidNum := regex.FindString(lines[i])
				pids = append(pids, pidNum)
			}
		}

		//fmt.Println(len(lines))
		//fmt.Println(lines[0])
		//fmt.Println(lines[1])
		//fmt.Println(lines[len(lines)-1])
		fmt.Println("PID numbers:")
		fmt.Println(pids)
		//recursion
		if len(pids) > 0 {
			for j := 0; j < len(pids); j++ {
				kill_prog := exec.Command("kill", pids[j])
				kill_out, kill_err := kill_prog.Output()

				if kill_err != nil {
					fmt.Println("Kill ERROR:")
					fmt.Println(kill_err.Error())
					log.Println("Kill ERROR at " + progName + ": " + kill_err.Error())
				} else {
					fmt.Println("Kill SUCCESS:", pids[j])
					fmt.Println(string(kill_out))
					log.Println("Kill SUCCESS "+progName+": ", pids[j]+string(kill_out))
				}
			}
		} else {
			fmt.Println("No running program found")
		}

	}
}
