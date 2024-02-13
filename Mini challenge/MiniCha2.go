package main

import (
	"fmt"
)

func main() {
	// Variable yang berisi string kalimat
	kalimat := "selamat malam"

	// Pecah kalimat menjadi huruf-huruf
	var hurufHuruf []string
	for _, huruf := range kalimat {
		hurufHuruf = append(hurufHuruf, string(huruf))
	}

	// Inisialisasi map untuk menyimpan jumlah kemunculan huruf
	kemunculanHuruf := make(map[string]int)

	// Looping untuk menghitung kemunculan huruf
	for _, huruf := range hurufHuruf {
		kemunculanHuruf[huruf]++
	}

	// Tampilkan hasil perhitungan kemunculan huruf
	for huruf, jumlah := range kemunculanHuruf {
		fmt.Printf("%s: %d\n", huruf, jumlah)
	}
}
