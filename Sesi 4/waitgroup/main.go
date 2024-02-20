package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go printName("Jennie", &wg)
	go printName("Rose", &wg)
	wg.Wait()
}

func printName(s string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}

	defer wg.Done()
}
