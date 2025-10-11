package main

import "fmt"

func main() {
	var h, d int
	fmt.Scan(&h, &d)
	diskon := ( h * d / 100)
	total := ( h - diskon )
	fmt.Println(total)
}