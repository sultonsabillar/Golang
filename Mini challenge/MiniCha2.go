package main

import "fmt"

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

	// Tampilkan hasil perhitungan kemunculan huruf
	fmt.Println("Hasil perhitungan kemunculan huruf:")
	for huruf, jumlah := range kemunculanHuruf {
		fmt.Printf("%c: %d\n", huruf, jumlah)
	}
}
