package main

import (
	"fmt"
	"math"
)

func main() {
	var ax, ay float64
	var bx, by float64
	var cx, cy float64

	fmt.Print("Masukkan x dan y untuk titik A: ")
	fmt.Scan(&ax, &ay)
	fmt.Print("Masukkan x dan y untuk titik B: ")
	fmt.Scan(&bx, &by)
	fmt.Print("Masukkan x dan y untuk titik C: ")
	fmt.Scan(&cx, &cy)

	ab := math.Sqrt(math.Pow(bx-ax, 2) + math.Pow(by-ay, 2))
	bc := math.Sqrt(math.Pow(cx-bx, 2) + math.Pow(cy-by, 2))
	ca := math.Sqrt(math.Pow(ax-cx, 2) + math.Pow(ay-cy, 2))

	sisitp := ab
	if bc > sisitp {
		sisitp = bc
	}
	if ca > sisitp {
		sisitp = ca
	}


	fmt.Printf("%.2f\n", sisitp)
}
