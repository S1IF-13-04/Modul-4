package main

import "fmt"

func main() {
	var harga, diskon, totaldiskon, setelahdiskon float64
	fmt.Print("Masukan total belanja : ")
	fmt.Scan(&harga)
	fmt.Print("Masukan total diskon : ")
	fmt.Scan(&diskon)
	totaldiskon = harga * (diskon/100.0)
	setelahdiskon = harga - totaldiskon
	fmt.Println("Total pembayaran adalah : ",setelahdiskon)
}
