package main

import "fmt"

func main() {
	var nama string
	var waktu int

	fmt.Print("Input:")
	fmt.Scan(&nama, &waktu)

	detik := waktu * 60
	bagi := detik / 7
	sisa := detik % 7
	genap := bagi%2 == 0

	fmt.Println(nama, detik, "detik", bagi, "sisa", sisa, genap)
}
