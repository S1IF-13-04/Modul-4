package main

import "fmt"

func main() {
	var jarakMeter int
	var waktuDetik int

	fmt.Print("Jarak (meter): ")
	fmt.Scanln(&jarakMeter)

	fmt.Print("Waktu (detik): ")
	fmt.Scanln(&waktuDetik)

	jarakTempuhPenuh := jarakMeter / 1000
	sisaJarak := jarakMeter % 1000

	jarakDalamKm := float64(jarakMeter) / 1000.0
	waktuDalamJam := float64(waktuDetik) / 3600.0
	kecepatanRataRata := jarakDalamKm / waktuDalamJam

	fmt.Println("\n--- Hasil Analisis Perjalanan ---")
	fmt.Printf("Jarak Tempuh Penuh (km): %d\n", jarakTempuhPenuh)
	fmt.Printf("Sisa Jarak yang tidak mencapai 1 km: %d meter\n", sisaJarak)
	fmt.Printf("Kecepatan Rata-Rata: %.2f km/jam\n", kecepatanRataRata)
}