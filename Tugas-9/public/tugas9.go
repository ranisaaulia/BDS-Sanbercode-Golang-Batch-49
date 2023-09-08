// mypackage/mypackage.go
package public

import (
	"fmt"
	"math"
)

type SegitigaSamaSisi struct {
	Alas, Tinggi int
}

type PersegiPanjang struct {
	Panjang, Lebar int
}

type Tabung struct {
	JariJari, Tinggi float64
}

type Balok struct {
	Panjang, Lebar, Tinggi int
}

type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}

type HitungBangunRuang interface {
	Volume() float64
	LuasPermukaan() float64
}

type Phone struct {
	Name, Brand string
	Year        int
	Colors      []string
}

// Method hitungBangunDatar untuk segitigaSamaSisi
func (s SegitigaSamaSisi) Luas() int {
	return (s.Alas * s.Tinggi) / 2
}

func (s SegitigaSamaSisi) Keliling() int {
	return 3 * s.Alas
}

// Method hitungBangunDatar untuk persegiPanjang
func (p PersegiPanjang) Luas() int {
	return p.Panjang * p.Lebar
}

func (p PersegiPanjang) Keliling() int {
	return 2 * (p.Panjang + p.Lebar)
}

// Method hitungBangunRuang untuk tabung
func (t Tabung) Volume() float64 {
	return math.Pi * math.Pow(t.JariJari, 2) * t.Tinggi
}

func (t Tabung) LuasPermukaan() float64 {
	return 2 * math.Pi * t.JariJari * (t.JariJari + t.Tinggi)
}

// Method hitungBangunRuang untuk balok
func (b Balok) Volume() float64 {
	return float64(b.Panjang * b.Lebar * b.Tinggi)
}

func (b Balok) LuasPermukaan() float64 {
	return 2 * float64(b.Panjang*b.Lebar+b.Panjang*b.Tinggi+b.Lebar*b.Tinggi)
}

func (p Phone) Display() string {
	colorStr := "Colors: " + p.Colors[0]
	for i := 1; i < len(p.Colors); i++ {
		colorStr += ", " + p.Colors[i]
	}
	return fmt.Sprintf("Name: %s\nBrand: %s\nYear: %d\n%s\n", p.Name, p.Brand, p.Year, colorStr)
}

func LuasPersegi(sisi int, tampilkanKeterangan bool) interface{} {
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

func HitungPenjumlahan(prefix string, angka1, angka2 []int) int {
	hasil := 0
	for _, angka := range angka1 {
		hasil += angka
	}
	for _, angka := range angka2 {
		hasil += angka
	}
	return hasil
}
