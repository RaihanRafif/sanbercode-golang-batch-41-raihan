package main

import "fmt"

func main() {
	// SOAL 1

	// var panjangPersegiPanjang string = "8"
	// var lebarPersegiPanjang string = "5"

	// var alasSegitiga string = "6"
	// var tinggiSegitiga string = "7"

	// var luasPersegiPanjang int
	// var kelilingPersegiPanjang int
	// var luasSegitiga int

	// newPanjangPersegiPanjang, _ := strconv.Atoi(panjangPersegiPanjang)
	// newLebarPersegiPanjang, _ := strconv.Atoi(lebarPersegiPanjang)
	// newAlasSegitiga, _ := strconv.Atoi(alasSegitiga)
	// newTinggiSegitiga, _ := strconv.Atoi(tinggiSegitiga)

	// luasPersegiPanjang = newPanjangPersegiPanjang * newLebarPersegiPanjang
	// fmt.Println("Luas persegi panjang adalah = ", luasPersegiPanjang)

	// kelilingPersegiPanjang = 2*newPanjangPersegiPanjang + 2*newLebarPersegiPanjang
	// fmt.Println("Keliling persegi panjang adalah = ", kelilingPersegiPanjang)

	// luasSegitiga = 1 / 2 * newAlasSegitiga * newTinggiSegitiga
	// fmt.Println("Luas segitiga adalah = ", luasSegitiga)

	// JAWABAN SOAL 1
	// 	Luas persegi panjang adalah =  40
	// Keliling persegi panjang adalah =  26
	// Luas segitiga adalah =  0 -> karena luas segitiga ber variable int

	//SOAL 2

	// var nilaiJohn = 80
	// var nilaiDoe = 50

	// if nilaiJohn >= 80 {
	// 	fmt.Println("INDEX A")
	// } else if nilaiJohn > 70 && nilaiJohn < 80 {
	// 	fmt.Println("INDEX B")
	// } else if nilaiJohn > 60 && nilaiJohn < 70 {
	// 	fmt.Println("INDEX C")
	// } else if nilaiJohn > 50 && nilaiJohn < 60 {
	// 	fmt.Println("INDEX D")
	// } else if nilaiJohn < 50 {
	// 	fmt.Println("INDEX E")
	// }

	// if nilaiDoe >= 80 {
	// 	fmt.Println("INDEX A")
	// } else if nilaiDoe > 70 && nilaiDoe < 80 {
	// 	fmt.Println("INDEX B")
	// } else if nilaiDoe > 60 && nilaiDoe < 70 {
	// 	fmt.Println("INDEX C")
	// } else if nilaiDoe > 50 && nilaiDoe < 60 {
	// 	fmt.Println("INDEX D")
	// } else if nilaiDoe < 50 {
	// 	fmt.Println("INDEX E")
	// }

	// SOAL 3

	// var tanggal = 12
	// var bulan = 2
	// var tahun = 2000

	// var newBulan string

	// switch bulan {
	// case 1:
	// 	newBulan = "januari"
	// case 2:
	// 	newBulan = "februari"
	// case 3:
	// 	newBulan = "maret"
	// case 4:
	// 	newBulan = "april"
	// case 5:
	// 	newBulan = "mei"
	// case 6:
	// 	newBulan = "juni"
	// case 7:
	// 	newBulan = "juli"
	// case 8:
	// 	newBulan = "agustus"
	// case 9:
	// 	newBulan = "september"
	// case 10:
	// 	newBulan = "oktober"
	// case 11:
	// 	newBulan = "november"
	// case 12:
	// 	newBulan = "desember"
	// }

	// fmt.Println(strconv.Itoa(tanggal) + " " + newBulan + " " + strconv.Itoa(tahun))

	// SOAL 4
	// 	Baby boomer, kelahiran 1944 s.d 1964
	// Generasi X, kelahiran 1965 s.d 1979
	// Generasi Y (Millenials), kelahiran 1980 s.d 1994
	// Generasi Z, kelahiran 1995 s.d 2015

	var tahun = 2000

	if tahun >= 1994 && tahun <= 1964 {
		fmt.Println("Baby Boomer")
	} else if tahun >= 1965 && tahun <= 1979 {
		fmt.Println("Generasi X")
	} else if tahun >= 1980 && tahun <= 1994 {
		fmt.Println("Generasi Y (Millenials)")
	} else if tahun >= 1995 && tahun <= 2015 {
		fmt.Println("Generasi Z")
	}
	//JAWABAN NO 4
	// GENERASI Z

}
