package main

import (
	"fmt"
)

func main() {
	// Membaca input dari pengguna
	fmt.Print("Masukkan kalimat: ")
	var kalimat string
	fmt.Scanln(&kalimat)

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

	// Konversi map menjadi slice
	var hasil []string
	for huruf, jumlah := range kemunculanHuruf {
		hasil = append(hasil, fmt.Sprintf("%s: %d", huruf, jumlah))
	}

	// Tampilkan hasil perhitungan kemunculan huruf
	fmt.Println(hasil)
}
