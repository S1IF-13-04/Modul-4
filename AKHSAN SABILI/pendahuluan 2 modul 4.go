package main

import "fmt"

func main() {
	var hari int
	fmt.Scan(&hari)
	tahun := hari / 365
	shari := hari % 365
	minggu := shari / 7
	fmt.Println(tahun, "tahun", minggu, "minggu")
}
