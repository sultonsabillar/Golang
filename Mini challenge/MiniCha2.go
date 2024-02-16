package main

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

	// Tampilkan hasil perhitungan kemunculan huruf
	fmt.Println("Hasil perhitungan kemunculan huruf:")
	for huruf, jumlah := range kemunculanHuruf {
		fmt.Printf("%s: %d\n", huruf, jumlah)
	}
}
package main

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
