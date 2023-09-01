package main

import (
	"fmt"
	"strings"
)

func luasPersegiPanjang() {

	panjang := 12
	lebar := 4

	luas := panjang * lebar

	fmt.Println("Luas Persegi Panjang:", luas)
}

func kelilingPersegiPanjang() {
	panjang := 12
	lebar := 4

	keliling := 2 * (panjang + lebar)

	fmt.Println("Keliling Persegi Panjang:", keliling)
}

func volumeBalok() {
	panjang := 12
	lebar := 4
	tinggi := 8

	volume := panjang * lebar * tinggi

	fmt.Println("Volume Balok:", volume)
}

func introduce(name, gender, occupation, age string) {
	salutation := ""
	rolePrefix := ""

	switch gender {
	case "laki-laki":
		salutation = "Pak"
		rolePrefix = "seorang"
	case "perempuan":
		salutation = "Bu"
		rolePrefix = "seorang"
	}

	introduction := fmt.Sprintf("%s %s adalah %s %s yang berusia %s tahun", salutation, name, rolePrefix, occupation, age)
	fmt.Println(introduction)
}

func buahFavorit(nama string, buah ...string) {

	buahFavorit := strings.Join(buah, "\", \"")

	fmt.Printf("Halo, nama saya %s dan buah favorit saya adalah \"%s\"\n", nama, buahFavorit)
}

func tambahDataFilm(title, duration, genre, year string) {
	film := map[string]string{
		"title":    title,
		"duration": duration,
		"genre":    genre,
		"year":     year,
	}

	dataFilm = append(dataFilm, film)
}

var dataFilm = []map[string]string{}

func main() {

	luasPersegiPanjang()
	kelilingPersegiPanjang()
	volumeBalok()
	fmt.Println()

	introduce("John", "laki-laki", "penulis", "30")
	introduce("Sarah", "perempuan", "model", "28")
	fmt.Println()

	buahFavorit("John", "semangka", "jeruk", "melon", "pepaya")
	fmt.Println()

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}
}
