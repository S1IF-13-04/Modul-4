package main

import "fmt"

func main() {
	var x, r, p, s int
	fmt.Print("Masukkan angkanya: ")
	fmt.Scanln(&x)

	r = x / 100
	p = (x / 10) % 10
	s = x % 10

	fmt.Println("Ratusan:", r, "Puluhan:", p,"Satuan:", s)
}