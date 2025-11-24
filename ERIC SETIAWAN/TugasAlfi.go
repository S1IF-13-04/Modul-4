package main

import "fmt"

func main() {
	var p, l, t float64
	fmt.Print("Masukkan panjang (p) dalam meter: ")
	fmt.Scan(&p)
	fmt.Print("Masukkan lebar (l) dalam meter: ")
	fmt.Scan(&l)
	fmt.Print("Masukkan tinggi (t) dalam meter: ")
	fmt.Scan(&t)
	volume := p * l * t
	liter := volume * 1000
	fmt.Printf("Volume air yang dapat dimasukkan adalah %.2f liter\n", liter)
}
