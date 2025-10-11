package main

import ("fmt"
		"math")

func main() {
	var r float64
	fmt.Printf("Masukan Jari-jari : ")
	fmt.Scan(&r)

	luas := math.Pi * r * r

	fmt.Printf("Luas lingkaran adalah: %.2f\n", luas)
}