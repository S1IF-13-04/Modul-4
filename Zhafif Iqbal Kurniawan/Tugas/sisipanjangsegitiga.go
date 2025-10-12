package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2, x3, y3 float64

	fmt.Scan(&x1, &y1)
	fmt.Scan(&x2, &y2)
	fmt.Scan(&x3, &y3)

	AB := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	BC := math.Sqrt(math.Pow(x3-x2, 2) + math.Pow(y3-y2, 2))
	AC := math.Sqrt(math.Pow(x3-x1, 2) + math.Pow(y3-y1, 2))

	sisiterpanjang := math.Max(math.Max(AB, BC), AC)

	fmt.Printf("%.2f\n", sisiterpanjang)
}
