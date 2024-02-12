package main

import "fmt"

func main() {
	score := 80

	if score > 80 {
		switch score {
		case 100:
			fmt.Println("perfect score")
		default:
			fmt.Println("nice score")
		}
	} else {
		if score == 50 {
			fmt.Println("not bad")
		} else if score == 30 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if score == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}
