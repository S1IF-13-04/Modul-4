package main

import "fmt"

func main() {
	var Nama string
	var Nim int
	var Kelas string

	fmt.Printf("Nama :")
	fmt.Scan(&Nama)

	fmt.Printf("NIM :")
	fmt.Scan(&Nim)

	fmt.Printf("Kelas :")
	fmt.Scan(&Kelas)

	fmt.Println("Perkenalkan saya adalah", Nama,
		"salah satu mahasiswa Prodi S1-IF dari kelas", Kelas,
		"dengan NIM", Nim)
}