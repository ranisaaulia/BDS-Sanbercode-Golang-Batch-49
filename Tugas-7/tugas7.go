package main

import "fmt"

type Buah struct {
	Nama       string
	Warna      string
	AdaBijinya bool
	Harga      int
}

type Segitiga struct {
	Alas, Tinggi int
}

type Persegi struct {
	Sisi int
}

type PersegiPanjang struct {
	Panjang, Lebar int
}

type Phone struct {
	Name, Brand string
	Year        int
	Colors      []string
}

type Movie struct {
	Property, Title, Genre string
	Duration, Year         int
}

var dataFilm = []Movie{}

func buah() {

	data := []Buah{
		{"Nanas", "Kuning", false, 9000},
		{"Jeruk", "Oranye", true, 8000},
		{"Semangka", "Hijau & Merah", true, 10000},
		{"Pisang", "Kuning", false, 5000},
	}

	fmt.Println("Daftar Buah:")
	fmt.Println("-------------------------------------------------------------")
	fmt.Printf("| %-10s | %-15s | %-10s | %-5s |\n", "Nama", "Warna", "Ada Bijinya", "Harga")
	fmt.Println("-------------------------------------------------------------")

	for _, buah := range data {
		adaBijinya := "Tidak"
		if buah.AdaBijinya {
			adaBijinya = "Ada"
		}
		fmt.Printf("| %-10s | %-15s | %-10s | %-5d |\n", buah.Nama, buah.Warna, adaBijinya, buah.Harga)
	}

	fmt.Println("-------------------------------------------------------------")
}

// Metode luas untuk segitiga
func (s Segitiga) luas() int {
	luasSegitiga := (s.Alas * s.Tinggi) / 2
	fmt.Printf("Luas Segitiga: %d\n", luasSegitiga)
	return luasSegitiga
}

// Metode luas untuk persegi
func (p Persegi) luas() int {
	sisiPersegi := p.Sisi
	luasPersegi := sisiPersegi * sisiPersegi
	fmt.Printf("Luas Persegi: %d\n", luasPersegi)
	return luasPersegi
}

// Metode luas untuk persegiPanjang
func (pp PersegiPanjang) luas() int {
	luasPersegiPanjang := pp.Panjang * pp.Lebar
	fmt.Printf("Luas Persegi Panjang: %d\n", luasPersegiPanjang)
	return luasPersegiPanjang
}

func (p *Phone) addColor(newColor string) {

	// Membuat objek phone dan menambahkan warna ke dalam color
	myPhone := Phone{
		Name:   "Phone Model X",
		Brand:  "Brand Y",
		Year:   2022,
		Colors: []string{"Black", "Silver"},
	}
	myPhone.Colors = append(myPhone.Colors, newColor)

	fmt.Printf("Phone Colors: %v\n", myPhone.Colors)
}

func tambahDataFilm(title string, duration int, genre string, year int, dataFilm *[]Movie) {
	for _, film := range *dataFilm {
		if film.Title == title {
			fmt.Printf("Judul film %s sudah ada\n", title)
			return
		}
	}

	filmBaru := Movie{
		Title:    title,
		Duration: duration,
		Genre:    genre,
		Year:     year,
	}
	*dataFilm = append(*dataFilm, filmBaru)

}

func main() {
	buah()
	fmt.Println()

	(Segitiga{Alas: 3, Tinggi: 6}.luas())
	(Persegi{Sisi: 6}.luas())
	(PersegiPanjang{Panjang: 3, Lebar: 2}.luas())
	fmt.Println()

	newColor := "Grey"
	var myPhone Phone
	myPhone.addColor(newColor)
	fmt.Println()

	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)
	fmt.Println("Data Film:")
	for _, film := range dataFilm {
		fmt.Printf("Title: %s, Duration: %d minutes, Genre: %s, Year: %d\n", film.Title, film.Duration, film.Genre, film.Year)
	}
}
