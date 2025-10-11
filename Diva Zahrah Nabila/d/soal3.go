package main

import (
	"fmt"
	"math"
)

func main() {
	var xA, yA, xB, yB, xC, yC float64

	fmt.Scan(&xA, &yA)
	fmt.Scan(&xB, &yB)
	fmt.Scan(&xC, &yC)

	AB := math.Sqrt(math.Pow(xB-xA, 2) + math.Pow(yB-yA, 2))
	BC := math.Sqrt(math.Pow(xC-xB, 2) + math.Pow(yC-yB, 2))
	CA := math.Sqrt(math.Pow(xA-xC, 2) + math.Pow(yA-yC, 2))

	longest := math.Max(AB, math.Max(BC, CA))
	fmt.Printf("%.2f\n", longest)
}
