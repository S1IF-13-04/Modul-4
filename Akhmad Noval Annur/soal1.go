package main

import "fmt"

func main() {
	var n1, n2, n3 float64
	fmt.Print("Masukkan nilai 1: ")
	fmt.Scan(&n1)
	fmt.Print("Masukkan nilai 2: ")
	fmt.Scan(&n2)
	fmt.Print("Masukkan nilai 3: ")
	fmt.Scan(&n3)

	rata := (n1 + n2 + n3) / 3
	fmt.Printf("Rata-rata: %.2f\n", rata)

	if rata >= 60 {
		fmt.Println("Keterangan: Lulus")
	} else {
		fmt.Println("Keterangan: Tidak Lulus")
	}
}