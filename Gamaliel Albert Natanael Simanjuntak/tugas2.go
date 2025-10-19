package main

import (
	"fmt"
	"math"
)

func main() {
	var bmi float64
	var tinggi float64

	fmt.Scan(&bmi)
	fmt.Scan(&tinggi)

	var beratFloat float64
	beratFloat = bmi * (tinggi * tinggi)

	var beratBulat int
	beratBulat = int(math.Round(beratFloat))

	fmt.Println(beratBulat)
}