package main

import (
	"fmt"
	"time"
)

func main() {
	// synchronous
	// printName("Jennie")
	// printName("Rose")

	// asynchronous
	go printName("Jennie")
	go printName("Rose")
	time.Sleep(time.Second * 5)
}

func printName(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}
}
