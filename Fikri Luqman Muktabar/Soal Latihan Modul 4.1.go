package main

import "fmt"

func main() {
	var harga, diskon, totaldiskon, setelahdiskon float64
	fmt.Scan(&harga)
	fmt.Scan(&diskon)
	
	totaldiskon = harga * (diskon / 100.0)
	setelahdiskon = harga - totaldiskon
	
	fmt.Println("Total belanja akhir setelah dipotong diskon: ", setelahdiskon)
}