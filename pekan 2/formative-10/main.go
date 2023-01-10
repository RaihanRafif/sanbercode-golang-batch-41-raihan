package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"sync"
	"time"
)

// soal 1
func soalSatu(kalimat string, tahun int) {
	fmt.Println(kalimat, " ", tahun)
}

// soal 2
func kelilingSegitigaSamaSisi(sisi int, status bool) (string, error) {
	if sisi != 0 && status == true {
		newSisi := strconv.Itoa(sisi)
		keliling := strconv.Itoa(3 * sisi)
		kalimat := "keliling segitiga sama sisinya dengan sisi " + newSisi + " cm adalah " + keliling + "cm"
		return kalimat, nil
	} else if sisi != 0 && status == false {
		keliling := strconv.Itoa(3 * sisi)
		return keliling, nil
	} else if sisi == 0 && status == true {
		return strconv.Itoa(sisi), errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
	} else {
		defer endApp()
		panic("Maaf anda belum menginput sisi dari segitiga sama sisi")
	}
}

func endApp() {
	fmt.Println("End App")
	message := recover()
	fmt.Println("Terjadi Error", message)
}

// soal 3
func tambahAngka(angka int, number *int) {
	*number = *number + angka
}

func cetakAngka(theNumber *int) {
	fmt.Println()
	fmt.Print(*theNumber)
	fmt.Println()
}

// soal 4
func soalEmpat(phones *[]string, data ...string) {
	*phones = append(*phones, data...)
	sort.Strings(*phones)
}

func sayEmpat(phones []string) {
	for i := 1; i < len(phones)+1; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i, "-", phones[i-1])
	}
}

// soal 5
func sayLima(phone []string, wg *sync.WaitGroup) {
	for i := 1; i < len(phone)+1; i++ {
		fmt.Println(i, "-", phone[i-1])
	}
	wg.Done()
}

func main() {
	//soal 1
	defer soalSatu("Golang Backend Development", 2021)

	//soal 2
	fmt.Println(kelilingSegitigaSamaSisi(4, true))

	fmt.Println(kelilingSegitigaSamaSisi(8, false))

	fmt.Println(kelilingSegitigaSamaSisi(0, true))

	fmt.Println(kelilingSegitigaSamaSisi(0, false))

	//soal 3

	angka := 1

	defer cetakAngka(&angka)

	tambahAngka(7, &angka)

	tambahAngka(6, &angka)

	tambahAngka(-1, &angka)

	tambahAngka(9, &angka)

	//soal 4
	var phones = []string{}

	soalEmpat(&phones, "Xiaomi", "Asus", "IPhone", "Samsung", "Oppo", "Realme", "Vivo")
	go sayEmpat(phones)

	//soal 5
	var phonesLima = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}
	sort.Strings(phonesLima)
	var wg sync.WaitGroup

	wg.Add(1)
	go sayLima(phonesLima, &wg)

	//SOAL 6

	var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}

	moviesChannel := make(chan string)

	go getMovies(moviesChannel, movies...)

	for value := range moviesChannel {
		fmt.Println(value)
	}
}

func getMovies(ch chan string, movie ...string) {
	for _, n := range movie {
		ch <- n
	}
	close(ch)
}
