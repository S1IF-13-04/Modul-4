package main
import "fmt"
func main() {
		
	var bmi, tinggi float64
	fmt.Print("Input: ")
	fmt.Scan(&bmi, &tinggi)
	
	berat := bmi * (tinggi * tinggi)

fmt.Printf("Berat Badannya Adalah : %.0f Kg\n", berat)


}
