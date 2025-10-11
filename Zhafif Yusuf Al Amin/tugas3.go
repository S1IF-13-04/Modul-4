package main
import (
	"fmt"
	"math"
)
func main() {
	var x1, y1, x2, y2, x3, y3 float64
	fmt.Print("Masukan x1 dan y1 di pisah dengan spasi: ")
	fmt.Scan(&x1, &y1)
	fmt.Print("Masukan x2 dan y2 di pisah dengan spasi: ")
	fmt.Scan(&x2, &y2)
	fmt.Print("Masukan x3 dan y3 di pisah dengan spasi: ")
	fmt.Scan(&x3, &y3)
	sisi1 := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	sisi2 := math.Sqrt(math.Pow(x3-x2, 2) + math.Pow(y3-y2, 2))
	sisi3 := math.Sqrt(math.Pow(x1-x3, 2) + math.Pow(y1-y3, 2))
	//sqrt itu akar pow itu kuadratnya, jadi akar dari a2 + b2
	//step pertama itu cari	panjang setiap sisi segitiga 
	//step kedua itu cari sisi panjang make teorema Pythagoras

	sisipanjang := math.Max(sisi1, math.Max(sisi2, sisi3))
	fmt.Printf("%.2f\n", sisipanjang)
}
