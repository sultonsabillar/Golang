package main

import "fmt"

func main() {

	// function closure yang bersifat anonymous
	// func(n string) {
	// 	fmt.Println(n)
	// }("Fitri")

	c := make(chan string)
	go func(n string) {
		c <- n
	}("Fitri")

	var chanFunc = func(n string) {
		c <- n
	}

	go chanFunc("hallo")

	b := <-c
	d := <-c
	fmt.Println(b)
	fmt.Println(d)

	chanSecond := make(chan string)
	go printName(chanSecond, "Jisoo")

	result := <-chanSecond
	fmt.Println(result)
}

func printName(c chan string, value string) {
	c <- value
}
