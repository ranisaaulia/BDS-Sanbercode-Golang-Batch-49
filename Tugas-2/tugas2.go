package main

import (
	"fmt"
	"strconv"
	"strings"
)

func soal1() {
	text1 := "Bootcamp"
	text2 := "Digital"
	text3 := "Skill"
	text4 := "Sanbercode"
	text5 := "Golang"

	fmt.Println(text1, text2, text3, text4, text5)
}

func soal2() {
	halo := "Halo Dunia"

	Text := strings.Replace(halo, "Dunia", "Golang", 2)

	fmt.Println(Text)
}

func soal3() {
	kataPertama := "saya"
	kataKedua := "senang"
	kataKetiga := "belajar"
	kataKeempat := "golang"

	// Mengonversi kataKetiga menjadi "belajaR" dengan huruf 'R' besar
	kataKetiga = strings.Replace(kataKetiga, "r", "R", 7)

	// Mengonversi kataKeempat menjadi "GOLANG" dengan huruf besar
	kataKeempat = strings.ToUpper(kataKeempat)

	fmt.Println(kataPertama, kataKedua, kataKetiga, kataKeempat)
}

func soal4() {
	angkaPertama := "8"
	angkaKedua := "5"
	angkaKetiga := "6"
	angkaKeempat := "7"

	// mengubah variabel string ke int
	angka1, err1 := strconv.Atoi(angkaPertama)
	angka2, err2 := strconv.Atoi(angkaKedua)
	angka3, err3 := strconv.Atoi(angkaKetiga)
	angka4, err4 := strconv.Atoi(angkaKeempat)

	if err1 == nil && err2 == nil && err3 == nil && err4 == nil {
		fmt.Println("angka1:", angka1)
		fmt.Println("angka2:", angka2)
		fmt.Println("angka3:", angka3)
		fmt.Println("angka4:", angka4)

		// Total
		total := angka1 + angka2 + angka3 + angka4
		fmt.Println("Total:", total)
	}
}

func soal5() {
	kalimat := "halo halo bandung"
	angka := 2021

	kalimat = strings.Replace(kalimat, "halo", "Hi", -1)

	fmt.Printf("%s - %d\n", kalimat, angka)
}

func main() {
	soal1()
	soal2()
	soal3()
	soal4()
	soal5()
}
