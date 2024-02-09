package main

import (
	"fmt"
)

var n int

func main() {

	fmt.Println("Enter Your Number: ")
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		if i == 1 {
			fmt.Println("squarecube")
			continue
		}
		if i == n^2 {
			fmt.Println("square")
			continue
		}
		fmt.Println(i)
	}

	// var n int
	// fmt.Print("Masukkan nilai n: ")
	// fmt.Scanln(&n)

	// for i := 1; i <= n; i++ {
	// 	isSquare := isPerfectSquare(i)
	// 	isCube := isPerfectCube(i)

	// 	switch {
	// 	case isSquare && isCube:
	// 		fmt.Println("SquareCube")
	// 	case isSquare:
	// 		fmt.Println("Square")
	// 	case isCube:
	// 		fmt.Println("Cube")
	// 	default:
	// 		fmt.Println(i)
	// 	}
	// }
}

// // Fungsi untuk mengecek apakah suatu angka adalah kuadrat sempurna
// func isPerfectSquare(num int) bool {
// 	root := int(math.Sqrt(float64(num)))
// 	return root*root == num
// }

// // Fungsi untuk mengecek apakah suatu angka adalah kubus sempurna
// func isPerfectCube(num int) bool {
// 	root := int(math.Round(math.Pow(float64(num), 1.0/3.0)))
// 	return root*root*root == num
// }
