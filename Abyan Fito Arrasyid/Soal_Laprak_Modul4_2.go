package main

import "fmt"

func main() {
	var BMI float64
	var TinggiBadan float64

	fmt.Print("Masukkan Berat Badan dan Tinggi Badan: ")
	fmt.Scan(&BMI, &TinggiBadan)

	BMI = BMI * (TinggiBadan * TinggiBadan)

	fmt.Printf("%.0f", BMI)
}
