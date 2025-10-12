package main

import "fmt"

func main() {
	var berat, air int
	fmt.Print("Masukkan berat badan (kg): ")
	fmt.Scan(&berat)

	air = berat * 35
	fmt.Printf("Kebutuhan air harian anda adalah: %d ml\n", air)
}
