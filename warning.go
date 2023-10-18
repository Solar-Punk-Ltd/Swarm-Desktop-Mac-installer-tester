package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func warning() (string, error) {
	fmt.Println("THIS IS SWARM INSTALLER TESTER")
	fmt.Println("Before run this app, DELETE SWARM DESKTOP app from /Applications, and IF YOU HAVE ANY FUNDS in your wallet backup them - https://docs.ethswarm.org/docs/desktop/backup-restore")
	fmt.Println("Type OK, if you are ready!")

	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	input := strings.Trim(str, "\n")

	return input, err
}
