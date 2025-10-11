package main
import (
	"fmt"
	"strconv"
)
func main(){
	var jam, hari string

	fmt.Print("Input: ")
	fmt.Scan(&jam, &hari)

	hasiljam, _ := strconv.Atoi(jam)
	hasilhari, _ := strconv.Atoi(hari)

	hasilperjam := hasiljam / hasilhari
	hasilperhari := hasiljam % hasilhari

	fmt.Println(hasilperjam,"Jam")
	fmt.Println(hasilperhari)
}