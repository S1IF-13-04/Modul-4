package main
import "fmt"
func main() {
	var harga, bayar, kembalian int
	fmt.Print("Masukkan total harga: ")
	fmt.Scan(&harga)
	fmt.Print("Masukkan uang dibayar: ")
	fmt.Scan(&bayar)

	kembalian = bayar - harga
	fmt.Println("\nKembalian:", kembalian)

	fmt.Println("50000 =", kembalian/50000, "lembar")
	kembalian %= 50000

	fmt.Println("20000 =", kembalian/20000, "lembar")
	kembalian %= 20000

	fmt.Println("10000 =", kembalian/10000, "lembar")
	kembalian %= 10000

	fmt.Println("5000  =", kembalian/5000, "lembar")
	kembalian %= 5000

	fmt.Println("2000  =", kembalian/2000, "lembar")
	kembalian %= 2000

	fmt.Println("1000  =", kembalian/1000, "lembar")
}