package main

import "fmt"

func main() {
	var nominal int
	fmt.Print("Masukkan nominal: ")
	fmt.Scan(&nominal)

	a := nominal / 1000
	b := (nominal % 1000) / 500
	c := (nominal % 500) / 200
	d := (nominal % 200) / 100

	fmt.Println(a, "lembar 1000,", b, "koin 500,", c, "koin 200,", d, "koin 100")
}
