package main

import (
	"fmt"
)

func main() {
	var j, d, m int
	fmt.Print("Masukkan total detik :")
	fmt.Scanln(&d)

	j = d / 3600
	m = (d % 3600) / 60
 	d = d % 60

	fmt.Println(j, "jam", m, "menit", d, "detik")
}
