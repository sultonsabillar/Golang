package main

import (
	"fmt"
	"time"
)

func other() {
	fmt.Println("My first go routine")
	go func() {
		fmt.Println("My other go routine")
	}()
	time.Sleep(time.Second)
}
