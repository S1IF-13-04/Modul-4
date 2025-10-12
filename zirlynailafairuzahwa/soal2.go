package main
import "fmt"

func main() {
	var tinggiBadan, bmi, beratBadan float64
	fmt.Scan(&bmi, &tinggiBadan)
	beratBadan = bmi * tinggiBadan * tinggiBadan
	fmt.Printf("%.0f", beratBadan)
}