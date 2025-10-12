package main

import "fmt"

func main() {
	var hargaBelanja, Diskon, potongan, hargaAkhir int

	fmt.Print("Masukkan harga barang: ")
	fmt.Scan(&hargaBelanja)

	fmt.Print("Masukkan diskon: ")
	fmt.Scan(&Diskon)

	potongan = hargaBelanja * Diskon / 100
	hargaAkhir = hargaBelanja - potongan

	fmt.Println("Harga akhir:", hargaAkhir)
}
