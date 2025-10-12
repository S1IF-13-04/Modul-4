package main

import (
	"fmt"
	"math"
)

func main() {
	var lama float64
	fmt.Print("Masukkan lama parkir (jam): ")
	fmt.Scan(&lama)

	jam := math.Ceil(lama)

	boolToFloat := func(b bool) float64 {
		return map[bool]float64{true: 1, false: 0}[b]
	}

	biaya := 3000 + (jam-1)*2000*boolToFloat(jam > 1)

	fmt.Printf("Total biaya parkir: Rp %.0f\n", biaya)
}