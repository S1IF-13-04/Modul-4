package main

import "fmt"

func main() {
	var totalAwal, diskon int
	fmt.Print("Masukkan total belanja awal: ")
	fmt.Scan(&totalAwal)
	fmt.Print("Masukkan besarnya diskon (%): ")
	fmt.Scan(&diskon)

	totalAkhir := totalAwal - (totalAwal * diskon / 100)

	fmt.Println("Total belanja akhir setelah diskon adalah:", totalAkhir)
}
