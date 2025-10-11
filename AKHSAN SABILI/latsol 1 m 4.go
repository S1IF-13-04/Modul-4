package main

import "fmt"

func main() {
	var d, m, j int
	fmt.Scan(&d)
	j = d / 3600
	m = (d % 3600) / 60
	d = d % 60
	fmt.Printf("%d jam %d menit dan %d detik", j, m, d)
}
