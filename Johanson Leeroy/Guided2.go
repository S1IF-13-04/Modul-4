package main

import "fmt"

func main() {
	var bil, b1, b2, b3 int
	var cek bool
	fmt.Scan(&bil)

	b1 = bil / 100
	b2 = bil % 100 / 10
	b3 = bil % 10
	cek = b1 <= b2 && b2 <= b3
	fmt.Println(cek)
}
