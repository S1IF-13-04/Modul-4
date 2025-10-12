package main

import "fmt"

func main() {
	var nama string
	var umur int
	
	fmt.Print("Masukkan nama dan umur: ")
	fmt.Scan(&nama, &umur)
	
	bulan := umur * 12
	bagi := float64(umur) / 5
	sisa := umur % 5
	genap := umur%2 == 0
	
	fmt.Printf("%s, %d tahun, %d, %.2f, %d sisa %d, %t\n", nama, umur, bulan, bagi, umur/5, sisa, genap)
}