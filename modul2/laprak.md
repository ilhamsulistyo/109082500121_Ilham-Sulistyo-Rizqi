# <h1 align="center">Laporan Praktikum Modul 2 - ... </h1>
<p align="center">Ilham Sulistyo Rizqi - 109082500121</p>

## Unguided 

### 1. Telusuri program berikut dengan cara mengkompilasi dan mengeksekusi program. Silakan
masukan data yang sesuai sebanyak yang diminta program. Perhatikan keluaran yang
diperoleh. Coba terangkan apa sebenarnya yang dilakukan program tersebut?

package main
import “fmt”
func main() {
var (
satu, dua, tiga string
temp string
)
fmt.Print("Masukan input string: ")
fmt.Scanln(&satu)
fmt.Print("Masukan input string: ")
fmt.Scanln(&dua)
fmt.Print("Masukan input string: ")
fmt.Scanln(&tiga)
fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
temp = satu
satu = dua
dua = tiga
tiga = temp
fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
#### soal1.go

```go
package main

import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp            string
	)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/ilhamsulistyo/109082500121_Ilham-Sulistyo-Rizqi/blob/main/modul2/output1.png)
Program diatas mengubah nilai var string dengan var string lainnya. Terdapat 3 variabel  string yang harus diinputkan(satu, dua, tiga) dan satu var string yang dibentuk untuk mengubah nilai var string diawal(temp). Contoh output, Input 3 string awal yaitu 1, 2, dan 3 yang menghasilkan output awal 1, 2, dan 3, kemudian var “temp” diisi dengan nilai var “satu”, var “satu” dengan var “dua”, var “dua” dengan var “tiga”, dan var “tiga” dengan var “temp”, sehingga output akhir yang dihasilkan adalah 2, 3, dan 1

## Unguided 

### 2. TSiswa kelas IPA di salah satu sekolah menengah atas di Indonesia sedang mengadakan
praktikum kimia. Di setiap percobaan akan menggunakan 4 tabung reaksi, yang mana
susunan warna cairan di setiap tabung akan menentukan hasil percobaan. Siswa diminta
untuk mencatat hasil percobaan tersebut. Percobaan dikatakan berhasil apabila susunan
warna zat cair pada gelas 1 hingga gelas 4 secara berturutan adalah ‘merah’, ‘kuning’,
‘hijau’, dan ‘ungu’ selama 5 kali percobaan berulang.
Buatlah sebuah program yang menerima input berupa warna dari ke 4 gelas reaksi
sebanyak 5 kali percobaan. Kemudian program akan menampilkan true apabila urutan
warna sesuai dengan informasi yang diberikan pada paragraf sebelumnya, dan false
untuk urutan warna lainnya.
Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergaris bawah
adalah input/read):
Percobaan 1: merah kuning hijau ungu
Percobaan 2: merah kuning hijau ungu
Percobaan 3: merah kuning hijau ungu
Percobaan 4: merah kuning hijau ungu
Percobaan 5: merah kuning hijau ungu
BERHASIL: true

Percobaan 1: merah kuning hijau ungu 
Percobaan 2: merah kuning hijau ungu 
Percobaan 3: merah kuning hijau ungu 
Percobaan 4: ungu kuning hiau merah
Percobaan 5: merah kuning hijau ungu 
BERHASIL: false
#### soal2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/ilhamsulistyo/109082500121_Ilham-Sulistyo-Rizqi/blob/main/modul2/output2.png)
sebuah program yang menerima input berupa warna dari ke 4 gelas reaksi sebanyak 5 kali percobaan. Kemudian program akan menampilkan true apabila urutan warna sesuai dengan informasi yang diberikan pada paragraf sebelumnya, dan false untuk urutan warna lainnya.
Sebagai contoh :
1.	Percobaan 1: merah kuning hijau ungu
Percobaan 2: merah kuning hijau ungu
Percobaan 3: merah kuning hijau ungu
Percobaan 4: merah kuning hijau ungu
Percobaan 5: merah kuning hijau ungu
BERHASIL: true
2.	Percobaan 1: merah kuning hijau ungu
Percobaan 2: merah kuning hijau ungu
Percobaan 3: merah kuning hijau ungu
Percobaan 4: ungu kuning hijau merah
Percobaan 5: merah kuning hijau ungu
BERHASIL: false

## Unguided 

### 3. PT POS membutuhkan aplikasi perhitungan biaya kirim berdasarkan berat parsel. Maka,
buatlah program BiayaPos untuk menghitung biaya pengiriman tersebut dengan ketentuan
sebagai berikut!
Dari berat parsel (dalam gram), harus dihitung total berat dalam kg dan sisanya (dalam
gram). Biaya jasa pengiriman adalah Rp. 10.000,- per kg. Jika sisa berat tidak kurang dari 500
gram, maka tambahan biaya kirim hanya Rp. 5,- per gram saja. Tetapi jika kurang dari 500
gram, maka tambahan biaya akan dibebankan sebesar Rp. 15,- per gram. Sisa berat (yang
kurang dari 1kg) digratiskan biayanya apabila total berat ternyata lebih dari 10kg.
Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergaris bawah
adalah input/read):
Contoh 1
Berat parsel (gram): 8500
Detail berat: 8 kg + 500 gr
Detail biaya: Rp. 80000 + Rp. 2500
Total biaya: Rp. 82500
Contoh 2
Berat parsel (gram): 9250
Detail berat: 9 kg + 250 gr
Detail biaya: Rp. 90000 + Rp. 3750
Total biaya: Rp. 93750
Contoh 3
Berat parsel (gram): 11750
Detail berat: 11 kg + 750 gr
Detail biaya: Rp. 110000 + Rp. 3750
Total biaya: Rp. 110000
#### soal3.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/ilhamsulistyo/109082500121_Ilham-Sulistyo-Rizqi/blob/main/modul2/output3.png)
Program diatas menghitung total biaya pengiriman suatu paket, total biaya dihitung berdasarkan total jumlah berat paket dalam kilogram dan sisa gram nya, biaya pengiriman paket Rp. 10.000,00 / Kg, Rp. 5,00 per sisa gram >= 500 g, dan Rp. 15,00 per sisa gram < 500 g, sisa gram berat paket tidak dikenakan biaya jika berat paket per Kg lebih dari 10 kg. Dideklarasikan 6 variabel berbentk integer, yaitu varibel n,biaya1, biaya2, total biaya, kg, dan gram, diinput bilangan bulat kedalam var n dalam satuan g, kemudian dioperasikan ke dalam var kg dan gram untuk di ketahui jumlah kg dan gramnya, setelah itu dioperasikan ke dalam var biaya1 untuk biaya per kg dan if else untuk menghitung biaya2 dalam gram, selanjutnya akan dioperasikan ke dalam variabel totalbiaya untuk menghitung total biaya, terakhir dioperasikan ke dalam if else untuk mengetahui sisa gram akan dihitung atau tidak. Sebagai contoh jika input var n 8500 g, maka total biaya adalah Rp. 82.500,00, jika input var n 9250 g, maka total biaya adalah Rp. 93.750,00, dan jika input var n 11750  g, maka total biaya adalah Rp. 11.000,00.