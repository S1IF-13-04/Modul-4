package main
import "fmt"

func main() {
	var p, l float64
	fmt.Print("Masukkan panjang (cm): ")
	fmt.Scan(&p)
	fmt.Print("Masukkan lebar (cm): ")
	fmt.Scan(&l)

	if p < 0 || l < 0 {
		fmt.Println("Input tidak valid")
	} else {
		luas := p * l
		keliling := 2 * (p + l)
		fmt.Printf("Luas: %.2f cm²\n", luas)
		fmt.Printf("Keliling: %.2f cm\n", keliling)
	}
}