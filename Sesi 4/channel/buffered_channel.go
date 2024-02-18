package main

import "fmt"

func bufferedChannel() {
	firstChan := make(chan int, 2)

	go func() {
		for value := range firstChan {
			value = <-firstChan
			fmt.Println("Terima data ========", value)
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Println("Kirim data", i)
		firstChan <- i
	}
}
