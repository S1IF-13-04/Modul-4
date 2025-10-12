package main

import "fmt"

func main() {
	var ipa, mtk, indo int
	fmt.Print("Masukkan nilai IPA: ")
	fmt.Scan(&ipa)
	fmt.Print("Masukkan nilai MTK: ")
	fmt.Scan(&mtk)
	fmt.Print("Masukkan nilai B.Indonesia: ")
	fmt.Scan(&indo)

	rata := float64(ipa+mtk+indo) / 3
	fmt.Printf("Rata-rata: %.2f\n", rata)
}
