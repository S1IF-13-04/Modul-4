package main

import (
	"fmt"
	"strconv"
)

func main() {
	var bil, b1, b2, b3, b4 int //temp
	fmt.Scan(&bil)
	b1 = bil % 10
	b2 = bil / 10 % 10
	b3 = bil / 100 % 10
	b4 = bil / 1000 % 10
	var b1s string = strconv.Itoa(b1)
	var b2s string = strconv.Itoa(b2)
	var b3s string = strconv.Itoa(b3)
	var b4s string = strconv.Itoa(b4)

	hasils := b1s + b2s + b3s + b4s
	hasil, _ := strconv.Atoi(hasils)

	fmt.Print(hasil)

}
