package main
import "fmt"
func main(){
	var jamkerja int
	const upah = 15000 

	fmt.Print("Input Jam Kerja: ")
	fmt.Scan(&jamkerja)

	hasil := jamkerja * upah

	fmt.Print("kamu mendapat gaji : ",hasil)
}