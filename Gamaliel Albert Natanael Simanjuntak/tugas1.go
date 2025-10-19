package main
import "fmt"
func main() {
	var totalBelanja int
	var diskonPersen int
	fmt.Scan(&totalBelanja)
	fmt.Scan(&diskonPersen)
	var besarDiskon int
	besarDiskon = (totalBelanja * diskonPersen) / 100
	var totalAkhir int
	totalAkhir = totalBelanja - besarDiskon
	fmt.Println(totalAkhir)
}