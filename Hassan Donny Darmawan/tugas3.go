package main
import (
	"fmt"
	"math"
)
func main () {
	var x1,x2,x3,y1,y2,y3,garis_c float64
	
	fmt.Print("masukan x1 dan y1: ")
	fmt.Scanln(&x1,&y1)
	fmt.Print("masukan nilai x2 dan y2: ")
	fmt.Scanln(&x2,&y2)
	fmt.Print("masukan nilai x3 dan y3: ")
	fmt.Scan(&x3,&y3)

	garis_a := x2 - x1
	garis_b := y3 - y2
	garis_c = math.Sqrt(math.Pow(garis_a,2) + math.Pow(garis_b,2))
	
	fmt.Printf("panjang sisi terpanjang yang didapat dari titik tersebut adalah: %g\n ", garis_c)
}