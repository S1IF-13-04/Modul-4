package main

import (
	"fmt"
)

func main() {
	var h, d, td, sd float64
	fmt.Print("Masukkan harga total belanja :")
	fmt.Scanln(&h)
	fmt.Print("Masukkan diskon :")
	fmt.Scanln(&d)

	td = h * (d / 100.0)
	sd = h - td
	fmt.Println("Total pembayaran setelah diskon :", sd)
}
