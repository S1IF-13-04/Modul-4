package main
import "fmt"

func main(){
	var jarak float64
	fmt.Scan(&jarak)
	kilometer := int(jarak) / 1000
	meter := int(jarak) % 1000
	fmt.Println("Hasil konversi: ", kilometer," km dan ", meter," meter")
	fmt.Printf("Jarak bulat: %.0f meter\n", jarak)
	fmt.Print("Apakah lebih dari 1 km? ", kilometer > 1)
}
