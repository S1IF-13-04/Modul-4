package main

import "fmt"

func main() {
	var bmi, tinggi float64
	var berat float64
	fmt.Print("masukan bmi: ")
	fmt.Scan(&bmi)
	fmt.Print("masukan tinggi: ")
	fmt.Scan(&tinggi)
	berat = bmi * (tinggi * tinggi)
	fmt.Printf("%.0f kg\n", berat)
}
