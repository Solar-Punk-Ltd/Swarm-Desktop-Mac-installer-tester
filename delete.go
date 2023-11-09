package main

import (
	"fmt"
	"os/exec"
)

func delete(arg string) {

	deleteCmd := exec.Command("rm", "-r", arg)
	_, deleteErr := shellCommandHandler(deleteCmd, "Delete: "+arg)

	if deleteErr != nil {
		//fmt.Println("Delete ERROR at", arg, ":", deleteErr.Error())
		fmt.Println("--------------------------")
	} else {
		//fmt.Println("Delete SUCCESS at", arg, ":", deleteOut)
		fmt.Println("--------------------------")
	}
}
