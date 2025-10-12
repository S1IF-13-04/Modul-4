package main

import (
	"fmt"
	"math"
)

func main() {
	var bmi, tinggi, berat float64

	fmt.Print("Masukan nilai:")
	fmt.Scan(&bmi, &tinggi)

	berat = bmi * math.Pow(tinggi, 2)

	fmt.Printf("%.0f\n", berat)
}
