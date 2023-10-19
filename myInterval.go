package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func myInterval(steps int, interval int) {
	type result struct {
		success int
		failure int
	}

	resultCounter := result{0, 0}

	//fmt.Println("result counter:", resultCounter.success, resultCounter.failure)

	for i := 0; i < steps; i++ {
		fmt.Printf("\n\n\n")
		fmt.Println("--------------------------")
		fmt.Println("TEST ROUND:", i+1)
		fmt.Println("--------------------------")

		var result_map map[string]bool = make(map[string]bool)

		file, file_err := os.OpenFile("installer_log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if file_err != nil {
			log.Fatal(file_err)
		} else {
			log.SetOutput(file)
			fmt.Println("log file opened")
			result_map["log_file"] = true
			fmt.Println("--------------------------")
			log.Printf("\n\n\n\n\nTEST ROUND No.%d at %s\n\n", i+1, time.Now())

			downloadDmg := exec.Command("wget", "https://github.com/ethersphere/swarm-desktop/releases/download/v0.30.0/Swarm.Desktop-0.30.0-arm64.dmg")
			_, downloadErr := shellCommandHandler(downloadDmg, "DOWNLOAD SWARM DESKTOP")
			if downloadErr != nil {
				result_map["download"] = false
			} else {
				result_map["download"] = true
				fmt.Println("--------------------------")
				time.Sleep(2 * time.Second)

				mount_and_install(&result_map)
				fmt.Println("--------------------------")
				time.Sleep(2 * time.Second)

				deleteDmg := exec.Command("rm", "-r", "Swarm.Desktop-0.30.0-arm64.dmg")
				shellCommandHandler(deleteDmg, "DELETE SWARM DESKTOP DMG")
				fmt.Println("--------------------------")
			}
		}

		//fmt.Println("result map length:", len(result_map))
		iteration := 0
		successCount := 0

		fmt.Print(
			"Test result at log-file      :", result_map["log_file"], "\n",
			"Test result at download      :", result_map["download"], "\n",
			"Test result at attach        :", result_map["attach"], "\n",
			"Test result at install       :", result_map["install"], "\n",
			"Test result at run           :", result_map["run"], "\n",
			"Test result at check-data-dir:", result_map["check_data_dir"], "\n",
			"Test result at check-bee-log :", result_map["check_bee_log"], "\n",
		)
		for item := range result_map {
			//fmt.Println("Test result at:", item, ":", result_map[item])
			iteration++

			if result_map[item] {
				successCount++
			}

			if iteration == len(result_map) {
				switch successCount {
				case 7:
					resultCounter.success++
					fmt.Println("TEST ROUND SUCCEED")
				default:
					resultCounter.failure++
					fmt.Println("TEST ROUND FAILED")
				}
			}
		}

		fmt.Println("\nTest round count: ", resultCounter.success+resultCounter.failure)
		fmt.Println("SUCCESS:", resultCounter.success)
		fmt.Println("FAILED:", resultCounter.failure)
		time.Sleep(time.Duration(interval) * time.Second)
	}

	/*
	   *i++

	   	if *i < steps {
	   		myInterval(i, steps, interval)
	   	}
	*/
}
