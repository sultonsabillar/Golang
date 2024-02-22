package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Interface1 struct {
	Data string
}

type Interface2 struct {
	Data int
}

var mutex sync.Mutex

func processData1(data Interface1, wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	defer mutex.Unlock()
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	fmt.Println("Processed Interface1 data:", data.Data)
}

func processData2(data Interface2, wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	defer mutex.Unlock()
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	fmt.Println("Processed Interface2 data:", data.Data)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup
	wg.Add(8)

	for i := 0; i < 4; i++ {
		data1 := Interface1{"Data1"}
		data2 := Interface2{42}

		go processData1(data1, &wg)
		go processData2(data2, &wg)
	}

	wg.Wait() // Waiting for all goroutines to finish
}
