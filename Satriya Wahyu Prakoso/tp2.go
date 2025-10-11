package main
import "fmt"
func main() {
	var totalHari int
	fmt.Print("Masukkan jumlah hari: ")
	fmt.Scan(&totalHari)
	tahun := totalHari / 365
	sisa := totalHari % 365
	bulan := sisa / 30
	hari := sisa % 30
	fmt.Printf("%d tahun, %d bulan, %d hari\n", tahun, bulan, hari)
}