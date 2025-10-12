package main
import "fmt"

func main() {
	var tinggi1, bmi1, tinggi2, bmi2, tinggi3, bmi3 float64
	fmt.Scan(&tinggi1, &bmi1)
	fmt.Scan(&tinggi2, &bmi2)
	fmt.Scan(&tinggi3, &bmi3)
	berat1 := bmi1 * (tinggi1 * tinggi1)
	berat2 := bmi2 * (tinggi2 * tinggi2)
	berat3 := bmi3 * (tinggi3 * tinggi3)
	rataRata := (berat1 + berat2 + berat3) / 3
	fmt.Printf("%.2f", rataRata)
}
