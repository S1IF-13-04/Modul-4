package main

import "fmt"

func main() {
	var meter float64
	var kilometer float64

	fmt.Print("Masukkan jarak (meter): ")
	fmt.Scan(&meter)

	kilometer = meter / 1000

	fmt.Printf("%.2f km\n", kilometer)
}
