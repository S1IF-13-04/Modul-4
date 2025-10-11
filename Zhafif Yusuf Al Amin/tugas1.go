package main
import "fmt"
func main(){
	var hargaAwal, diskon int

	fmt.Print("Total Belanja Awal: ")
	fmt.Scan(&hargaAwal)
	fmt.Print("Dapat Diskon Dalam Persen: ")
	fmt.Scan(&diskon)

	potongan := (hargaAwal * diskon) / 100
	total := hargaAwal - potongan

	fmt.Print(total)
}