package main

import "fmt"

func main() {
	var hargaBelanja, Diskon, potongan, hargaAkhir int
	fmt.Print("Masukan harga barang: ")
	fmt.Scan(&hargaBelanja)
	fmt.Print("Masukan diskon: ")
	fmt.Scan(&Diskon)
	potongan = hargaBelanja * Diskon / 100
	hargaAkhir = hargaBelanja - potongan
	fmt.Println("harga:", hargaAkhir)

}
