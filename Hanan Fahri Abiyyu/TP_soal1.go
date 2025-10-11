package main
import (
	"fmt"
	"strconv"
)
func main() {
	var barangx, barangy string
	const  pajak = 12.0
	fmt.Print("Masukkan harga barang x :")
	fmt.Scan(&barangx)
	fmt.Print("Masukkan harga barang y :")
	fmt.Scan(&barangy)

	barang1, _ := strconv.Atoi(barangx)
	barang2, _ := strconv.Atoi(barangy)

	th := barang1 + barang2
	dp := th + (th * pajak / 100)
	fmt.Println("Harga barang setelah terkena pajak adalah", dp)
}
