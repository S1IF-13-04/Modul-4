package main

import "fmt"

func main() {
	var bmi, tinggi, berat float64
	fmt.Print("Masukkan nilai BMI: ")
	fmt.Scan(&bmi)
	fmt.Print("Masukkan tinggi badan (meter): ")
	fmt.Scan(&tinggi)
	berat = bmi * (tinggi * tinggi)
	fmt.Printf("Berat badan adalah: %.2f kg\n", berat)
}
