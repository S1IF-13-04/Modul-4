package main

import "fmt"

func main() {
	var nama string
	var n1, n2, n3 int

	fmt.Print("Masukkan nama dan tiga nilai ujian: ")
	fmt.Scan(&nama, &n1, &n2, &n3)

	jumlah := n1 + n2 + n3
	rata := float64(jumlah) / 3
	rataBulat := int(rata + 0.5)
	sisa := jumlah % 3

	fmt.Printf("%s, %d, %.2f, %d, %d\n", nama, jumlah, rata, rataBulat, sisa)
}