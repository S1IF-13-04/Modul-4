package main
import "fmt"
func main() {
	var jam int
	fmt.Print("Masukkan waktu :")
	fmt.Scan(&jam)

	WIB := (jam + 0) % 24
	WITA := (jam + 1) % 24
	WIT := (jam + 2) % 24

	fmt.Printf("%02d:00 WIB\n", WIB)
	fmt.Printf("%02d:00 WITA \n",WITA)
	fmt.Printf("%02d:00 WIT\n",WIT)


}
