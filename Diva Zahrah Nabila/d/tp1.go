package main

import "fmt"

func main() {
	var a, b, c, nim, kelas string
	fmt.Print("Masukkan nama (3 kata): ")
	fmt.Scan(&a, &b, &c)
	fmt.Print("Masukkan NIM: ")
	fmt.Scan(&nim)
	fmt.Print("Masukkan Kelas: ")
	fmt.Scan(&kelas)

	namaLengkap := a + " " + b + " " + c
	fmt.Println("\nNAMA  :", namaLengkap)
	fmt.Println("NIM   :", nim)
	fmt.Println("KELAS :", kelas)
}
