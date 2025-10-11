package main
import "fmt"
func main(){
	var tb, bb, bmi float64
	fmt.Scan(&bb, &tb)
	bmi = bb / (tb * tb)
	fmt.Printf("%.2f", bmi)
}