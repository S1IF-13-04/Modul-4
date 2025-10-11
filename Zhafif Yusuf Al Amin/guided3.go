package main
import "fmt"
func main() {
	var berat, tinggi float64

	fmt.Print("Input: ")
	fmt.Scan(&berat, &tinggi)

	bmi := berat / (tinggi * tinggi)

	fmt.Printf("BMI : %.2f\n", bmi)

}
