package main

import "fmt"

func main() {
	var panjang, lebar float64

	fmt.Print("Masukkan panjang: ")
	fmt.Scan(&panjang)
	fmt.Print("Masukkan lebar: ")
	fmt.Scan(&lebar)
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)

	fmt.Printf("\nLuas persegi panjang = %.2f\n", luas)
	fmt.Printf("Keliling persegi panjang = %.2f\n", keliling)
}
