package main

import "fmt"

func main() {
	var angka int
	fmt.Print("Masukkan bilangan (100-999): ")
	fmt.Scan(&angka)

	puluhan := (angka / 10) % 10
	fmt.Println(puluhan%2 == 0)
}
