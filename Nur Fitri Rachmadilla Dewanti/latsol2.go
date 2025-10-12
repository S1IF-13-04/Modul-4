package main

import (
	"fmt"
	"math"
)

func main() {
	var bmi, tinggi float64

	fmt.Print("Masukkan nilai BMI: ")
	fmt.Scan(&bmi)
	fmt.Print("Masukkan tinggi badan (meter): ")
	fmt.Scan(&tinggi)

	berat := bmi * (tinggi * tinggi)
	hasilBerat := math.Round(berat) 

	fmt.Println("Berat badan (kg):", hasilBerat)
}
