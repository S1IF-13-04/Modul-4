package main

import "fmt"

func main() {
	var totalBelanja int
	var uangTunai int

	fmt.Print("Total Belanja: ")
	fmt.Scanln(&totalBelanja)

	fmt.Print("Uang Tunai: ")
	fmt.Scanln(&uangTunai)

	kembalian := uangTunai - totalBelanja
	poinLoyalitas := totalBelanja / 10000

	fmt.Println("\n--- Hasil Transaksi ---")
	fmt.Printf("Kembalian Anda: Rp %d\n", kembalian)
	fmt.Printf("Poin Loyalitas Didapat: %d poin\n", poinLoyalitas)
}