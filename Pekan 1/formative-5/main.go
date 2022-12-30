package main

import (
	"fmt"
	"strings"
)

func luasPersegiPanjang(panjang int, lebar int) (luas int) {
	luas = panjang * lebar
	return
}

func kelilingPersegiPanjang(panjang int, lebar int) (keliling int) {
	keliling = panjang + lebar
	return
}

func volumeBalok(panjang, lebar, tinggi int) (volume int) {
	volume = panjang * lebar * tinggi
	return
}

func main() {
	//SOAL 1
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)

	//SOAL 2
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	//SOAL 3
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}

	var buahFavoritJohn = buahFavorit("John", buah...)

	fmt.Println(buahFavoritJohn)
	// halo nama saya john dan buah favorit saya adalah "semangka", "jeruk", "melon", "pepaya"

	//SOAL 4
	var dataFilm = []map[string]string{}
	// buatlah closure function disini

	var tambahDataFilm = func(str ...string) {
		dataFilm = append(dataFilm, map[string]string{
			"genre": str[2],
			"jam":   str[1],
			"tahun": str[3],
			"title": str[0],
		})
	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}

}

func buahFavorit(name string, buah ...string) string {
	return "halo nama saya " + name + " dan buah favorit saya adalah " + strings.Join(buah, ",")
}

func introduce(kata ...string) (kalimat string) {
	if kata[1] == "laki-laki" {
		jk := "pak"
		kalimat = jk + " " + kata[0] + " adalah seorang " + kata[2] + " " + kata[3]
		return
	} else {
		jk := "bu"
		kalimat = jk + " " + kata[0] + " adalah seorang " + kata[2] + " " + kata[3]
		return
	}
}
