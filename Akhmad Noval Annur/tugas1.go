package main
import "fmt"

func main(){
	var h,d int 
	fmt.Print("Masukkan Total Belanja: ")
	fmt.Scan(&h)
	fmt.Print("Masukkan Diskon(dalam persen): ")
	fmt.Scan(&d)

	d = h * d/100
	akhir := h - d

	fmt.Printf("Harga setelah diskon adalah %d",akhir)
}