package main  
import "fmt"  
func main() {  
var beratBadan, tinggiBadan, bmi float64  
fmt.Print("masukan berat badan dan tinggi badan: ")

fmt.Scan(&beratBadan, &tinggiBadan)  
bmi = beratBadan / (tinggiBadan * tinggiBadan)  

fmt.Printf("%.2f", bmi)  
}