package main

import 
	"fmt"

func main () {
	var x,y,z int
	var temp int
	
	fmt.Print("masukan nilai 3 nilai: ")
	fmt.Scan(&x,&y,&z)

	if x>y {
		temp = x
		x = y
		y = temp
	} 
	if y>z {
	temp = y
	y = z
	z = temp
	}
	if x>y {
		temp = x
		x = y 
		y = temp
	}

	fmt.Println(x, y, z)
}