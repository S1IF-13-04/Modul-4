package main

import "fmt"

func main() {
	var in, p1, p2, p3 int

	fmt.Scan(&in)

	p1 = in % 10
	p2 = (in / 10) % 10
	p3 = in / 100

	fmt.Println(p3 < p2 && p2 < p1)
}
