package main

import "fmt"

func main() {

	var A, B int

	fmt.Scan(&A, &B)

	Af := float64(A)
	Bf := float64(B)

	rata_rata := (Af + Bf) / 2

	fmt.Print(rata_rata)

}
