package main

import "fmt"

func main() {
	var totalBelanja, diskonPersen int
	fmt.Scan(&totalBelanja, &diskonPersen)
	diskon := totalBelanja * diskonPersen / 100
	totalAkhir := totalBelanja - diskon
	fmt.Println(totalAkhir)
}