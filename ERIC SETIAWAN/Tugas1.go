package main

import "fmt"

func main() {

	var harga, diskon int
	fmt.Print("Masukkan Total Belanja: ")
	fmt.Scan(&harga)
	fmt.Print("Masukkan Diskon: ")
	fmt.Scan(&diskon)

	diskon = harga * diskon / 100
	total := harga - diskon

	fmt.Printf("Harga setelah diskon adalah %d", total)
}
