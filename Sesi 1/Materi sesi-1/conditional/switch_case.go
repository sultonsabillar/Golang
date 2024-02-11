package main

import "fmt"

func main() {
	score := 88
	switch {
	case score == 100:
		fmt.Println("Perfect!")
	case (score < 90) && (score > 60):
		fmt.Println("Not bad")
	default:
		fmt.Println("Please study harder")
	}
}
