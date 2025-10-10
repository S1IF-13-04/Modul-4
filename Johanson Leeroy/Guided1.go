package main

import "fmt"

func main() {
	var totaldetik, jam, menit, detik, sisadetik int

	fmt.Scan(&totaldetik)

	jam = totaldetik / 3600
	sisadetik = totaldetik % 3600
	menit = sisadetik / 60
	detik = sisadetik % 60

	fmt.Println(jam, "jam", menit, "menit", detik, "detik")
}
