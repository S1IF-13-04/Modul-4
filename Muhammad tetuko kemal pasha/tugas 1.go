package main

import "fmt"

func main() {
	var beli, diskon int
	var hasil int
	fmt.Println("masukan harga beli dan diskon: ")
	fmt.Scan(&beli)
	fmt.Printf("masukan diskon: ")
	fmt.Scan(&diskon)
	hasil = beli * diskon / 100
	hasil2 := beli - hasil
	fmt.Print(hasil2)
}
