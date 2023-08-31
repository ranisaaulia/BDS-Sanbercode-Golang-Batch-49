package main

import "fmt"

func soal1() {

	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Println("Berkualitas")
		} else if i%3 == 0 && i%2 != 0 {
			fmt.Println("I Love Coding")
		} else if 1%2 != 0 {
			fmt.Println("Santai")
		}
	}
}

func soal2() {

	for row := 1; row <= 7; row++ {
		for col := 1; col <= 7; col++ {
			if col <= row {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

func soal3() {

	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	var kalimatBaru = kalimat[2:7]

	fmt.Println(kalimatBaru)
}

func soal4() {

	var sayuran = []string{}
	var sayurSayuran = append(sayuran, "bayam", "buncis", "kangkung", "kubis", "seledri", "tauge", "timun")

	for i, sayurSayuran := range sayurSayuran {
		fmt.Printf("%d. %s\n", i+1, sayurSayuran)

	}
}

func soal5() {

	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}

	for i, v := range satuan {
		fmt.Printf("%s = %d\n", i, v)

	}

	var volumeBalok = satuan["panjang"] * satuan["lebar"] * satuan["tinggi"]
	fmt.Printf("volume balok = %d", volumeBalok)
}

func main() {
	soal1()
	soal2()
	soal3()
	soal4()
	soal5()

}
