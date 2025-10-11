package main
import "fmt"
func main() {
    var jam, menit, detik,sisa int
	fmt.Print("Input: ")
	fmt.Scan(&detik)

	jam = detik / 3600
	menit = (detik % 3600)/ 60
	sisa = detik % 60

	fmt.Println(jam, "jam: " ,menit,"menit: " ,sisa ,"detik: ")

}