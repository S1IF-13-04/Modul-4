package main

import "fmt"

func main() {
	var totalDetik, jam, menit, detik, sisaDetik int
	fmt.Print("Masukkan total detik: ")
	fmt.Scan(&totalDetik)
	jam = totalDetik / 3600
	sisaDetik = totalDetik % 3600
	menit = sisaDetik / 60
	detik = sisaDetik % 60
	fmt.Println(jam, "jam,", menit, "menit,", detik, "detik")
}
