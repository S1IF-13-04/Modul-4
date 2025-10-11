package main
import "fmt"
func main() {
	var awal, diskon, akhir int
	print("Masukkan nilai awal: ")
	fmt.Scan(&awal)
	print("Masukkan nilai diskon: ")
	fmt.Scan(&diskon)
	akhir = awal * (100 - diskon) / 100
	fmt.Println("total belanja akhir", akhir)
}
