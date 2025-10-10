package main

import "fmt"

func main() {
	var a, b, c, d, e int
	var rata float64

	fmt.Print("Masukkan 5 bilangan bulat: ")
	fmt.Scan(&a, &b, &c, &d, &e)

	rata = float64(a+b+c+d+e) / 5

	fmt.Printf("Rata-rata = %.1f\n", rata)
}
