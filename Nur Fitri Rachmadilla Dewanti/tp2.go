package main

import "fmt"

func main() {
	var hari int
	fmt.Print("Masukkan jumlah hari: ")
	fmt.Scan(&hari)

	tahun := hari / 365
	sisa := hari % 365
	bulan := sisa / 30
	hariSisa := sisa % 30

	fmt.Printf("Umur %d tahun %d bulan %d hari\n", tahun, bulan, hariSisa)
}
