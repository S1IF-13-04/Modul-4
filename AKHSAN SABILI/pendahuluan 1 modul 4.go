package main

import "fmt"

func main() {
	var h, d int
	fmt.Print("harga barang : ")
	fmt.Scan(&h)
	fmt.Print("Diskon : ")
	fmt.Scan(&d)
	diskon := h * d / 100
	hakhir := h - diskon
	fmt.Println("Harga setelah diskon : ", hakhir)

}
