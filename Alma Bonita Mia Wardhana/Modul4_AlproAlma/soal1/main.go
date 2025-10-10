package main

import "fmt"

func main() {
	var beli, diskon int
	var hasil int

	fmt.Println("Masukkan harga beli:")
	fmt.Scan(&beli)

	fmt.Print("Masukkan diskon (%): ")
	fmt.Scan(&diskon)

	hasil = beli * diskon / 100
	hasil2 := beli - hasil

	fmt.Println("Harga setelah diskon adalah:", hasil2)
}
