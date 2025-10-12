package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2, x3, y3 float64
	var AB, BC, AC, ABsqr, BCsqr, ACsqr, BCakr float64

	fmt.Scan(&x1, &y1)
	fmt.Scan(&x2, &y2)
	fmt.Scan(&x3, &y3)

	ABsqr = math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2)
	BCsqr = math.Pow(x3-x2, 2) + math.Pow(y3-y2, 2)
	ACsqr = math.Pow(x3-x1, 2) + math.Pow(y3-y1, 2)

	AB = math.Sqrt(ABsqr)
	BC = math.Sqrt(BCsqr)
	AC = math.Sqrt(ACsqr)

	BCakr = (AB + BC + math.Sqrt(math.Pow(AB-BC, 2))) / 2
	BCakr = (BCakr + AC + math.Sqrt(math.Pow(BCakr-AC, 2))) / 2

	fmt.Printf("%.2f", BCakr)
}
