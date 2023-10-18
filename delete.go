package main

import (
	//"fmt"
	"os/exec"
)

func delete(arg string) string {

	deleteCmd := exec.Command("rm", "-r", arg)
	deleteOut, deleteErr := shellCommandHandler(deleteCmd, "Delete: "+arg)

	if deleteErr != nil {
		//fmt.Println("Delete ERROR at", arg, ":", deleteErr.Error())
		return "Delete ERROR at " + arg + ": " + deleteErr.Error()
	} else {
		//fmt.Println("Delete SUCCESS at", arg, ":", deleteOut)
		return "Delete SUCCESS at " + arg + ": " + deleteOut
	}
}
