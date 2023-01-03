package main

import (
	"fmt"
	// "math"
)

func main() {
	//Soal 1
	var luasLingkaran float64
	var kelilingLingkaran float64

	rumusLuasLingkaran(&luasLingkaran, 5)
	rumusKelilingLingkaran(&kelilingLingkaran, 5)

	//Soal 2
	var sentence string
	introduce(&sentence, "John", "laki-laki", "penulis", "30")
	fmt.Println(sentence) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	introduce(&sentence, "Sarah", "perempuan", "model", "28")
	fmt.Println(sentence) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	//Soal 3
	var buah = []string{}

	nomorTiga(&buah, "Jeruk", "Semangka", "Mangga", "Strawberry", "Durian", "Manggis", "Alpukat")

	//Soal 4

	var dataFilm = []map[string]string{}

	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	for i, m := range dataFilm {
		fmt.Print(i+1, "\n")
		for k, v := range m {
			fmt.Println(k, "value is", v)
		}
	}

}

func tambahDataFilm(judul string, duration string, genre string, year string, dataFilm *[]map[string]string) {
	*dataFilm = append(*dataFilm, map[string]string{
		"title":    judul,
		"duration": duration,
		"genre":    genre,
		"tahun":    year,
	})
}

func nomorTiga(buah *[]string, buahs ...string) {
	*buah = append(*buah, buahs...)
	fmt.Println(*buah)
	temp := *buah
	for i := 1; i < len(*buah)+1; i++ {
		// fmt.Println(i, "-", *buah)
		fmt.Println(i, "-", temp[i-1])
	}
}

func introduce(sentence *string, kata ...string) {
	if kata[1] == "laki-laki" {
		jk := "pak"
		*sentence = jk + " " + kata[0] + " adalah seorang " + kata[2] + " yang berusia " + kata[3]
	} else {
		jk := "bu"
		*sentence = jk + " " + kata[0] + " adalah seorang " + kata[2] + " yang berusia " + kata[3]
	}
}
func rumusLuasLingkaran(luasLingkaran *float64, r float64) float64 {
	*luasLingkaran = r * r * 22 / 7
	fmt.Println(*luasLingkaran)
	return *luasLingkaran
}

func rumusKelilingLingkaran(kelilingLingkaran *float64, r float64) float64 {
	*kelilingLingkaran = 2 * 22 / 7 * r
	return *kelilingLingkaran
}
