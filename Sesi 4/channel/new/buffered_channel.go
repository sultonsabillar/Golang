package main

import "fmt"

func bufferedChannel() {
	fmt.Println("---- Buffered Channel ----")
	data := make(chan string, 2)
	data <- "Hello"
	data <- "World"

	firstPrint, secondPrint := <-data, <-data

	fmt.Println(firstPrint, secondPrint)
}
