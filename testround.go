package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func testRound(steps int, interval int, currentTask *string, round int, resultCounter *result) {

	fmt.Printf("\n\n\n")
	fmt.Println("--------------------------")
	fmt.Println("TEST ROUND:", round+1)
	fmt.Println("--------------------------")

	var resultMap map[string]bool = make(map[string]bool)
	resultMap["log_file"] = false
	resultMap["download"] = false
	resultMap["attach"] = false
	resultMap["install"] = false
	resultMap["run"] = false
	resultMap["check_data_dir"] = false
	resultMap["check_bee_log"] = false
	resultMap["check_desktop_log"] = false

	file, file_err := os.OpenFile("installer_log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if file_err != nil {
		log.Fatal(file_err)
	} else {
		log.SetOutput(file)
		fmt.Println("log file opened")
		resultMap["log_file"] = true
		fmt.Println("--------------------------")
		log.Printf("\n\n\n\n\nTEST ROUND No.%d at %s\n\n", round+1, time.Now())
		download(&resultMap, currentTask)
	}

	//fmt.Println("result map length:", len(resultMap))
	iteration := 0
	successCount := 0

	fmt.Print(
		"Test result at log-file          :", resultMap["log_file"], "\n",
		"Test result at download          :", resultMap["download"], "\n",
		"Test result at attach            :", resultMap["attach"], "\n",
		"Test result at install           :", resultMap["install"], "\n",
		"Test result at run               :", resultMap["run"], "\n",
		"Test result at check-data-dir    :", resultMap["check_data_dir"], "\n",
		"Test result at check-bee-log     :", resultMap["check_bee_log"], "\n",
		"Test result at check-desktop-log :", resultMap["check_desktop_log"], "\n",
	)
	for item := range resultMap {
		//fmt.Println("Test result at:", item, ":", resultMap[item])
		iteration++

		if resultMap[item] {
			successCount++
		}

		if iteration == len(resultMap) {
			switch successCount {
			case 8:
				(*resultCounter).success++
				fmt.Println("TEST ROUND SUCCEED")
			default:
				(*resultCounter).failure++
				fmt.Println("TEST ROUND FAILED")
			}
		}
	}

	fmt.Println("\nTest round count: ", (*resultCounter).success+(*resultCounter).failure, "/", steps)
	fmt.Println("SUCCESS:", (*resultCounter).success)
	fmt.Println("FAILED:", (*resultCounter).failure)

	*currentTask = "finish"
}
