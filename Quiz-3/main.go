package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"quiz-3/config"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func main() {
	db, e := config.MySQL()

	if e != nil {
		log.Fatal(e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}

	fmt.Println("Success")

	router := httprouter.New()
	router.GET("/bangun-datar/persegi-panjang", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		lebar := r.URL.Query().Get("lebar")
		panjang := r.URL.Query().Get("panjang")
		hitung := r.URL.Query().Get("hitung")

		go DatarHandler(w, r, lebar, panjang, hitung)
	})

	router.GET("/bangun-ruang/tabung", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		jariJari := r.URL.Query().Get("jariJari")
		tinggi := r.URL.Query().Get("tinggi")
		hitung := r.URL.Query().Get("hitung")

		go RuangHandler(w, r, jariJari, tinggi, hitung)
	})

	fmt.Println("Server Running at Port 7723")
	log.Fatal(http.ListenAndServe(":7723", router))
}

func DatarHandler(w http.ResponseWriter, r *http.Request, lebar, panjang, hitung string) {
	if hitung != "luas" && hitung != "keliling" {
		http.Error(w, "Parameter hitung tidak valid", http.StatusBadRequest)
		return
	}

	var hasil float64
	if hitung == "luas" {
		lebarFloat, err1 := strconv.ParseFloat(lebar, 64)
		panjangFloat, err2 := strconv.ParseFloat(panjang, 64)
		if err1 != nil || err2 != nil {
			http.Error(w, "Parameter lebar atau panjang tidak valid", http.StatusBadRequest)
			return
		}

		hasil = lebarFloat * panjangFloat
	} else {
		lebarFloat, err1 := strconv.ParseFloat(lebar, 64)
		panjangFloat, err2 := strconv.ParseFloat(panjang, 64)
		if err1 != nil || err2 != nil {
			http.Error(w, "Parameter lebar atau panjang tidak valid", http.StatusBadRequest)
			return
		}

		hasil = 2 * (lebarFloat + panjangFloat)
	}

	fmt.Fprintf(w, "Hasil perhitungan: %f", hasil)
}

func RuangHandler(w http.ResponseWriter, r *http.Request, jariJari, tinggi, hitung string) {

	if hitung != "luasPermukaan" && hitung != "volume" {
		http.Error(w, "Parameter hitung tidak valid", http.StatusBadRequest)
		return
	}

	var hasil float64
	if hitung == "luasPermukaan" {
		jariJariFloat, err1 := strconv.ParseFloat(jariJari, 64)
		tinggiFloat, err2 := strconv.ParseFloat(tinggi, 64)
		if err1 != nil || err2 != nil {
			http.Error(w, "Parameter jariJari atau tinggi tidak valid", http.StatusBadRequest)
			return
		}

		hasil = 2 * math.Pi * jariJariFloat * (jariJariFloat + tinggiFloat)
	} else {
		jariJariFloat, err1 := strconv.ParseFloat(jariJari, 64)
		tinggiFloat, err2 := strconv.ParseFloat(tinggi, 64)
		if err1 != nil || err2 != nil {
			http.Error(w, "Parameter jariJari atau tinggi tidak valid", http.StatusBadRequest)
			return
		}

		hasil = math.Pi * jariJariFloat * jariJariFloat * tinggiFloat
	}

	fmt.Fprintf(w, "Hasil perhitungan: %f", hasil)
}
