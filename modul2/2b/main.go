package main

import "fmt"

func main() {
	var (
		berhasil       bool
		g1, g2, g3, g4 string
		i              int
	)

	berhasil = true
	for i = 1; i <= 5; i++ {
		fmt.Print("Percobaan ", i, " : ")
		fmt.Scan(&g1, &g2, &g3, &g4)
		if g1 != "merah" || g2 != "kuning" || g3 != "hijau" || g4 != "ungu" {
			berhasil = false
		}

	}
	fmt.Println("Berhasil: ", berhasil)
}
