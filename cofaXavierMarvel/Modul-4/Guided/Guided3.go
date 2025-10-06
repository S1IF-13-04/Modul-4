package main

import "fmt"

func main() {

	var bodyWeight, bodyHeight, bmi float64

	fmt.Scan(&bodyWeight, &bodyHeight)

	bmi = bodyWeight / (bodyHeight * bodyHeight)

	fmt.Printf("%.2f", bmi)
}
