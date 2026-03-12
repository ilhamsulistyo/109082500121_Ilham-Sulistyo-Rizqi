package main
import "fmt"
func main() {
	var (
		berat, kg, sisaberat, biaya1, biaya2, totalbiaya int

	)

	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&berat)

	kg = berat / 1000
	sisaberat = berat % 1000
	biaya1 = kg*10000
    
	if sisaberat >= 500 {
		biaya2 = sisaberat*5
	} else {
		biaya2 = sisaberat*15
	}
    
	if kg > 10{
		totalbiaya = biaya1
	}else{
		totalbiaya = biaya1 + biaya2
	}
	
	fmt.Printf("\nDetail berat: %d kg + %d gr\n",kg,sisaberat)
	fmt.Printf("\nDetail biaya:Rp.%d kg + Rp. %d gr\n",biaya1,biaya2)
	fmt.Printf("\nTotal biaya:Rp.%d",totalbiaya)
}