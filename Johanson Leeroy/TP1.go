package main

import "fmt"

func main() {
	var bil, b1, b2, b3, b4, jumlah int
	fmt.Scan(&bil)
	b1 = bil % 10
	b2 = bil / 10 % 10
	b3 = bil / 100 % 10
	b4 = bil / 1000
	jumlah = b1 + b2 + b3 + b4
	fmt.Print(jumlah)

}
