package main

import "fmt"

func main() {
	var bmi, tb, bb float64
	fmt.Print("Masukan BMI: ")
	fmt.Scan(&bmi)
	fmt.Print("Masukan tinggi dama meter: ")
	fmt.Scan(&tb)

	bb = bmi * tb * tb

	berat := int(bb + 0.5)

	fmt.Printf("Berat badanya adalah: %d",berat)
}