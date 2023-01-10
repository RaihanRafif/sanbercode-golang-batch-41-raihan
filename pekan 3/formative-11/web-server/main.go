package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
)

func volume(r, t float64) float64 {
	volume := math.Pi * r * r * t
	return volume
}

func luasAlas(r float64) float64 {
	luasAlas := math.Pi * r * r
	return luasAlas
}

func kelilingAlas(r float64) float64 {
	kelilingAlas := math.Pi * r * 2
	return kelilingAlas
}

func index(w http.ResponseWriter, r *http.Request) {
	volume := volume(7, 10)
	luasAlas := luasAlas(7)
	kelilingAlas := kelilingAlas(7)

	kalimat := "jariJari : 7, tinggi: 10, volume : " + strconv.Itoa(int(volume)) + ", luas alas: " + strconv.Itoa(int(luasAlas)) + ", keliling alas: " + strconv.Itoa(int(kelilingAlas))
	fmt.Fprintln(w, kalimat)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "halo!")
	})

	http.HandleFunc("/index", index)

	fmt.Println("starting web server at http://localhost:8080/")

	http.ListenAndServe(":8080", nil)
}
