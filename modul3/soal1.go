package main

import "fmt"

func faktorial(n int) int{
	var hasil = 1
	for i := 1; i <= n;i++{
		hasil = hasil*i
	}
	return hasil
}

func permutasi(n,r int) int{
	var hasil int
	if n >= r {
		hasil = faktorial(n)/faktorial(n-r)
	}
	return hasil
}

func kombinasi(n,r int) int{
	var hasil int
	if n >= r {
		hasil = faktorial(n)/(faktorial(n-r)*faktorial(r))
	}
	return hasil
}

func main() {
	var (
		a,b,c,d int
	)
	fmt.Print("Input nilai a, b, c, dan d : ")
	fmt.Scan(&a,&b,&c,&d)	

	fmt.Println(permutasi(a,c),kombinasi(a,c))
	fmt.Println(permutasi(b,d),kombinasi(b,d))
}
