package components

import (
	"fmt"
	"strconv"
)

type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}

type HitungBangunRuang interface {
	Volume() float64
	LuasPermukaan() float64
}

// SOAL 2
type SoalDua interface {
	TampilDataDua() string
}

// SOAL 3
func LuasPersegi(sisi int, status bool) interface{} {
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
//SOAL 4

var Prefix interface{} = "hasil penjumlahan dari "

var KumpulanAngkaPertama interface{} = []int{6, 8}

var KumpulanAngkaKedua interface{} = []int{12, 14}
