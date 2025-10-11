package main
import ( 
	"fmt"
	"math"
)

func main () {
	var bmi,tinggi float64

	fmt.Print("masukan nilai bmi & tinggi badan (m): ")
	fmt.Scan(&bmi,&tinggi)

	bb := bmi * math.Pow(tinggi,2)

	fmt.Printf("berat badan yaitu: %0.0f", bb)
}