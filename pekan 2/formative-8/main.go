package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// SOAL 1
type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

func (p persegiPanjang) luas() int {
	return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() int {
	keliling := 2 * (p.lebar + p.panjang)
	return keliling
}

func (s segitigaSamaSisi) luas() int {
	return 1 / 2 * s.alas * s.tinggi
}

func (s segitigaSamaSisi) keliling() int {
	sisi := math.Sqrt(float64(s.alas)*float64(s.alas) + float64(s.tinggi) + float64(s.tinggi))
	keliling := sisi * 3
	return int(keliling)
}

func (t tabung) volume() float64 {
	volume := t.jariJari * t.jariJari * 2 * math.Phi * t.tinggi
	return volume
}

func (t tabung) luasPermukaan() float64 {
	luasPermukaan := t.jariJari*t.jariJari*math.Pi + 2*t.jariJari*math.Pi*t.tinggi
	return luasPermukaan
}

func (b balok) volume() float64 {
	volume := b.lebar * b.panjang * b.tinggi
	return float64(volume)
}

func (b balok) luasPermukaan() float64 {
	luasPermukaan := 2*b.lebar*b.tinggi + 2*b.panjang*b.tinggi + 2*b.lebar*b.panjang
	return float64(luasPermukaan)
}

//SOAL 2

type phone struct {
	name, brand string
	year        int
	colors      []string
}

type soalDua interface {
	tampilDataDua() string
}

func (p phone) tampilDataDua() string {
	tes := "name : " + p.name + "\nbrand : " + p.brand + "\nyear : " + strconv.Itoa(p.year) + "\ncolors : " + strings.Join(p.colors, " ")

	return tes
}

// SOAL 3
func luasPersegi(sisi int, status bool) interface{} {
	if status == true {
		fmt.Println("luas persegi dengan sisi " + strconv.Itoa(sisi) + " cm adalah" + strconv.Itoa(sisi*2) + "cm")
	} else if status == false {
		fmt.Println(sisi * 2)
	} else if sisi == 0 && status == true {
		fmt.Println("Maaf anda belum menginput sisi dari persegi")
	} else if sisi == 0 && status == false {
		fmt.Println(nil)
	}
	return ""
}

//SOAL 4

func main() {
	var bangunDatar hitungBangunDatar
	var bangunRuang hitungBangunRuang

	bangunDatar = segitigaSamaSisi{2.0, 4.0}
	fmt.Println("===== segitiga Sama Sisi")
	fmt.Println("luas      :", bangunDatar.luas())
	fmt.Println("keliling  :", bangunDatar.keliling())

	bangunDatar = persegiPanjang{2.0, 4.0}
	fmt.Println("===== persegi panjang")
	fmt.Println("luas      :", bangunDatar.luas())
	fmt.Println("keliling  :", bangunDatar.keliling())

	bangunRuang = balok{2, 3, 4}
	fmt.Println("===== balok")
	fmt.Println("volume      :", bangunRuang.volume())
	fmt.Println("luas permukaan  :", bangunRuang.luasPermukaan())

	bangunRuang = tabung{2, 3}
	fmt.Println("===== tabung")
	fmt.Println("volume      :", bangunRuang.volume())
	fmt.Println("luas permukaan  :", bangunRuang.luasPermukaan())

	//SOAL 2
	var hpKu soalDua
	color := []string{"red", "blue", "green"}
	hpKu = phone{"tes", "asdas", 2000, color}
	fmt.Println(hpKu.tampilDataDua())

	//SOAL 3
	fmt.Println(luasPersegi(4, true))

	fmt.Println(luasPersegi(8, false))

	fmt.Println(luasPersegi(0, true))

	fmt.Println(luasPersegi(0, false))

	//SOAL 4

	var prefix interface{} = "hasil penjumlahan dari "

	var kumpulanAngkaPertama interface{} = []int{6, 8}
	// angkaPertama := strings.Join(kumpulanAngkaPertama.([]int), ", ")

	var kumpulanAngkaKedua interface{} = []int{12, 14}

	// tulis jawaban anda disini
	// fmt.Println(angkaPertama)
	fmt.Println(prefix, kumpulanAngkaPertama.([]int)[0], " + ", kumpulanAngkaPertama.([]int)[1], kumpulanAngkaKedua.([]int)[0], " + ", kumpulanAngkaPertama.([]int)[1], " = ", kumpulanAngkaPertama.([]int)[0]+kumpulanAngkaPertama.([]int)[1]+kumpulanAngkaKedua.([]int)[0]+kumpulanAngkaKedua.([]int)[1])

}
