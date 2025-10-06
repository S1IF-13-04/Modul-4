package main

import "fmt"

func main() {
	var price, discount int
	//declares price and discount as intergers

	fmt.Scan(&price, &discount)
	//Scans for them both

	fmt.Println(price - ((price * discount) / 100))
	//Print the result of the price being discounted by the discount

}
