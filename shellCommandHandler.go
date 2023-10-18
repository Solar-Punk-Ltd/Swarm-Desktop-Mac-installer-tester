package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os/exec"
)

func shellCommandHandler(cmd *exec.Cmd, cmdName string) (string, error) {
	var stdErr error
	var stdOut string
	var stdoutPipe io.ReadCloser
	var stdoutPipeErr error

	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println("StderrPipe ERROR at", cmdName, ":", err)
	}

	if cmdName != "DOWNLOAD SWARM DESKTOP" {
		stdoutPipe, stdoutPipeErr = cmd.StdoutPipe()
		if stdoutPipeErr != nil {
			fmt.Println("StdoutPipe ERROR at", cmdName, ":", stdoutPipeErr)
		}
	}

	startErr := cmd.Start()
	if startErr != nil {
		fmt.Println("Start ERROR at", cmdName, ":", startErr)
	} else {
		fmt.Println("Start", cmdName, "SUCCESS")
	}

	if cmdName != "DOWNLOAD SWARM DESKTOP" {
		stdoutPipeContent, stdoutPipeContentErr := io.ReadAll(stdoutPipe)
		if stdoutPipeContentErr != nil {
			fmt.Println("Read stdout pipe ERROR at", cmdName, ":", stdoutPipeContentErr)
		} else {
			stdlen := len(stdoutPipeContent)
			fmt.Println("stdout length:", stdlen)
			stdOut = string(stdoutPipeContent)
			var stdOutLog string
			if stdlen < 2500 {
				stdOutLog = string(stdoutPipeContent)
			} else {
				for i := 0; i < 2000; i++ {
					stdOutLog += string(stdoutPipeContent[i])
				}
				for j := stdlen - 500; j < stdlen; j++ {
					stdOutLog += string(stdoutPipeContent[j])
				}
			}
			log.Printf("stdout %s: %s\n", cmdName, stdOutLog)
		}
	}

	stderrPipeContent, stderrPipeContentErr := io.ReadAll(stderrPipe)
	if stderrPipeContentErr != nil {
		fmt.Println("Read stderr pipe ERROR at", cmdName, ":", stderrPipeContentErr)
	} else {
		stdlen := len(stderrPipeContent)
		fmt.Println("stderr length:", stdlen)
		var stderr string
		if stdlen < 2500 {
			stderr = string(stderrPipeContent)
		} else {
			for i := 0; i < 2000; i++ {
				//s := []string{stdout, string(stderrPipeContentlurp[i])}
				//stdout = strings.Join(s, "")
				stderr += string(stderrPipeContent[i])
			}
			for j := stdlen - 500; j < stdlen; j++ {
				stderr += string(stderrPipeContent[j])
			}
		}
		log.Printf("stderr %s: %s\n", cmdName, stderr)
		stdErr = errors.New(cmdName + "stderr :" + stderr)
	}

	waitErr := cmd.Wait()
	if waitErr != nil {
		fmt.Println("Wait ERROR at", cmdName, ":", waitErr, "stderr:", stdErr)
		return "ERROR", stdErr
	} else {
		fmt.Println(cmdName, "finished wit SUCCESS")
		if cmdName == "DOWNLOAD SWARM DESKTOP" {
			stdOut = stdErr.Error()
		}
	}

	return stdOut, nil
}
