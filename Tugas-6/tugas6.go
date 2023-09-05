package main

import "fmt"

func soal1(jariJari float64, luas *float64, keliling *float64) {

	*luas = 3.14 * jariJari * jariJari
	*keliling = 2 * 3.14 * jariJari

	fmt.Printf("Luas Lingkaran: %.2f\n", *luas)
	fmt.Printf("Keliling Lingkaran: %.2f\n", *keliling)
}

func soal2(sentence *string, name string, gender string, occupation string, age string) {
	var salutation string

	if gender == "laki-laki" {
		salutation = "Pak"
	} else if gender == "perempuan" {
		salutation = "Bu"
	} else {
		salutation = ""
	}

	*sentence = fmt.Sprintf("%s %s adalah seorang %s yang berusia %s tahun", salutation, name, occupation, age)
}

func soal3(buah *[]string, namaBuah string) {
	*buah = append(*buah, namaBuah)
}

var dataFilm = []map[string]string{}

func soal4(judul string, durasi string, genre string, tahun string, dataFilm *[]map[string]string) {
	film := make(map[string]string)
	film["Judul"] = judul
	film["Durasi"] = durasi
	film["Genre"] = genre
	film["Tahun"] = tahun
	*dataFilm = append(*dataFilm, film)
}

func main() {

	var luasLingkaran, kelilingLingkaran float64
	jariJariLingkaran := 5.0
	soal1(jariJariLingkaran, &luasLingkaran, &kelilingLingkaran)
	fmt.Println("-----------------------------------------------------------------------------------------")

	var sentence string
	soal2(&sentence, "John", "laki-laki", "penulis", "30")
	fmt.Println(sentence)

	soal2(&sentence, "Sarah", "perempuan", "model", "28")
	fmt.Println(sentence)
	fmt.Println("-----------------------------------------------------------------------------------------")

	var buah = []string{}
	soal3(&buah, "Jeruk")
	soal3(&buah, "Semangka")
	soal3(&buah, "Mangga")
	soal3(&buah, "Strawberry")
	soal3(&buah, "Durian")
	soal3(&buah, "Manggis")
	soal3(&buah, "Alpukat")

	for i, nama := range buah {
		fmt.Printf("%d. %s\n", i+1, nama)
	}
	fmt.Println("-----------------------------------------------------------------------------------------")

	soal4("LOTR", "2 jam", "action", "1999", &dataFilm)
	soal4("Avenger", "2 jam", "action", "2019", &dataFilm)
	soal4("Spiderman", "2 jam", "action", "2004", &dataFilm)
	soal4("Juon", "2 jam", "horror", "2004", &dataFilm)

	for i, film := range dataFilm {
		fmt.Printf("%d. Judul: %s\n", i+1, film["Judul"])
		fmt.Printf("   Durasi: %s\n", film["Durasi"])
		fmt.Printf("   Genre: %s\n", film["Genre"])
		fmt.Printf("   Tahun: %s\n", film["Tahun"])
	}
}
