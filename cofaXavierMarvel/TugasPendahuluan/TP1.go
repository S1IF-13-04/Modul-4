package main

import (
	"fmt"
)

func main() {

	var A, B int

	fmt.Scan(&A, &B)

	C := A / B

	D := A % B

	fmt.Printf("\nHasil pembarian bulat : %v\nHasil sisa bagi : %v", C, D)
}
