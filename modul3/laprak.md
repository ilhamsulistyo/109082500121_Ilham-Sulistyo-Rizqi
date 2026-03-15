# <h1 align="center">Laporan Praktikum Modul 2 - ... </h1>
<p align="center">Ilham Sulistyo Rizqi - 109082500121</p>

## Unguided 

### 1. Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika
diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng
untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian
membantu Jonas? (tidak tentunya ya :p)
Masukan terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi,
dengan syarat a ≥ c dan b ≥ d.
Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi a
terhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d.

#### soal1.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/ilhamsulistyo/109082500121_Ilham-Sulistyo-Rizqi/blob/main/modul3/output1.png)
Program diatas menghitung permutasi dan kombinasi, dimana bilangan akan difaktorialkan di dalam suatu function kemudian dioperasikan ke dalam rumus kombinasi dan permutasi yang telah dibentuk dalam suatu function juga. fungsi tersebut akan dipanggil melalui "func main" dengan meng-inputkan masing-masing variabel int di dalam fungsi kombinasi dan permutasi tersebut.  
## Unguided 

### 2. Diberikan tiga buah fungsi matematika yaitu f (x) = x*x
, g (x) = x − 2 dan h (x) = x +
1. Fungsi komposisi (fogoh)(x) artinya adalah f(g(h(x))). Tuliskan f(x), g(x) dan h(x)
dalam bentuk function.
Masukan terdiri dari sebuah bilangan bulat a, b dan c yang dipisahkan oleh spasi.
Keluaran terdiri dari tiga baris. Baris pertama adalah (fogoh)(a), baris kedua (gohof)(b),
dan baris ketiga adalah (hofog)(c)!
#### soal2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/ilhamsulistyo/109082500121_Ilham-Sulistyo-Rizqi/blob/main/modul3/output2.png)
Program diatas menghitung nilai fungsi fogoh, gohof, dan hofog, dimana setiap fungsi sudah dibentuk di dalam suatu function f, g, dan h. setiap fungsi memiliki operasi nya masing-masing, sebagai berikut :
1. f (x) = x*x
2. g (x) = x − 2
3. h (x) = x + 1.
fungsi tersebut akan dipanggil melalui "func main" dengan meng-inputkan masing-masing variabel int (a, b, dan c) dan dioperasikan kembali di fungsi fogoh, gohof, dan hofog.
## Unguided 

### 3. [Lingkaran] Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius
r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y)
berdasarkan dua lingkaran tersebut.
Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat
dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik
sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan
bilangan bulat.
Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik
di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".
#### soal3.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/ilhamsulistyo/109082500121_Ilham-Sulistyo-Rizqi/blob/main/modul3/output3.png)
Program diatas menentukan suatu titik di dalam atau di luar lingkaran 1 dan 2 dan di dalam lingkaran 1 atau lingkaran 2, suatu titik di dalam atau luar lingkaran ditentukan dengan jarak <= r, rumus jarak = √(a − c)2 + (b − d)2. rumus jarak dan penentuan posisi titik (jarak <= r) dibentk dalam function jarak dan didalam, kemudian dipanggil di dalam func main,dan akhirnya output ditentukan oleh if else dan else if.