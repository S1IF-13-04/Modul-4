package main
import "fmt"
func main() {
	var tugas, uts, uas float64
	fmt.Print("Masukkan nilai tugas, UTS, dan UAS: ")
	fmt.Scan(&tugas, &uts, &uas)
	nilaiAkhir := (tugas * 0.20) + (uts * 0.35) + (uas * 0.45)
	status := nilaiAkhir >= 75
	fmt.Printf("Nilai akhir: %.2f\n", nilaiAkhir)
	fmt.Printf("Status: %t\n", status)
}