package main

import "fmt"

func main() {
	var belanja int
	var diskon int

	fmt.Scan(&belanja)
	fmt.Scan(&diskon)

	belanjaFloat := float64(belanja)
	diskonFloat := float64(diskon)

	jumlah := belanjaFloat * (diskonFloat / 100.0)
	harga := belanjaFloat - jumlah

	fmt.Println(int(harga))
}
