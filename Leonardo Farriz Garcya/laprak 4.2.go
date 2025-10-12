package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan bilangan 3 digit: ")
	fmt.Scan(&n)

	d1 := n / 100
	d2 := (n / 10) % 10
	d3 := n % 10 

	hasil := d1 < d2 && d2 < d3

	fmt.Println(hasil)
}
