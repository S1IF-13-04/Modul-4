package main

import "fmt"

func main() {
	var totalBelanja, diskonPersen int
	var totalAkhir float64

	fmt.Print("Masukkan total belanja: ")
	fmt.Scan(&totalBelanja)
	fmt.Print("Masukkan besaran diskon (%): ")
	fmt.Scan(&diskonPersen)

	totalAkhir = float64(totalBelanja) - (float64(totalBelanja) * float64(diskonPersen) / 100)

	fmt.Printf("Total belanja setelah diskon: %.0f\n", totalAkhir)
}
