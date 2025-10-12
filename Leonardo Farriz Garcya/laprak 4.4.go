package main

import "fmt"

func main() {
	var totalAwal, diskon int
	fmt.Print("Masukkan total belanja: ")
	fmt.Scan(&totalAwal)
	fmt.Print("Masukkan besar diskon (%): ")
	fmt.Scan(&diskon)

	totalAkhir := totalAwal - (totalAwal * diskon / 100)

	fmt.Println("Total setelah diskon:", totalAkhir)
}
