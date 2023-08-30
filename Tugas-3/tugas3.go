package main

import (
	"fmt"
	"strconv"
)

func soal1() {

	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"

	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	// Mengubah string ke int
	var panjangPP, _ = strconv.Atoi(panjangPersegiPanjang)
	var lebarPP, _ = strconv.Atoi(lebarPersegiPanjang)
	var alasSg, _ = strconv.Atoi(alasSegitiga)
	var tinggiSg, _ = strconv.Atoi(tinggiSegitiga)

	// Operasi Perhitungan
	luasPersegiPanjang := panjangPP * lebarPP
	kelilingPersegiPanjang := 2 * (panjangPP + lebarPP)
	luasSegitiga := (alasSg * tinggiSg) / 2

	hasil := fmt.Sprintf("Luas Persegi Panjang: %d\nKeliling Persegi Panjang: %d\nLuas Segitiga: %d", luasPersegiPanjang, kelilingPersegiPanjang, luasSegitiga)
	fmt.Println(hasil)

}

func soal2() {

	var nilaiJohn = 80
	var nilaiDoe = 50

	if nilaiJohn >= 80 {
		fmt.Println("A")
	} else if nilaiJohn >= 70 {
		fmt.Println("B")
	} else if nilaiJohn >= 60 {
		fmt.Println("C")
	} else if nilaiJohn >= 50 {
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}

	if nilaiDoe >= 80 {
		fmt.Println("A")
	} else if nilaiDoe >= 70 {
		fmt.Println("B")
	} else if nilaiDoe >= 60 {
		fmt.Println("C")
	} else if nilaiDoe >= 50 {
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}
}

func soal3() {

	var tanggal = 31
	var bulan = 12
	var tahun = 2000
	var Bulan string

	switch bulan {
	case 1:
		Bulan = "Januari"
	case 2:
		Bulan = "Februari"
	case 3:
		Bulan = "Maret"
	case 4:
		Bulan = "April"
	case 5:
		Bulan = "Mei"
	case 6:
		Bulan = "Juni"
	case 7:
		Bulan = "Juli"
	case 8:
		Bulan = "Agustus"
	case 9:
		Bulan = "September"
	case 10:
		Bulan = "Oktober"
	case 11:
		Bulan = "November"
	case 12:
		Bulan = "Desember"
	default:
		Bulan = ""
	}

	tanggalLahir := fmt.Sprintf("%d %s %d", tanggal, Bulan, tahun)
	fmt.Println(tanggalLahir)

}

func soal4() {

	tahunKelahiran := 2000
	var generasi string

	if tahunKelahiran >= 1944 && tahunKelahiran <= 1964 {
		generasi = "Baby Boomer"
	} else if tahunKelahiran >= 1965 && tahunKelahiran <= 1979 {
		generasi = "X"
	} else if tahunKelahiran >= 1980 && tahunKelahiran <= 1994 {
		generasi = "Y (Millennials)"
	} else if tahunKelahiran >= 1995 && tahunKelahiran <= 2015 {
		generasi = "Z"
	} else {
		generasi = ""
	}

	fmt.Printf("Saya merupakan generasi: %s\n", generasi)
}

func main() {
	soal1()
	soal2()
	soal3()
	soal4()
}
