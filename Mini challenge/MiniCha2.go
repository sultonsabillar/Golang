package main

import (
	"fmt"
)

func main() {
	// Membaca input dari pengguna
	fmt.Print("Masukkan kalimat: ")
	var kalimat string
	fmt.Scanln(&kalimat)

	// Inisialisasi map untuk menyimpan jumlah kemunculan huruf
	kemunculanHuruf := make(map[rune]int)

	// Looping untuk menghitung kemunculan huruf
	for _, huruf := range kalimat {
		kemunculanHuruf[huruf]++
	}

	// Tampilkan huruf satu per satu ke bawah
	for _, huruf := range kalimat {
		fmt.Println(string(huruf))
	}

	// Tampilkan hasil perhitungan kemunculan huruf
	fmt.Println()
	fmt.Println("map", kemunculanHuruf)
}
