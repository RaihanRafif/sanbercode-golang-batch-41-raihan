package components

import (
	"math"
	"strconv"
	"strings"
)

// SOAL 1
func (p PersegiPanjang) Luas() int {
	return p.Panjang * p.Lebar
}

func (p PersegiPanjang) Keliling() int {
	keliling := 2 * (p.Lebar + p.Panjang)
	return keliling
}

func (s SegitigaSamaSisi) Luas() int {
	return 1 / 2 * s.Alas * s.Tinggi
}

func (s SegitigaSamaSisi) Keliling() int {
	sisi := math.Sqrt(float64(s.Alas)*float64(s.Alas) + float64(s.Tinggi) + float64(s.Tinggi))
	keliling := sisi * 3
	return int(keliling)
}

func (t Tabung) Volume() float64 {
	volume := t.JariJari * t.JariJari * 2 * math.Phi * t.Tinggi
	return volume
}

func (t Tabung) LuasPermukaan() float64 {
	luasPermukaan := t.JariJari*t.JariJari*math.Pi + 2*t.JariJari*math.Pi*t.Tinggi
	return luasPermukaan
}

func (b Balok) Volume() float64 {
	volume := b.Lebar * b.Panjang * b.Tinggi
	return float64(volume)
}

func (b Balok) LuasPermukaan() float64 {
	luasPermukaan := 2*b.Lebar*b.Tinggi + 2*b.Panjang*b.Tinggi + 2*b.Lebar*b.Panjang
	return float64(luasPermukaan)
}

// SOAL 2
func (p Phone) TampilDataDua() string {
	tes := "Name : " + p.Name + "\nBrand : " + p.Brand + "\nYear : " + strconv.Itoa(p.Year) + "\nColors : " + strings.Join(p.Colors, " ")

	return tes
}
