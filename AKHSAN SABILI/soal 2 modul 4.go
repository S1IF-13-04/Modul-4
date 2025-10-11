package main
import ("fmt"
		"math")
func main(){
	var bmi, tb float64
	fmt.Scan(&bmi, &tb)
	bb := bmi * (tb * tb)
	bbint := int (math.Floor(bb + 0.5)) 
	fmt.Println(bbint)
}