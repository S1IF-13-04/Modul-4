package main
import "fmt"

func main () {
	var harga,diskon int

	fmt.Print("masukan total harga: ")
	fmt.Scanln(&harga)
	fmt.Print("masukan diskon: ")
	fmt.Scanln(&diskon)

	dibayar := float64(harga) * (100-float64(diskon))/100

	fmt.Print("harga yang dibayar adalah: ", dibayar)
}