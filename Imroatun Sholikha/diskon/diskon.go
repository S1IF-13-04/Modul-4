package main

import "fmt"

func main() {
	var hasil, diskon, totalAwal int

	fmt.Print("total awal:")
	fmt.Scan(&totalAwal)
	fmt.Print("diskon:")
	fmt.Scan(&diskon)

	hasil = totalAwal - (totalAwal * diskon / 100)

	fmt.Println(hasil)
}
