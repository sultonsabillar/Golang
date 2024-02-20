package main

import (
	"fmt"
	"math"
)

type hitung interface {
	luas() float64
	keliling() float64
}

type persegi struct {
	sisi float64
}

type lingkaran struct {
	diameter float64
}

func (l lingkaran) jariJariMethod() float64 {
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJariMethod(), 2)
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}

func main() {
	var lingkaran1 hitung = lingkaran{diameter: 14}
	var persegi1 hitung = persegi{5}

	fmt.Printf("Type of lingkaran: %T \n", lingkaran1)
	fmt.Printf("Type of persegi: %T \n", persegi1)

	fmt.Println("===== lingkaran")
	fmt.Println("jari-jari :", lingkaran1.(lingkaran).jariJariMethod())
	fmt.Println("luas      :", lingkaran1.luas())
	fmt.Println("keliling  :", lingkaran1.keliling())

	fmt.Println("===== persegi")
	fmt.Println("luas      :", persegi1.luas())
	fmt.Println("keliling  :", persegi1.keliling())
}
