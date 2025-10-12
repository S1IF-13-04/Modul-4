package main

import "fmt"

func main() {
	var tahunLahir, tahunSekarang, umur int

	fmt.Print("Masukkan tahun lahir: ")
	fmt.Scan(&tahunLahir)

	fmt.Print("Masukkan tahun sekarang: ")
	fmt.Scan(&tahunSekarang)

	umur = tahunSekarang - tahunLahir

	fmt.Printf("Umur kamu sekarang adalah: %d tahun\n", umur)
}
