package main

import "fmt"

func main() {

	var idr int
	var sgd float64

	fmt.Scan(&idr)

	rp := float64(idr)

	sgd = rp / 12732.72

	fmt.Printf("%.2f", sgd)
}
