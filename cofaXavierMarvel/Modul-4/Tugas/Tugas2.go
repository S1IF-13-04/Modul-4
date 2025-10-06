package main

import "fmt"

func main() {

	var bodyWeight, bodyHeight, bmi float64

	fmt.Scan(&bmi, &bodyHeight)

	bodyWeight = bmi * (bodyHeight * bodyHeight)

	fmt.Printf("%.f", bodyWeight)

}
