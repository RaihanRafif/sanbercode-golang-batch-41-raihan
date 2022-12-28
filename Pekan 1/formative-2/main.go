// Bootcamp Digital Skill Sanbercode Golang

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// SOAL 1
	text1 := "Bootcamp"
	text2 := "Digital"
	text3 := "Skill"
	text4 := "Sanbercode"
	text5 := "Golang"

	fmt.Println(text1 + ` ` + text2 + ` ` + text3 + ` ` + text4 + ` ` + text5)

	// JAWABAN SOAL 1
	// Bootcamp Digital Skill Sanbercode Golang

	// -------------------------------------------------------------------------------

	// SOAL 2
	// halo := "Halo Dunia"
	// newHalo := strings.Replace(halo, "Dunia", "Golang", 1)

	// fmt.Println(newHalo)

	// JAWABAN SOAL 2
	// Halo Golang

	// -------------------------------------------------------------------------------

	// SOAL 3
	// kataPertama := "saya"
	// kataKedua := "senang"
	// kataKetiga := "belajar"
	// kataKeempat := "golang"

	// newKataKedua := strings.Replace(kataKedua, "senang", "Senang", 1)
	// newKataKetiga := strings.Replace(kataKetiga, "belajar", "belajaR", 1)
	// newKataKeempat := strings.ToUpper(kataKeempat)

	// fmt.Println(kataPertama + " " + newKataKedua + " " + newKataKetiga + " " + newKataKeempat)
	//Jawaban Soal 3
	// saya Senang belajaR GOLANG

	// Soal 4
	// var angkaPertama = "8"
	// var angkaKedua = "5"
	// var angkaKetiga = "6"
	// var angkaKeempat = "7"

	// newAngkaPertama, _ := strconv.Atoi(angkaPertama)

	// newAngkaKedua, _ := strconv.Atoi(angkaKedua)
	// newAngkaKetiga, _ := strconv.Atoi(angkaKetiga)
	// newAngkaKeempat, _ := strconv.Atoi(angkaKeempat)

	// fmt.Println(newAngkaPertama + newAngkaKedua + newAngkaKetiga + newAngkaKeempat)
	//Jawaban SOAL 4
	//26

	// -------------------------------------------------------------------------------

	// SOAL 5
	kalimat := "halo halo bandung"
	angka := 2021

	newKalimat := strings.Replace(kalimat, "halo", "Hi", 2)
	newAngka := strconv.Itoa(angka)

	fmt.Println(`"` + newKalimat + `"` + " " + "-" + " " + newAngka)
	//Jawaban Soal 5
	// "Hi Hi bandung" - 2021

}
