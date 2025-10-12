package main

import "fmt"

func main() {
	var hb, d, o, td, hd, ht float64
	fmt.Print("Masukan total belanja : ")
	fmt.Scan(&hb)
	fmt.Print("Masukan total diskon : ")
	fmt.Scan(&d)
	fmt.Print("Masukan harga ongkir : ")
	fmt.Scan(&o)
	td = hb * (d/100.0)
	hd = hb - td
	ht = hd + o
	fmt.Println("Total pembayaran adalah : ",ht)
}
