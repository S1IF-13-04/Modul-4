package main

import "fmt"

func main() {
	var total int
	fmt.Scan(&total)

	const (
		tahun = 365
		bulan = 30
	)

	t := total / tahun
	b := (total % tahun) / bulan
	h := (total % tahun) % bulan

	fmt.Println(t, "Tahun,", b, "Bulan,", h, "Hari")
}
