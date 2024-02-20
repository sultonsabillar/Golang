package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("My first go routine")

	go secondPrint()
	go firstPrint()

	time.Sleep(time.Millisecond * 5)
}

func firstPrint() {
	fmt.Println("My other go routines!")
}

func secondPrint() {
	fmt.Println("My other go routines 2!")
}
