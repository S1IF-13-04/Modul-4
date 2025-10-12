package main

import "fmt"

func main() {
	var toeflScore int
	var wawancaraScore int
	var status bool

	fmt.Print("Masukkan nilai TOEFL dan nilai wawancara: ")
	fmt.Scan(&toeflScore, &wawancaraScore)

	status = toeflScore >= 550 && wawancaraScore >= 80
	fmt.Println(status)
}
