package main

import "fmt"

func main() {
	var r, luas float64
	const pi = 3.141592653589793

	fmt.Print("Masukan jari-jari lingkaran: ")
	fmt.Scan(&r)

	luas = pi * r * r
	fmt.Printf("Luas lingkaran dengan jari-jari %.0f adalah: %.2f\n", r, luas)

}
