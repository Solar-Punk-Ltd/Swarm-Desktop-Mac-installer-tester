package main

import (
	"fmt"
	"os/exec"
)

func exitProg(name string, command string, args ...string) {
	exit := exec.Command(command, args...)
	shellCommandHandler(exit, "EXIT "+name)
	fmt.Println("--------------------------")
}
