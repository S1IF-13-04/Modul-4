package main

import (
	"fmt"
)

func main() {
	var totalDetik int
	fmt.Scan(&totalDetik)

	hari := totalDetik / 86400
	jam := (totalDetik % 86400) / 3600
	menit := (totalDetik % 3600) / 60
	detik := totalDetik % 60

	fmt.Printf("%d hari, %d jam, %d menit, %d detik\n", hari, jam, menit, detik)
}
