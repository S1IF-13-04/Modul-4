package main

import "fmt"

func main() {
	var total, discount int
	fmt.Scanln(&total)
	fmt.Scanln(&discount)

	totalakhir := total - (total * discount / 100)
	fmt.Println(totalakhir)

}
