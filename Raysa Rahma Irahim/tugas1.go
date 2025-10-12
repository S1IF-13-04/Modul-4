package main

import "fmt"

func main() {
	var harga, diskon, total float64

	fmt.Print("Masukkan total belanja awal:")
	fmt.Scan(&harga)
	fmt.Print("Masukkan diskon (%):")
	fmt.Scan(&diskon)
	total = harga * (1 - diskon/100)

	fmt.Println("harga setelah diskon: ", total)
}
