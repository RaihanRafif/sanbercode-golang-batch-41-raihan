package main

import (
	"fmt"
	. "formative-9/components"
)

func main() {
	var bangunDatar HitungBangunDatar
	var bangunRuang HitungBangunRuang

	bangunDatar = SegitigaSamaSisi{Alas: 2.0, Tinggi: 3.0}
	fmt.Println("===== segitiga Sama Sisi")
	fmt.Println("luas      :", bangunDatar.Luas())
	fmt.Println("keliling  :", bangunDatar.Keliling())

	bangunDatar = PersegiPanjang{Panjang: 2.0, Lebar: 4.0}
	fmt.Println("===== persegi panjang")
	fmt.Println("luas      :", bangunDatar.Luas())
	fmt.Println("keliling  :", bangunDatar.Keliling())

	bangunRuang = Balok{Panjang: 2, Lebar: 3, Tinggi: 4}
	fmt.Println("===== balok")
	fmt.Println("volume      :", bangunRuang.Volume())
	fmt.Println("luas permukaan  :", bangunRuang.LuasPermukaan())

	bangunRuang = Tabung{JariJari: 2, Tinggi: 3}
	fmt.Println("===== tabung")
	fmt.Println("volume      :", bangunRuang.Volume())
	fmt.Println("luas permukaan  :", bangunRuang.LuasPermukaan())

	//SOAL 2
	var hpKu SoalDua
	color := []string{"red", "blue", "green"}
	hpKu = Phone{Name: "tes", Brand: "asdas", Year: 2000, Colors: color}
	fmt.Println(hpKu.TampilDataDua())

	//SOAL 3
	fmt.Println(LuasPersegi(4, true))

	fmt.Println(LuasPersegi(8, false))

	fmt.Println(LuasPersegi(0, true))

	fmt.Println(LuasPersegi(0, false))

	//SOAL 4

	fmt.Println(Prefix, KumpulanAngkaPertama.([]int)[0], " + ", KumpulanAngkaPertama.([]int)[1], KumpulanAngkaKedua.([]int)[0], " + ", KumpulanAngkaPertama.([]int)[1], " = ", KumpulanAngkaPertama.([]int)[0]+KumpulanAngkaPertama.([]int)[1]+KumpulanAngkaKedua.([]int)[0]+KumpulanAngkaKedua.([]int)[1])

}
