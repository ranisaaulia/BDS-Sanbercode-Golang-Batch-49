package repository

import (
	"fmt"
	"net/http"
)

func DatarHandler(w http.ResponseWriter, r *http.Request) {

	lebar := r.URL.Query().Get("lebar")
	panjang := r.URL.Query().Get("panjang")
	hitung := r.URL.Query().Get("hitung")

	if !isValidParameterDatar(lebar, panjang, hitung) {
		http.Error(w, "Parameter tidak valid", http.StatusBadRequest)
		return
	}

	hasil := hitungLuasAtauKeliling(lebar, panjang, hitung)

	fmt.Fprintf(w, "Hasil perhitungan: %f", hasil)
}

func RuangHandler(w http.ResponseWriter, r *http.Request) {

	jariJari := r.URL.Query().Get("jariJari")
	tinggi := r.URL.Query().Get("tinggi")
	hitung := r.URL.Query().Get("hitung")

	if !isValidParameterRuang(jariJari, tinggi, hitung) {
		http.Error(w, "Parameter tidak valid", http.StatusBadRequest)
		return
	}

	hasil := hitungVolumeAtauLuasPermukaan(jariJari, tinggi, hitung)

	fmt.Fprintf(w, "Hasil perhitungan: %f", hasil)
}

func isValidParameterDatar(lebar, panjang, hitung string) bool {

	return true
}

func isValidParameterRuang(jariJari, tinggi, hitung string) bool {

	return true
}

func hitungLuasAtauKeliling(lebar, panjang, hitung string) float64 {

	return 0.0
}

func hitungVolumeAtauLuasPermukaan(jariJari, tinggi, hitung string) float64 {

	return 0.0
}
