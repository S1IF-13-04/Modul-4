package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung jarak antara dua titik (x1, y1) dan (x2, y2)
func distance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

func main() {
	var x1, y1, x2, y2, x3, y3 float64

	// Input koordinat titik A, B, dan C
	fmt.Scan(&x1, &y1)
	fmt.Scan(&x2, &y2)
	fmt.Scan(&x3, &y3)

	// Hitung panjang sisi-sisi segitiga
	ab := distance(x1, y1, x2, y2)
	bc := distance(x2, y2, x3, y3)
	ca := distance(x3, y3, x1, y1)

	// Tentukan sisi terpanjang
	max := ab
	if bc > max {
		max = bc
	}
	if ca > max {
		max = ca
	}

	// Tampilkan hasil dengan dua angka di belakang koma
	fmt.Printf("%.2f\n", max)
}
