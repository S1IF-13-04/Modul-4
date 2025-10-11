package main

import "fmt"

func main() {
	var x int
	var d1, d2, d3 int
	fmt.Print("Masukkan Input :")
	fmt.Scanln(&x)

	d1 = x / 100
	d2 = (x % 100) / 10
	d3 = x % 100 % 10

	fmt.Print(d1 <= d2 && d2 <= d3)

}
