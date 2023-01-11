package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type NilaiMahasiswa struct {
	Nama, MataKuliah, IndeksNilai string
	Nilai, ID                     uint
}

var nilaiNilaiMahasiswa = []NilaiMahasiswa{}

func Index(nilai int) string {
	if nilai >= 80 {
		index := "A"
		return index
	} else if nilai >= 70 && nilai < 80 {
		index := "B"
		return index
	} else if nilai >= 60 && nilai < 70 {
		index := "C"
		return index
	} else if nilai >= 50 && nilai < 60 {
		index := "D"
		return index
	} else {
		index := "E"
		return index
	}
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uname, pwd, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte("Username atau Password tidak boleh kosong"))
			return
		}

		if uname == "admin" && pwd == "admin" {
			next.ServeHTTP(w, r)
			return
		}

		w.Write([]byte("Username atau password tidak sesuai"))
		return
	})
}

func PostNilai(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var DataMahasiswa NilaiMahasiswa

	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&DataMahasiswa); err != nil {
				log.Fatal(err)
			}

			if DataMahasiswa.Nilai > 100 {
				w.Write([]byte("Nilai melebihi batas maksimal"))
				return
			}

			jumlahData := len(nilaiNilaiMahasiswa)
			DataMahasiswa.IndeksNilai = Index(int(DataMahasiswa.Nilai))
			DataMahasiswa.ID = uint(jumlahData) + 1

			nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, DataMahasiswa)

		}

		dataMahasiswa, _ := json.Marshal(DataMahasiswa)
		w.Write(dataMahasiswa)
		return
	}
	http.Error(w, "NOT FOUND", http.StatusNotFound)
	return
}

func GetNilai(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		dataNilai, err := json.Marshal(nilaiNilaiMahasiswa)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(dataNilai)
		return
	}
	http.Error(w, "Error...", http.StatusNotFound)
}

func main() {
	server := &http.Server{
		Addr: ":8080",
	}

	//routing
	http.Handle("/post_data", Auth(http.HandlerFunc(PostNilai)))
	http.Handle("/get_data", http.HandlerFunc(GetNilai))

	fmt.Println("Server running at http://localhost:8080")
	server.ListenAndServe()
}
