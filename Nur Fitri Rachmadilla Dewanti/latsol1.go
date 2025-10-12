package main

import "fmt"

func main() {
	var harga float64
	var diskon float64

	fmt.Print("Masukkan harga barang: ")
	fmt.Scan(&harga)
	fmt.Print("Masukkan persentase diskon: ")
	fmt.Scan(&diskon)

	hargaAkhir := harga - (harga * diskon / 100)

	fmt.Println("Harga setelah diskon:", hargaAkhir)
}
