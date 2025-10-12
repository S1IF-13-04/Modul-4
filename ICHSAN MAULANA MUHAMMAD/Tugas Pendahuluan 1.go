package main

import "fmt"

func main() {
	var bulan, minggu, hari, jam, menit, detik int
	fmt.Scan(&detik)

	bulan = detik / 2592000
	minggu = (detik % 2592000) / 604800
	hari = (detik % 604800) / 86400
	jam = detik % 86400 / 3600
	menit = detik % 3600 / 60
	detik = detik % 60

	fmt.Println(bulan, "bulan", minggu, "minggu", hari, "hari", jam, "jam", menit, "menit", detik, "detik")
}
