package main

import "fmt"

func f(x int) int{
	var hasil int
	hasil = x*x
	return hasil
}

func g(x int) int{
	var hasil int
	hasil = x-2
	return hasil
}

func h(x int) int{
	var hasil int
	hasil = x + 1
	return hasil
}

func main() {
	var a,b,c int
	
	fmt.Print("Input nilai a, b, c: ")
	fmt.Scan(&a,&b,&c)	

	fmt.Println(f(g(h(a))))
	fmt.Println(g(h(f(b))))
	fmt.Println(h(f(g(c))))
}
