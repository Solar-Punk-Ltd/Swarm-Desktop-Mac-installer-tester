package main

import (
	"fmt"
	"regexp"
	"strings"
)

// param: str string, return: int
func regex() {
	str := `\d+` //`[0-9]+`
	re := regexp.MustCompile(str)
	resAll := re.FindAllString("We were on Safari 297 times, Safar is 757654 a good adventure", -1)
	resFirst := re.FindString("We were on Safari 297 times, Safar is 757654 a good adventure")
	fmt.Println("all:", resAll, "first:", resFirst)

	str2 := []byte(" 234 zoli")
	re2 := regexp.MustCompile(`\d`)
	res2 := re2.FindIndex(str2)

	before, after, found := strings.Cut(string(str2), string(str2[res2[0]-1]))
	if found {
		fmt.Println(before)
		fmt.Println(after)
	}

}
