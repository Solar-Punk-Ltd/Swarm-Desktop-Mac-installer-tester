package main

import (
	"fmt"
	"log"
	"os/exec"
)

func runprog(name string) {
	run := exec.Command("open", name)
	run_out, run_err := run.Output()

	if run_err != nil {
		fmt.Println("Run ERROR:", run_err.Error())
		log.Println("Run ERROR:", run_err.Error())
	} else {
		fmt.Println("Run SUCCESS:", string(run_out))
		log.Println("Run SUCCESS:", string(run_out))
	}
}
