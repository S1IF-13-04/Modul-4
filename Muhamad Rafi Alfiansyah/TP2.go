package main
import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif (≤15): ")
	fmt.Scan(&n)

	if n <= 0 || n > 15 {
		fmt.Println("Input tidak valid")
	} else {
		biner := ""
		temp := n
		for temp > 0 {
			sisa := temp % 2
			biner = fmt.Sprintf("%d%s", sisa, biner)
			temp = temp / 2
		}
		fmt.Printf("Desimal: %d\n", n)
		fmt.Printf("Biner: %s\n", biner)
	}
}