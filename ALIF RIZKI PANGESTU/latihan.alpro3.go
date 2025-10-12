package main

import (
	"fmt"
)

func main() {
	var berat, tinggi, bmi float64
	fmt.Print("Masukkan berat dan tinggi: ")
	fmt.Scan(&berat)
	fmt.Scan(&tinggi)

	bmi = berat / (tinggi * tinggi)

	fmt.Printf("Nilai BMI Anda adalah: %.2f\n", bmi)
}
