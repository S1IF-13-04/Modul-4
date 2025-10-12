package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2, x3, y3 float64

	fmt.Println("Masukkan koordinat titik A (x y):")
	fmt.Scan(&x1, &y1)
	fmt.Println("Masukkan koordinat titik B (x y):")
	fmt.Scan(&x2, &y2)
	fmt.Println("Masukkan koordinat titik C (x y):")
	fmt.Scan(&x3, &y3)

	AB := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	BC := math.Sqrt(math.Pow(x3-x2, 2) + math.Pow(y3-y2, 2))
	CA := math.Sqrt(math.Pow(x3-x1, 2) + math.Pow(y3-y1, 2))

	fmt.Printf("Sisi terpanjang: %.2f\n", math.Max(AB, math.Max(BC, CA)))
}
