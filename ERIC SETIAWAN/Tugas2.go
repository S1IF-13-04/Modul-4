package main

import "fmt"

func main() {
	var nilai, tinggiBadan float64
	fmt.Print("Masukkan Nilai BMI: ")
	fmt.Scan(&nilai)
	fmt.Print("Masukkan Tinggi Badan: ")
	fmt.Scan(&tinggiBadan)
	beratBadan := nilai * (tinggiBadan * tinggiBadan)
	fmt.Printf("Berat Badan: %.0f", beratBadan)
}
