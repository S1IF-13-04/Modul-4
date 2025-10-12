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

	d1 := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	d2 := math.Sqrt(math.Pow(x3-x2, 2) + math.Pow(y3-y2, 2))
	d3 := math.Sqrt(math.Pow(x3-x1, 2) + math.Pow(y3-y1, 2))

	longest := math.Max(d1, math.Max(d2, d3))

	fmt.Printf("%.2f\n", longest)
}
