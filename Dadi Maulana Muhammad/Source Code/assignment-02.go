package main

import (
	"fmt"
	"math"
)

func main() {
	var bmi, tinggi float64
	fmt.Scan(&bmi, &tinggi)

	berat := bmi * tinggi * tinggi
	fmt.Println(int(math.Round(berat)))
}
