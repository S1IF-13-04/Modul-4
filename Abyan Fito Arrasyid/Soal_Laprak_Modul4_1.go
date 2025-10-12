package main

import "fmt"

func main() {
	var x float64
	var diskon float64

	fmt.Println("Masukkan harga barang: ")
	fmt.Scan(&x)
	fmt.Println("Masukkan diskon: ")
	fmt.Scan(&diskon)

	diskon = x - (x * diskon / 100)

	fmt.Printf("Harga: %.0f", diskon)

}
