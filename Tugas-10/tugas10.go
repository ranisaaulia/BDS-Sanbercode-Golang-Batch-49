package main

import (
	"flag"
	"fmt"
	"math"
	"sort"
)

func sentence(kalimat string, tahun int) {
	fmt.Println("Run Application...")
	fmt.Printf("Kalimat: %s\nTahun: %d\n", kalimat, tahun)
}

func kelilingSegitigaSamaSisi(sisi int, tampilkanString bool) (string, error) {
	if sisi == 0 {
		errMsg := "Maaf anda belum menginput sisi dari segitiga sama sisi"
		if tampilkanString {
			return "", fmt.Errorf(errMsg)
		} else {
			panic(errMsg)
		}
	}

	keliling := sisi * 3

	if tampilkanString {
		return fmt.Sprintf("keliling segitiga sama sisinya dengan sisi %d cm adalah %d cm", sisi, keliling), nil
	} else {
		return fmt.Sprintf("%d", keliling), nil
	}
}

func tambahAngka(nilai int, total *int) {
	*total += nilai
}

func cetakAngka(total *int) {
	fmt.Printf("Total angka: %d\n", *total)
}

var phones = []string{}

func addPhones(phones *[]string) {
	*phones = append(*phones, "Xiaomi", "Asus", "IPhone", "Samsung", "Oppo", "Realme", "Vivo")
}

func hitungLuasLingkaran(r float64) float64 {
	luas := math.Pi * math.Pow(r, 2)
	return math.Round(luas)
}

func hitungKelilingLingkaran(r float64) float64 {
	keliling := 2 * math.Pi * r
	return math.Round(keliling)
}

func hitungLuasPersegiPanjang(panjang, lebar float64) float64 {
	return panjang * lebar
}

func hitungKelilingPersegiPanjang(panjang, lebar float64) float64 {
	return 2 * (panjang + lebar)
}

func main() {
	var kalimat = "Golang Backend Development"
	var tahun = 2021

	defer sentence(kalimat, tahun)

	angka := 1
	defer cetakAngka(&angka)

	tambahAngka(7, &angka)
	tambahAngka(6, &angka)
	tambahAngka(-1, &angka)
	tambahAngka(9, &angka)

	addPhones(&phones)
	sort.Strings(phones)

	for i, phone := range phones {
		fmt.Printf("%d. %s\n", i+1, phone)
	}

	radii := []float64{7, 10, 15}

	for _, r := range radii {
		luas := hitungLuasLingkaran(r)
		keliling := hitungKelilingLingkaran(r)

		fmt.Printf("Jari-jari lingkaran: %.0f\n", r)
		fmt.Printf("Luas lingkaran: %.0f\n", luas)
		fmt.Printf("Keliling lingkaran: %.0f\n", keliling)
	}

	// Mendefinisikan flag untuk panjang dan lebar persegi panjang
	var panjang, lebar float64
	flag.Float64Var(&panjang, "panjang", 0.0, "Panjang dari persegi panjang")
	flag.Float64Var(&lebar, "lebar", 0.0, "Lebar dari persegi panjang")

	// Parsing argumen dari command line
	flag.Parse()

	if panjang <= 0 || lebar <= 0 {
		fmt.Println("Panjang dan lebar harus lebih dari 0.")
		return
	}

	luas := hitungLuasPersegiPanjang(panjang, lebar)
	keliling := hitungKelilingPersegiPanjang(panjang, lebar)

	fmt.Printf("Panjang: %.2f\n", panjang)
	fmt.Printf("Lebar: %.2f\n", lebar)
	fmt.Printf("Luas Persegi Panjang: %.2f\n", luas)
	fmt.Printf("Keliling Persegi Panjang: %.2f\n", keliling)
}
