package main

import "fmt"

func main() {
	var bmi, tinggibadan, beratbadan float64
	fmt.Print("Nilai BMI: ")
	fmt.Scan(&bmi)
	fmt.Print("Tinggi Badan (m): ")
	fmt.Scan(&tinggibadan)

	beratbadan = bmi * tinggibadan * tinggibadan 
	berat := int(beratbadan + 0.5)

	fmt.Printf("Berat Badan (Kg): %d", berat)
}