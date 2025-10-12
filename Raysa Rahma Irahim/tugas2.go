package main

import "fmt"

func main() {
	var tinggiBadan, bmi, beratBadan float64
	
	fmt.Print("Masukkan nilai bmi: ")
	fmt.Scan(&bmi)
	fmt.Print("Masukkan tinggi badan: ")
	fmt.Scan(&tinggiBadan)

	beratBadan = bmi * (tinggiBadan * tinggiBadan)

	fmt.Printf("Berat badan: %0.f kg\n", beratBadan)
}
