package main

import (
	"fmt"
	"strings"
)

func main() {
	// Variable yang berisi string kalimat
	kalimat := "Halo teman-teman, apa kabar hari ini?"

	// Pecah kalimat menjadi kata-kata
	kataKata := strings.Fields(kalimat)

	// Inisialisasi map untuk menyimpan jumlah kemunculan kata
	kemunculanKata := make(map[string]int)

	// Looping untuk menghitung kemunculan kata
	for _, kata := range kataKata {
		kemunculanKata[kata]++
	}

	// Tampilkan hasil perhitungan kemunculan kata
	for kata, jumlah := range kemunculanKata {
		fmt.Printf("Kata: %s, Jumlah: %d\n", kata, jumlah)
	}
}
