package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Interface1 struct {
	Data string
}

type Interface2 struct {
	Data int
}

func processData1(data Interface1) {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	fmt.Println("Processed Interface1 data:", data.Data)
}

func processData2(data Interface2) {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	fmt.Println("Processed Interface2 data:", data.Data)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 4; i++ {
		data1 := Interface1{"Data1"}
		data2 := Interface2{42}

		go processData1(data1)
		go processData2(data2)
	}

	time.Sleep(1 * time.Second) // Waiting for goroutines to finish
}
