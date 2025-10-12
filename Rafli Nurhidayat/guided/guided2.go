package main

import "fmt"

func main() {
	var bilangan, d1, d2, d3 int
	var status bool
	fmt.Scan(&bilangan)
	d1 = bilangan / 100
	d2 = bilangan % 100 / 10
	d3 = bilangan % 100 % 10

	status = d1 <= d2 && d2 <= d3
	fmt.Print(status)
}
