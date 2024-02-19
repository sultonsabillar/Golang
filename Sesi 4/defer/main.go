package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Haiiiiii")
	fmt.Println("Halloooo")

	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}

	// exit example
	defer fmt.Println("Invoke with defer")
	fmt.Println("Test before exiting")
	os.Exit(0)
}
