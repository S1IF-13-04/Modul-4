package main

import "fmt"

func main() {
	var harga int
	var totalHarga float64

	fmt.Print("Masukkan harga makanan: ")
	fmt.Scan(&harga)

	totalHarga = float64(harga) * 1.15
	fmt.Printf("Total harga setelah pajak 15%%: %.0f\n", totalHarga)
}
