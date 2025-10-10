package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2, x3, y3, AB, AC, ABsqr, ACsqr, BC float64
	fmt.Println("Masukan titik koordinat segitiga (siku-siku) ABC :")
	fmt.Scan(&x1, &y1)
	fmt.Scan(&x2, &y2)
	fmt.Scan(&x3, &y3)
	AB = x3 - x1
	AC = y3 - y1
	ABsqr = AB * AB
	ACsqr = AC * AC
	BC = ABsqr + ACsqr
	BCakr := math.Sqrt(BC)
	fmt.Printf("%.2f", BCakr)
}
