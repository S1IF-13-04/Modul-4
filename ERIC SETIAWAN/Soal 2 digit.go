package main

import "fmt"

func main() {
	var bilangan, digit1, digit2, digit3 int
	fmt.Scan(&bilangan)

	digit1 = bilangan / 100
	digit2 = bilangan % 100 / 10
	digit3 = bilangan % 100 % 10

	fmt.Println(digit1 <= digit2 && digit2 <= digit3)

}
