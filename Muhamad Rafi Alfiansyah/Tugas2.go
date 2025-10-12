package main
import "fmt"

func main(){
	var n,tb float64 
	fmt.Print("Masukkan Nilai BMI: ")
	fmt.Scan(&n)
	fmt.Print("Masukkan Tinggi Badan(dalam meter): ")
	fmt.Scan(&tb)

	bb := n * (tb*tb)

	fmt.Printf("Berat Badan: %.0f",bb)
}