package main

import "fmt"

func Pointer_test(i *int) {
	fmt.Println("before:", *i)
	*i++
	fmt.Println("after:", *i)
	if *i < 3 {
		fmt.Println("counters i:", i)
		Pointer_test(i)
	}

}
