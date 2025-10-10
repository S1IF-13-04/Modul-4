package main

import "fmt"

func main() {
	var bmi, tinggi float64
	var berat float64

	fmt.Print("Masukkan BMI: ")
	fmt.Scan(&bmi)

	fmt.Print("Masukkan tinggi (meter): ")
	fmt.Scan(&tinggi)

	berat = bmi * (tinggi * tinggi)
	fmt.Printf("Perkiraan berat badan: %.0f kg\n", berat)
}
