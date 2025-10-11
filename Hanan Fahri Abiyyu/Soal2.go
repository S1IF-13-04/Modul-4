package main

import (
	"fmt"
)

func main() {
	var bmi, tinggibadan, beratbadan float64
	fmt.Print("Masukkan BMI :")
	fmt.Scan(&bmi)
	fmt.Print("Masukkan tinggi badan :")
	fmt.Scan(&tinggibadan)

	beratbadan = bmi * tinggibadan * tinggibadan
	berat := int(beratbadan + 0.5)
	fmt.Print("Hasil berat badan : ", berat)

}
