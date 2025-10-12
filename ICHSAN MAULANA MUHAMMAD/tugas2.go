package main

import (
	"fmt"
	"math"
)

func main() {
	var BMI, tinggiBadan, beratBadan, beratBadanAkhir float64

	fmt.Scan(&BMI, &tinggiBadan)

	beratBadan = (tinggiBadan * tinggiBadan) * BMI
	beratBadanAkhir = math.Ceil(beratBadan)

	fmt.Println(beratBadanAkhir, "kg")
}
