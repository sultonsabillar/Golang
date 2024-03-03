package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan nilai n: ")
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		isSquare := isPerfectSquare(i)
		isCube := isPerfectCube(i)

		switch {
		case isSquare && isCube:
			fmt.Println("SquareCube")
		case isSquare:
			fmt.Println("Square")
		case isCube:
			fmt.Println("Cube")
		default:package main

		import (
			"bufio"
			"fmt"
			"os"
		)
		
		func main() {
			// Membaca input dari pengguna
			fmt.Print("Masukkan kalimat: ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			kalimat := scanner.Text()
		
			// Inisialisasi map untuk menyimpan jumlah kemunculan huruf
			kemunculanHuruf := make(map[string]int)
		
			// Looping untuk menghitung kemunculan huruf
			for _, huruf := range kalimat {
				kemunculanHuruf[string(huruf)]++
			}
		
			// Tampilkan huruf satu per satu ke bawah
			for _, huruf := range kalimat {
				fmt.Println(string(huruf))
			}
		
			// Tampilkan hasil perhitungan kemunculan huruf
			fmt.Println("\n", kemunculanHuruf)
		}
		
			fmt.Println(i)
		}
	}
}

// Fungsi untuk mengecek apakah suatu angka adalah kuadrat sempurna
func isPerfectSquare(num int) bool {
	root := int(math.Sqrt(float64(num)))
	return root*root == num
}

// Fungsi untuk mengecek apakah suatu angka adalah kubus sempurna
func isPerfectCube(num int) bool {
	root := int(math.Round(math.Pow(float64(num), 1.0/3.0)))
	return root*root*root == num
}
