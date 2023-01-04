package main

import "fmt"

type buah struct {
	name       string
	warna      string
	adaBijinya bool
	harga      int
}

type segitiga struct {
	alas, tinggi int
}

type persegi struct {
	sisi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type phone struct {
	name, brand string
	year        int
	colors      []string
}

type movie struct {
	title    string
	genre    string
	duration int
	year     int
}

func main() {
	var nanas buah
	nanas.name = "Nanas"
	nanas.warna = "Orange"
	nanas.adaBijinya = false
	nanas.harga = 9000
	fmt.Print(nanas)

	var Jeruk buah
	Jeruk.name = "Jeruk"
	Jeruk.warna = "Orange"
	Jeruk.adaBijinya = true
	Jeruk.harga = 8000
	fmt.Print(Jeruk)

	var semangka buah
	semangka.name = "semangka"
	semangka.warna = "Hijau dan Merah"
	semangka.adaBijinya = true
	semangka.harga = 10000
	fmt.Print(semangka)

	var pisang buah
	pisang.name = "pisang"
	pisang.warna = "kuning"
	pisang.adaBijinya = false
	pisang.harga = 5000
	fmt.Print(pisang)

	//SOAL 2

	var komponenSegitiga segitiga
	komponenSegitiga.alas = 2
	komponenSegitiga.tinggi = 3
	luasSegitiga(komponenSegitiga.alas, komponenSegitiga.tinggi)

	var komponenPersegi persegi
	komponenPersegi.sisi = 3
	luasPersegi(komponenPersegi.sisi)

	type persegiPanjang struct {
		panjang, lebar int
	}
	var komponenPersegiPanjang persegiPanjang
	komponenPersegiPanjang.panjang = 2
	komponenPersegiPanjang.lebar = 3
	luasPersegiPanjang(komponenPersegiPanjang.lebar, komponenPersegiPanjang.panjang)

	//SOAL 3
	var myPhone phone
	myPhone.brand = "samsung"
	myPhone.name = "ultra"
	myPhone.year = 2020
	addColor(myPhone, "red")

	//SOAL 4
	var dataFilm = []movie{}

	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	for i, item := range dataFilm {
		fmt.Println(i+1, " ", item)
	}

}

func tambahDataFilm(judul string, duration int, genre string, year int, dataFilm *[]movie) {

	// type movie struct {
	// 	title    string
	// 	genre    string
	// 	duration int
	// 	year     int
	// }

	var film movie
	film.title = judul
	film.genre = genre
	film.duration = duration
	film.year = year
	*dataFilm = append(*dataFilm, film)
	// "title":    judul,
	// "duration": duration,
	// "genre":    genre,
	// "tahun":    year,

}

func addColor(myPhone phone, color string) {
	myPhone.colors = append(myPhone.colors, color)
	fmt.Println(myPhone)
}

func luasPersegiPanjang(lebar, panjang int) int {
	luas := lebar * panjang
	return luas
}

func luasSegitiga(alas, tinggi int) int {
	luas := alas * tinggi / 2
	return luas
}

func luasPersegi(sisi int) int {
	luas := sisi * sisi
	return luas
}
