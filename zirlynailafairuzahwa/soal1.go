package main
import "fmt"

func main(){
	var belanjaAwal, diskon, belanjaAkhir int
	fmt.Print("Masukkan total belanja: ")
	fmt.Scan(&belanjaAwal)
	fmt.Print("Masukkan diskon: ")
	fmt.Scan(&diskon)
	belanjaAkhir = belanjaAwal - (belanjaAwal * diskon / 100)
	fmt.Print(belanjaAkhir)
}