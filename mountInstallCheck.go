package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func mount_and_install(resultMap *map[string]bool, currentTask *string) {
	*currentTask = "hdiutil"
	attachDmg := exec.Command("hdiutil", "attach", APPNAME)
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

		*currentTask = "cp"
		defer delete("/Applications/Swarm Desktop.app")
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
			*currentTask = "open"
			runProg := exec.Command("open", "/Applications/Swarm Desktop.app")
			_, runProgErr := shellCommandHandler(runProg, "OPEN SWARM DESKTOP.APP")

			if runProgErr != nil {
				(*resultMap)["run"] = false
			} else {
				(*resultMap)["run"] = true
				defer delete(HOMEDIR + "/Library/Application Support/Swarm Desktop")
				defer delete(HOMEDIR + "/Library/Logs/Swarm Desktop/")
				defer exitProg(BROWSER, "osascript", "-e", "tell application \""+BROWSER+"\" to quit")
				defer exitProg("SWARM DESKTOP.APP", "pkill", "swarm-desktop")
				fmt.Println("--------------------------")
				time.Sleep(30 * time.Second)

				//CD data-dir
				*currentTask = "cd"
				checkDataDir := exec.Command("cd", HOMEDIR+"/Library/Application Support/Swarm Desktop")
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
				beeLogKeyFrases := []string{
					"\"msg\"=\"using chain with network network\" \"chain_id\"=100 \"network_id\"=1",
					"\"msg\"=\"starting debug server\" \"address\"=\"127.0.0.1:1635\"",
					"\"msg\"=\"using datadir\" \"path\"=\"" + HOMEDIR + "/Library/Application Support/Swarm Desktop/data-dir\"",
					"\"msg\"=\"starting in ultra-light mode\"",
					"\"msg\"=\"connected\" \"tld\"=\"\" \"endpoint\"=\"https://cloudflare-eth.com\"",
					"\"msg\"=\"starting api server\" \"address\"=\"127.0.0.1:1633\"",
				}

				beeLog, beeLogErr := os.ReadFile(HOMEDIR + "/Library/Logs/Swarm Desktop/bee.current.log")
				if beeLogErr != nil {
					log.Println(beeLogErr)
					(*resultMap)["check_bee_log"] = false
				} else {
					checkBeeLog(string(beeLog), resultMap, "bee", beeLogKeyFrases)
				}

				desktopLogKeyFrases := []string{
					"msg=\"OK\"",
					"msg=\"Serving UI from path: /Applications/Swarm Desktop.app/Contents/Resources/app.asar/dist/ui\"",
					"msg=\"Creating new Bee config.yaml\"",
				}

				desktopLog, desktopLogErr := os.ReadFile(HOMEDIR + "/Library/Logs/Swarm Desktop/bee-desktop.log") //??
				if desktopLogErr != nil {
					log.Println(desktopLogErr)
					(*resultMap)["check_desktop_log"] = false
				} else {
					checkBeeLog(string(desktopLog), resultMap, "desktop", desktopLogKeyFrases)
				}
			}

		}
		fmt.Println("--------------------------")
	}
}
