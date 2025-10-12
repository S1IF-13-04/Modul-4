package main

import (
	"fmt"
)

func main() {
	var x, fx float64
	var luas int
	fmt.Print("Masukkan nilai x : ")
	fmt.Scanln(&x)
	fx = (x * x * x) + 1/x
	fmt.Println("Maka hasil persamaan f(x) adalah ", fx)
	fxInt := int(fx)
	luas = fxInt * fxInt
	fmt.Print("Maka luas persegi adalah ", luas)
}
