package main

import (
	"fmt"
	"os"
)

// Struct untuk menyimpan data teman
type Teman struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

// Data teman-teman di kelas
var temanTeman = []Teman{
	{"Fitri", "Jakarta", "Developer", "Ingin mempelajari Golang untuk pengembangan aplikasi backend."},
	{"Budi", "Bandung", "Designer", "Tertarik dengan kemampuan konkurensi dan kinerja yang cepat pada Golang."},
	{"Cindy", "Surabaya", "Data Analyst", "Ingin menggunakan Golang untuk analisis data skala besar."},
	{"Eko", "Malang", "DevOps Engineer", "Tertarik dengan deployment aplikasi menggunakan Golang."},
}

// Function untuk mencari teman berdasarkan nama
func cariTeman(nama string) (Teman, bool) {
	for _, teman := range temanTeman {
		if teman.Nama == nama {
			return teman, true
		}
	}
	return Teman{}, false
}

func main() {
	// Mendapatkan argumen dari CLI
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Harap berikan nama teman sebagai argumen CLI.")
		return
	}

	// Mengambil nama teman dari argumen CLI
	namaTeman := args[1]

	// Mencari teman berdasarkan nama
	teman, found := cariTeman(namaTeman)
	if !found {
		fmt.Printf("Data dengan nama %s tidak tersedia.\n", namaTeman)
		return
	}

	// Menampilkan data teman
	fmt.Println("Nama:", teman.Nama)
	fmt.Println("Alamat:", teman.Alamat)
	fmt.Println("Pekerjaan:", teman.Pekerjaan)
	fmt.Println("Alasan memilih kelas Golang:", teman.Alasan)
}
