package main

import "fmt"

func main() {
	var satu, dua, tiga string
	var temp string

	fmt.Print("Input huruf pertama: ")
	fmt.Scan(&satu)
	fmt.Print("Input huruf kedua: ")
	fmt.Scan(&dua)
	fmt.Print("Input huruf pertama: ")
	fmt.Scan(&tiga)
	fmt.Println("output : ", satu, " + ", dua, " + ", tiga)
	satu = dua
	dua = tiga
	tiga = satu
	temp = satu
	fmt.Println("output akhir : ", temp, " + ", dua, " + ", tiga)
}
