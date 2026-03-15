package main

import (
	"fmt"
	"math"
)

func jarak(a,b,c,d float64) float64{
	var hasil, ac, bd float64
	ac = a-c
	bd = b-d
	hasil = float64(math.Sqrt(ac*ac + bd*bd))
		return hasil
}

func didalam(cx, cy, r, x, y float64) bool{
	var j float64
	j = jarak(x ,y ,cx ,cy)
	return j <= r 
}

func main() {
	var (
		cx1, cy1, r1 float64
		cx2, cy2, r2 float64
		x, y float64
		dalam1, dalam2 bool
	)
	
	fmt.Scan(&cx1,&cy1,&r1)	
	fmt.Scan(&cx2,&cy2,&r2)
	fmt.Scan(&x, &y)
	
	dalam1 = didalam(cx1, cy1, r1, x, y)
	dalam2 = didalam(cx2, cy2, r2, x, y)

	if dalam1 && dalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}