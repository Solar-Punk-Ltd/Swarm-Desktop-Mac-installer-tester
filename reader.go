package main

import (
	"bufio"
	"os"
	"strings"
)

func myReader() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	input := strings.Trim(str, "\n")

	return input, err
}
