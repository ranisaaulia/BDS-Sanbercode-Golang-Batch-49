package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type NilaiMahasiswa struct {
	Nama        string `json:"nama"`
	MataKuliah  string `json:"mata_kuliah"`
	IndeksNilai string `json:"indeks_nilai"`
	Nilai       uint   `json:"nilai"`
	ID          uint   `json:"id"`
}

var (
	nilaiNilaiMahasiswa = []NilaiMahasiswa{}
	mutex               = sync.Mutex{}
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uname, pwd, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte("Username atau Password tidak boleh kosong"))
			return
		}

		if uname == "ranisa" && pwd == "123113" {
			next.ServeHTTP(w, r)
			return
		}
		w.Write([]byte("Username atau Password tidak sesuai"))
		return
	})
}

func AddNilaiMahasiswa(w http.ResponseWriter, r *http.Request) {

	// Input data JSON
	var input struct {
		Nama       string `json:"nama" binding:"required"`
		MataKuliah string `json:"mata_kuliah" binding:"required"`
		Nilai      uint   `json:"nilai" binding:"required"`
	}

	// menguraikan data JSON
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Menghitung Indeks Nilai
	var indeksNilai string
	switch {
	case input.Nilai >= 80:
		indeksNilai = "A"
	case input.Nilai >= 70:
		indeksNilai = "B"
	case input.Nilai >= 60:
		indeksNilai = "C"
	case input.Nilai >= 50:
		indeksNilai = "D"
	default:
		indeksNilai = "E"
	}

	// Membuat objek Nilai Mahasiswa
	nilaiMahasiswa := NilaiMahasiswa{
		Nama:        input.Nama,
		MataKuliah:  input.MataKuliah,
		Nilai:       input.Nilai,
		IndeksNilai: indeksNilai,
		ID:          uint(len(nilaiNilaiMahasiswa) + 1),
	}

	// Lock mutex untuk menghindari data race
	mutex.Lock()
	defer mutex.Unlock()

	// Menambahkan Nilai Mahasiswa ke slice Nilai Nilai Mahasiswa
	nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, nilaiMahasiswa)

	// Mengirim respons JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nilaiMahasiswa)
}

func GetNilaiMahasiswa(w http.ResponseWriter, r *http.Request) {
	// Lock mutex untuk menghindari data race
	mutex.Lock()
	defer mutex.Unlock()

	// Mengirim respons JSON dengan semua data Nilai Mahasiswa
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nilaiNilaiMahasiswa)
}

func main() {

	http.Handle("/tambah-nilai", Auth(http.HandlerFunc(AddNilaiMahasiswa)))
	http.HandleFunc("/nilai", GetNilaiMahasiswa)

	fmt.Println("Server Running On Port 2210")
	http.ListenAndServe(":2210", nil)
}
