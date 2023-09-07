package main

import (
	"fmt"
	"math"
)

type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

type phone struct {
	name, brand string
	year        int
	colors      []string
}

// Method hitungBangunDatar untuk segitigaSamaSisi
func (s segitigaSamaSisi) luas() int {
	return (s.alas * s.tinggi) / 2
}

func (s segitigaSamaSisi) keliling() int {
	return 3 * s.alas
}

// Method hitungBangunDatar untuk persegiPanjang
func (p persegiPanjang) luas() int {
	return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() int {
	return 2 * (p.panjang + p.lebar)
}

// Method hitungBangunRuang untuk tabung
func (t tabung) volume() float64 {
	return math.Pi * math.Pow(t.jariJari, 2) * t.tinggi
}

func (t tabung) luasPermukaan() float64 {
	return 2 * math.Pi * t.jariJari * (t.jariJari + t.tinggi)
}

// Method hitungBangunRuang untuk balok
func (b balok) volume() float64 {
	return float64(b.panjang * b.lebar * b.tinggi)
}

func (b balok) luasPermukaan() float64 {
	return 2 * float64(b.panjang*b.lebar+b.panjang*b.tinggi+b.lebar*b.tinggi)
}

func (p phone) display() string {
	colorStr := "Colors: " + p.colors[0]
	for i := 1; i < len(p.colors); i++ {
		colorStr += ", " + p.colors[i]
	}
	return fmt.Sprintf("Name: %s\nBrand: %s\nYear: %d\n%s\n", p.name, p.brand, p.year, colorStr)
}

func luasPersegi(sisi int, tampilkanKeterangan bool) interface{} {
	if sisi == 0 {
		if tampilkanKeterangan {
			return "Maaf anda belum menginput sisi dari persegi"
		}
		return nil
	}

	luas := sisi * sisi
	if tampilkanKeterangan {
		return fmt.Sprintf("Luas persegi dengan sisi %d cm adalah %d cm", sisi, luas)
	}
	return luas
}

func hitungPenjumlahan(prefix string, angka1, angka2 []int) int {
	hasil := 0
	for _, angka := range angka1 {
		hasil += angka
	}
	for _, angka := range angka2 {
		hasil += angka
	}
	return hasil
}

func main() {
	// Menampilkan Segitiga Sama Sisi
	segitiga := segitigaSamaSisi{alas: 3, tinggi: 5}
	fmt.Println("Segitiga Sama Sisi:")
	fmt.Printf("Luas: %d\n", segitiga.luas())
	fmt.Printf("Keliling: %d\n", segitiga.keliling())

	// Menampilkan Persegi Panjang
	persegi := persegiPanjang{panjang: 3, lebar: 7}
	fmt.Println("\nPersegi Panjang:")
	fmt.Printf("Luas: %d\n", persegi.luas())
	fmt.Printf("Keliling: %d\n", persegi.keliling())

	// Menampilkan Tabung
	silinder := tabung{jariJari: 2.5, tinggi: 9}
	fmt.Println("\nTabung:")
	fmt.Printf("Volume: %.2f\n", silinder.volume())
	fmt.Printf("Luas Permukaan: %.2f\n", silinder.luasPermukaan())

	// Menampilkan Balok
	kubus := balok{panjang: 3, lebar: 3, tinggi: 3}
	fmt.Println("\nBalok:")
	fmt.Printf("Volume: %.2f\n", kubus.volume())
	fmt.Printf("Luas Permukaan: %.2f\n", kubus.luasPermukaan())

	SamsungGalaxy := phone{
		name:   "Samsung Galaxy Note 20",
		brand:  "Samsung Galaxy Note 20",
		year:   2020,
		colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}

	// Menampilkan data ponsel
	fmt.Println()
	fmt.Println(SamsungGalaxy.display())

	// Menampilkan Luas Persegi
	fmt.Println(luasPersegi(4, true))
	fmt.Println(luasPersegi(8, false))
	fmt.Println(luasPersegi(0, true))
	fmt.Println(luasPersegi(0, false))

	// Menampilkan hitung penjumlahan
	fmt.Println()
	prefix := "hasil penjumlahan dari"
	kumpulanAngkaPertama := []int{6, 8}
	kumpulanAngkaKedua := []int{12, 14}

	hasil := hitungPenjumlahan(prefix, kumpulanAngkaPertama, kumpulanAngkaKedua)

	fmt.Printf("%s %v + %v = %d\n", prefix, kumpulanAngkaPertama, kumpulanAngkaKedua, hasil)
}
