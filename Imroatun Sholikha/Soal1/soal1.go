package main

import "fmt"

func main() {
	var nama string
	var s1, s2, s3 int

	fmt.Print("input:")
	fmt.Scan(&nama, &s1, &s2, &s3)

	jumlah := s1 + s2 + s3
	rata := float64(jumlah) / 3.0
	sisa := jumlah % 5

	fmt.Printf("%s %d %.2f %d", nama, jumlah, rata, sisa)
}
