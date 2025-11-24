package main

import "fmt"

func main() {
	var kecepatan, jarak float64

	fmt.Print("Masukkan kecepatan (km/jam): ")
	fmt.Scan(&kecepatan)
	fmt.Print("Masukkan jarak (km): ")
	fmt.Scan(&jarak)

	waktu := jarak / kecepatan

	fmt.Printf("Waktu yang dibutuhkan adalah %.2f jam\n", waktu)
}
