package main

import (
	"fmt"
	"os/exec"
	"time"
)

func download(resultMap *map[string]bool, currentTask *string) {
	*currentTask = "wget"
	defer delete(APPNAME)
	downloadDmg := exec.Command("wget", "https://github.com/ethersphere/swarm-desktop/releases/download/"+LIB+"/"+APPNAME)
	_, downloadErr := shellCommandHandler(downloadDmg, "DOWNLOAD SWARM DESKTOP")
	if downloadErr != nil {
		(*resultMap)["download"] = false
	} else {
		(*resultMap)["download"] = true
		fmt.Println("--------------------------")
		time.Sleep(2 * time.Second)

		mount_and_install(resultMap, currentTask)
		fmt.Println("--------------------------")
		time.Sleep(2 * time.Second)

	}
}
