package main

import "fmt"

func main() {
	var a, b, c float64
	fmt.Print("Masukkan 3 nilai: ")
	fmt.Scan(&a, &b, &c)

	rata := (a + b + c) / 3
	fmt.Println("Rata-rata:", int(rata+0.5))
}
