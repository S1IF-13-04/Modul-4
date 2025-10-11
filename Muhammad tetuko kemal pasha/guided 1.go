package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Print("masukan detik: ")
	fmt.Scan(&x)
	y = x / 3600
	z = (x % 3600) / 60
	x = x % 60
	fmt.Printf("%d jam, %d menit, %d detik\n", y, z, x)
}
