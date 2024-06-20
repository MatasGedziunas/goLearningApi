package main

import (
	"fmt"
)

func helloSir() {
	fmt.Println("hello sir")
}

func arrays() {
	var intArr [3]int
	fmt.Println(&intArr[0])
}

func slices() {
	var intSlice []int
	// intSlice = make([]int, 0)
	intSlice = append(intSlice, 5)
	fmt.Print(intSlice, len(intSlice))
}

func maps() {
	var maps map[string]int = make(map[string]int)
	var hello, exists = maps["hello"]
	println(hello, exists)
	maps["hello"] = 8
	hello, exists = maps["hello"]
	println(hello, exists)
}

func loops() {
	var hello map[string]int = make(map[string]int)
	hello["hello"] = 8
	hello["goodbye"] = 5
	for key, val := range hello {
		println(key, val)
	}
}
