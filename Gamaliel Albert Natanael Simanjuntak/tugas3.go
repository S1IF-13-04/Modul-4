package main
import (
	"fmt"
	"math"
)
func hitungJarak(x1, y1, x2, y2 float64) float64 {
	deltaX := math.Pow(x2-x1, 2)
	deltaY := math.Pow(y2-y1, 2)
	return math.Sqrt(deltaX + deltaY)
}
func main() {
	var xA, yA float64
	var xB, yB float64
	var xC, yC float64
	fmt.Scan(&xA, &yA)
	fmt.Scan(&xB, &yB)
	fmt.Scan(&xC, &yC)
	sisiAB := hitungJarak(xA, yA, xB, yB)
	sisiBC := hitungJarak(xB, yB, xC, yC)
	sisiAC := hitungJarak(xA, yA, xC, yC)
	var sisiTerpanjang float64
	sisiTerpanjang = math.Max(sisiAB, sisiBC)
	sisiTerpanjang = math.Max(sisiTerpanjang, sisiAC)
	fmt.Printf("%.2f\n", sisiTerpanjang)
}