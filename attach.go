package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func mount_and_install(resultMap *map[string]bool) {
	attachDmg := exec.Command("hdiutil", "attach", "Swarm.Desktop-0.30.0-arm64.dmg")
	_, attachErr := shellCommandHandler(attachDmg, "MOUNT SWARM DESKTOP.DMG")

	if attachErr != nil {
		fmt.Println("Attach error:", attachErr)
		log.Println("Attach error:", string(attachErr.Error()))
		(*resultMap)["attach"] = false
	} else {
		//fmt.Println("Attach success:", attachOut)
		//log.Println("Attach success:", attachOut)
		(*resultMap)["attach"] = true
		fmt.Println("--------------------------")
		defer detach()
		time.Sleep(3 * time.Second)

		install := exec.Command("cp", "-R", "/Volumes/Swarm Desktop/Swarm Desktop.app", "/Applications")
		installOut, installErr := shellCommandHandler(install, "INSTALL SWARM DESKTOP.APP")

		if installErr != nil {
			//fmt.Println("Install error:", installErr)
			//log.Println("Install error:", string(installErr.Error()))
			(*resultMap)["install"] = false
		} else {
			//fmt.Println("Installation success:", installOut)
			log.Println("Installation success", installOut)
			(*resultMap)["install"] = true
			fmt.Println("--------------------------")
			time.Sleep(3 * time.Second)
			//RUN
			runProg := exec.Command("open", "/Applications/Swarm Desktop.app")
			_, runProgErr := shellCommandHandler(runProg, "OPEN SWARM DESKTOP.APP")

			if runProgErr != nil {
				(*resultMap)["run"] = false
			} else {
				(*resultMap)["run"] = true
				fmt.Println("--------------------------")
				time.Sleep(30 * time.Second)

				//CD data-dir
				checkDataDir := exec.Command("cd", "/Users/zolmac/Library/Application Support/Swarm Desktop")
				_, checkDataDirErr := shellCommandHandler(checkDataDir, "CHECK DATA-DIR EXISTENCE")
				if checkDataDirErr != nil {
					//log.Println("Data-dir check error:", string(checkDataDirErr.Error()))
					(*resultMap)["check_data_dir"] = false
				} else {
					//log.Println("Data-dir check success:", checkDataDirOut)
					(*resultMap)["check_data_dir"] = true
				}
				fmt.Println("--------------------------")
				time.Sleep(3 * time.Second)

				//CHECK bee.log /Users/zolmac/Library/Logs/Swarm Desktop
				beeLog, beeLogErr := os.ReadFile("/Users/zolmac/Library/Logs/Swarm Desktop/bee.current.log")
				if beeLogErr != nil {
					log.Println(beeLogErr)
					(*resultMap)["check_data_dir"] = false
				} else {
					checkBeeLog(string(beeLog), resultMap)
				}

				fmt.Println("--------------------------")
				time.Sleep(3 * time.Second)
				exitSwarm := exec.Command("pkill", "swarm-desktop")
				shellCommandHandler(exitSwarm, "EXIT SWARM DESKTOP.APP")
				fmt.Println("--------------------------")
				exitSafari := exec.Command("pkill", "Safari")
				shellCommandHandler(exitSafari, "EXIT SAFARI")
				fmt.Println("--------------------------")
			}

			time.Sleep(3 * time.Second)
			delete("/Applications/Swarm Desktop.app")
			fmt.Println("--------------------------")
			delete("/Users/zolmac/Library/Application Support/Swarm Desktop")
			fmt.Println("--------------------------")
			delete("/Users/zolmac/Library/Logs/Swarm Desktop/")
		}
		fmt.Println("--------------------------")
	}
}
