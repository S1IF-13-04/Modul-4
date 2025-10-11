package main
import "fmt"
func main() {
	var BMI, tinggi float64
	print("Masukkan nilai BMI dan tinggi badan : ")
	fmt.Scan(&BMI, &tinggi)
	berat := BMI * (tinggi * tinggi)
	fmt.Printf("Berat badan = %.0f", berat)
}
