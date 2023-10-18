package main

import (
	"fmt"
	"log"
	"os/exec"
)

func detach() {
	detachDmg := exec.Command("hdiutil", "detach", "/dev/disk4s1")
	detachOut, detachErr := shellCommandHandler(detachDmg, "UNMOUNT SWARM DESKTOP.DMG")
	if detachErr != nil {
		fmt.Println("Detach error:", detachErr)
		log.Println("Detach error", string(detachErr.Error()))
	} else {
		//fmt.Println("Detach success:", detachOut)
		log.Println("Detach success:", detachOut)
	}

}
