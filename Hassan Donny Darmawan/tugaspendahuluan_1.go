package main
import "fmt"

func main () {
	var km,hm,dam,m,dm,cm,mm float64
	fmt.Print("masukan nilai dalam satuan meter: ")
	fmt.Scanln(&m)
	km = m / 1000
	hm = m / 100
	dam = m / 10
	dm = m * 10
	cm = m * 100
	mm = m * 1000

	fmt.Printf("kilometer: %0.2f\n", km)
	fmt.Printf("hektometer: %0.2f\n", hm)
	fmt.Printf("deksameter: %0.2f\n", dam)
	fmt.Printf("meter: %0.0f\n", m)
	fmt.Printf("desimeter: %0.0f\n", dm)
	fmt.Printf("centimeter: %0.0f\n", cm)
	fmt.Printf("milimeter: %0.0f", mm)
}