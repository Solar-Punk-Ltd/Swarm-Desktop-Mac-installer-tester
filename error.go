package main

import (
	"os/exec"
)

func myError() {
	cmd := exec.Command("cp", "pending", "../")
	shellCommandHandler(cmd, "cp")
}
